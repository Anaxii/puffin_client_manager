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
				err = p.DB.UpdateClientSettings(e.ClientUUID, "payment_expiration", temp.PaymentExpiration)
				if err != nil {
					log.WithFields(log.Fields{"error": err}).Error("Error updating client settings payment_expiration")
				}
			}
		}
	}
}

func (p *PaymentsHandler) handleClientChanges(
	clients *map[int]global.ClientSettings,
	clientChanges chan primitive.ObjectID,
	listeners *map[int]global.ClientSettings,
	event chan PaymentsLog,
	listenerStatus *map[int]bool) {
	for {
		select {
		case id := <-clientChanges:
			log.Info("handle change:", id)
			updatedClient, err := p.DB.GetOneClient(id)

			if err != nil {
				delete(*listeners, updatedClient.UUID)
				delete(*clients, updatedClient.UUID)
				(*listenerStatus)[updatedClient.UUID] = false
				client.SetClients(*clients)
				log.Info("client deleted:", id)
				return
			}

			clientChange(updatedClient, clients, listeners, event, listenerStatus)
		}
	}
}

func clientChange(
	updatedClient global.ClientSettings,
	clients *map[int]global.ClientSettings,
	listeners *map[int]global.ClientSettings,
	event chan PaymentsLog,
	listenerStatus *map[int]bool) {

	_, ok := (*listeners)[updatedClient.UUID]
	if !ok {
		log.Info("new client:", updatedClient.UUID)
		(*listeners)[updatedClient.UUID] = updatedClient
		(*clients)[updatedClient.UUID] = updatedClient
		client.SetClients(*clients)
		go ListenForEvents(updatedClient.WSURL, updatedClient.UUID, updatedClient.PuffinClientAddress, event, listenerStatus)
	} else {
		(*clients)[updatedClient.UUID] = updatedClient
		c := (*listeners)[updatedClient.UUID]
		if updatedClient.Status != c.Status {
			if updatedClient.Status == "inactive" {
				log.Info("client now inactive:", updatedClient.UUID)
				// stop listener
				(*listenerStatus)[c.UUID] = false
				delete(*listeners, c.UUID)

			} else if updatedClient.Status == "active" {
				log.Info("client now active:", updatedClient.UUID)
				(*listenerStatus)[c.UUID] = true
				go ListenForEvents(updatedClient.WSURL, updatedClient.UUID, updatedClient.PuffinClientAddress, event, listenerStatus)
			}
		}

		if updatedClient.PuffinClientAddress != c.PuffinClientAddress {
			log.Info("client updated kyc address:", updatedClient.UUID)
			go ListenForEvents(updatedClient.WSURL, updatedClient.UUID, updatedClient.PuffinClientAddress, event, listenerStatus)
		}
		(*clients)[updatedClient.UUID] = updatedClient
		(*listeners)[updatedClient.UUID] = updatedClient

		client.SetClients(*clients)
	}
}
