package payments

import (
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"puffin_client_manager/internal/database"
	"puffin_client_manager/pkg/abi"
	"puffin_client_manager/pkg/global"
	"strings"
	"time"
)

type PaymentsHandler struct {
	Config global.Config
	DB     database.Database
}

func (p *PaymentsHandler) StartPaymentsHandler() {

	clients, err := p.DB.GetClients()
	if err != nil {
		log.Fatal("could not get clients")
	}

	listeners := map[int]global.ClientSettings{}
	clientChanges := make(chan primitive.ObjectID)

	for _, v := range clients {
		listeners[v.UUID] = v
	}

	go p.DB.ClientStream(clientChanges)
	go p.handleClientChanges(clientChanges, &listeners)

	event := make(chan PaymentsLog)

	log.Info("Starting event listeners")
	for _, v := range listeners {
		go ListenForEvents(v.WSURL, v.UUID, v.PuffinClientAddress, event)
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

func (p *PaymentsHandler) handleClientChanges(clientChanges chan primitive.ObjectID, listeners *map[int]global.ClientSettings) {
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
			} else {
				c := (*listeners)[updatedClient.UUID]
				if updatedClient.Status != c.Status {
					if updatedClient.Status == "inactive" {
						log.Println("Client now inactive")
						// stop listener
						delete(*listeners, c.UUID)

					} else if updatedClient.Status == "active" {
						log.Println("client now active")
						// start listener

					}
				}

				if updatedClient.PuffinClientAddress != c.PuffinClientAddress {
					log.Println("New KYC addresses")
					// restart listener
				}

				(*listeners)[updatedClient.UUID] = updatedClient
			}

			log.Println(listeners, err)
		}
	}
}

