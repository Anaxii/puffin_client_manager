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



	// frequently check for new/removed clients
	// handle clients that are running but no longer need to be
	// create list of networks and append smart contracts and UUID to it

	event := make(chan PaymentsLog)
	//
	//go h.handleQueue()
	//
	//go func() {
	//	ticker := time.NewTicker(15 * time.Second)
	//	x := 3
	//	for range ticker.C {
	//		for _, v := range config.NetworksMap {
	//			h.updateLastBlock(v, x)
	//		}
	//		h.updateLastBlock(config.Subnet, x)
	//
	//		if x == 3 {
	//			x = 0
	//		} else {
	//			x++
	//		}
	//	}
	//}()
	//
	log.Info("Starting event listeners")
	for _, v := range listeners {
		go ListenForEvents(v.WSURL, v.UUID, v.PuffinClientAddress, event)
	}

	clientABI, _ := ethABI.JSON(strings.NewReader(abi.PuffinClientABI))

	for {
		select {
		case e := <-event:
			data, method, err := findEvent([]types.Log{e.Log}, clientABI)
			if err != nil {
				log.WithFields(log.Fields{"err": err}).Info("Unable to parse event")
			}
			log.Println(data, method)
			//h.handleEvent(data, method, vLog.Network)
			//if len(h.BridgeQueue) == 0 {
			//	db.Write([]byte("block"), []byte(vLog.Network.Name), []byte(fmt.Sprintf("%v", vLog.Log.BlockNumber)))
			//}
			//log.Println(e)
		}
	}
}

func (p *PaymentsHandler) handleClientChanges(clientChanges chan primitive.ObjectID, listeners *map[int]global.ClientSettings) {
	for {
		select {
		case id := <-clientChanges:
			go func() {
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
							delete((*listeners), c.UUID)

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
			}()
		}
	}
}

//func (h *Handler) updateLastBlock(v config.Networks, x int) {
//	walletBlock, _ := wallet.Block(v)
//	h.Blocks[v.Name] = int(walletBlock.Int64())
//	if walletBlock.Int64() > 0 {
//		if len(h.BridgeQueue) == 0 {
//			log.WithFields(log.Fields{
//				"block":   walletBlock.Int64(),
//				"network": v.Name,
//			}).Info("Updated last synced block")
//
//			if x == 3 {
//				go api.Log(LogHistory{Status: "Updated last synced block", Message: "lastBlock", Log: BridgeRequest{Block: walletBlock.Int64(), NetworkIn: v}, Timestamp: time.Now().Unix()})
//			}
//
//			db.Write([]byte("block"), []byte(v.Name), []byte(fmt.Sprintf("%v", walletBlock.Int64())))
//		} else {
//			log.WithFields(log.Fields{
//				"block":      walletBlock.Int64(),
//				"queue_size": len(h.BridgeQueue),
//				"network":    v.Name,
//			}).Info("Last block")
//			for _, v := range h.BridgeQueue {
//				log.Info(v.Method)
//			}
//		}
//	}
//}
