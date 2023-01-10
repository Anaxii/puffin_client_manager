package payments

import (
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"puffin_client_manager/internal/database"
	"puffin_client_manager/pkg/abi"
	"puffin_client_manager/pkg/client"
	"puffin_client_manager/pkg/global"
	"strings"
	"time"
)

type PaymentsHandler struct {
	Config  global.Config
	DB      database.Database
}

func (p *PaymentsHandler) StartPaymentsHandler() {

	var clients = map[int]global.ClientSettings{}
	_clients, err := p.DB.GetClients()
	if err != nil {
		log.Fatal("could not get clients")
	}

	for _, v := range _clients {
		clients[v.UUID] = v
	}

	go client.SetClients(clients)

	listeners := map[int]global.ClientSettings{}
	listenerStatus := map[int]bool{}
	clientChanges := make(chan primitive.ObjectID)

	for _, v := range clients {
		listeners[v.UUID] = v
		listenerStatus[v.UUID] = true
	}

	event := make(chan PaymentsLog)

	go p.DB.ClientStream(clientChanges)
	go p.handleClientChanges(&clients, clientChanges, &listeners, event, &listenerStatus)


	log.Info("Starting event listeners")
	for _, v := range listeners {
		go ListenForEvents(v.WSURL, v.UUID, v.PuffinClientAddress, event, &listenerStatus)
	}

	clientABI, _ := ethABI.JSON(strings.NewReader(abi.PuffinClientABI))

	for {
		select {
		case e := <-event:
			_, ok := listeners[e.ClientUUID]
			if !ok {
				continue
			}
			_, method, err := findEvent([]types.Log{e.Log}, clientABI)
			if err != nil {
				log.WithFields(log.Fields{"err": err}).Info("Unable to parse event")
			}

			if method == "NewPayment" {
				temp := listeners[e.ClientUUID]
				increment := 28 * 24 * int(time.Hour.Seconds())
				temp.PaymentExpiration += increment
				listeners[e.ClientUUID] = temp
				err = p.DB.UpdateClientSettings(e.ClientUUID, "payment_expiration", listeners[e.ClientUUID])
				if err != nil {
					log.WithFields(log.Fields{"error": err}).Error("Error updating client settings payment_expiration")
				}
			}
		}
	}
}

func (p *PaymentsHandler) handleClientChanges(clients *map[int]global.ClientSettings, clientChanges chan primitive.ObjectID, listeners *map[int]global.ClientSettings, event chan PaymentsLog, listenerStatus *map[int]bool) {
	for {
		select {
		case id := <-clientChanges:
			log.Println("handle ", id)
			updatedClient, err := p.DB.GetOneClient(id)
			if err != nil {
				log.Println("Client Deleted")
				return
			}

			_, ok := (*listeners)[updatedClient.UUID]
			if !ok {
				log.Println("New client")
				(*listeners)[updatedClient.UUID] = updatedClient
				(*clients)[updatedClient.UUID] = updatedClient
				client.SetClients(*clients)
				go ListenForEvents(updatedClient.WSURL, updatedClient.UUID, updatedClient.PuffinClientAddress, event, listenerStatus)
			} else {
				c := (*listeners)[updatedClient.UUID]
				if updatedClient.Status != c.Status {
					if updatedClient.Status == "inactive" {
						log.Println("Client now inactive")
						// stop listener
						(*listenerStatus)[c.UUID] = false
						delete(*listeners, c.UUID)

					} else if updatedClient.Status == "active" {
						log.Println("client now active")
						(*listenerStatus)[c.UUID] = true
						go ListenForEvents(updatedClient.WSURL, updatedClient.UUID, updatedClient.PuffinClientAddress, event, listenerStatus)
					}
				}

				if updatedClient.PuffinClientAddress != c.PuffinClientAddress {
					log.Println("New KYC addresses")
					// restart listener
				}
				(*clients)[updatedClient.UUID] = updatedClient

				client.SetClients(*clients)
				(*listeners)[updatedClient.UUID] = updatedClient
			}

			log.Println(listeners, err)
		}
	}
}
