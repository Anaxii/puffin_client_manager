package payments

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"puffin_client_manager/internal/database"
	"puffin_client_manager/pkg/global"
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

	p.handleClientChanges(clientChanges, &listeners)



	// frequently check for new/removed clients
	// handle clients that are running but no longer need to be
	// create list of networks and append smart contracts and UUID to it

	//var event = make(chan PaymentsLog)
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
	//log.Info("Starting event listeners")
	//for _, v := range config.NetworksMap {
	//	go blockchain.ListenToEvents(v, v.BridgeAddress, event)
	//}
	//go blockchain.ListenToEvents(config.Subnet, config.Subnet.BridgeAddress, event)
	//
	//for {
	//	select {
	//	case vLog := <-event:
	//		data, method, err := events.FindEvent([]types.Log{vLog.Log}, h.BridgeABI)
	//		if err != nil {
	//			log.WithFields(log.Fields{"err": err}).Info("Unable to parse event")
	//		}
	//		h.handleEvent(data, method, vLog.Network)
	//		if len(h.BridgeQueue) == 0 {
	//			db.Write([]byte("block"), []byte(vLog.Network.Name), []byte(fmt.Sprintf("%v", vLog.Log.BlockNumber)))
	//		}
	//	}
	//}
}

func (p *PaymentsHandler) handleClientChanges(clientChanges chan primitive.ObjectID, listeners *map[int]global.ClientSettings) {


	for {
		select {
		case id := <-clientChanges:
			log.Println("handle ", id)
			updatedClient, err := p.DB.GetOneClient(id)
			if err != nil {
				log.Println("Client Deleted")
				continue
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
					// restart contract listeners
				}

				(*listeners)[updatedClient.UUID] = updatedClient
			}

			log.Println(listeners, err)
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
