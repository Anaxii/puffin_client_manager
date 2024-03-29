package client

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"puffin_client_manager/internal/database"
	"puffin_client_manager/pkg/blockchain"
	"puffin_client_manager/pkg/db"
	"puffin_client_manager/pkg/global"
	"puffin_client_manager/pkg/util"
	"puffin_client_manager/pkg/wallet"
)

var newClientQueue = map[string]global.ClientSettings{}
var clients = map[int]global.ClientSettings{}

func GetClients() map[int]global.ClientSettings {
	return clients
}

func SetClients(c map[int]global.ClientSettings) {
	clients = c
}

func QueueNewClient(c global.ClientSettings, id string) {
	db.Write([]byte("new_client"), []byte(id), []byte("pending"))
	newClientQueue[id] = c
}

func StartClientQueueTicker(d database.Database) {
	t := util.SecondsTicker(15)
	for {
		<-t.C
		checkClientQueue(d)
		t = util.SecondsTicker(15)
	}
}

func checkClientQueue(d database.Database) {
	for id, newClient := range newClientQueue {
		db.Write([]byte("new_client"), []byte(id), []byte("verifying"))
		isVerified, reason, _ := verifyClientInfo(newClient)
		log.Println(id, isVerified, reason)
		if isVerified {
			val, err := d.GetCurrentUUID()
			if err != nil {
				db.Write([]byte("new_client"), []byte(id), []byte("denied | reason: Internal error"))
				delete(newClientQueue, id)
				continue
			}
			log.Println(val)
			newClient.UUID = val + 1
			err = d.AddClient(newClient)
			if err != nil {
				db.Write([]byte("new_client"), []byte(id), []byte("denied | reason: Internal error"))
				delete(newClientQueue, id)
				continue
			}
			db.Write([]byte("new_client"), []byte(id), []byte("approved | UUID: " + fmt.Sprintf("%v", newClient.UUID)))
			delete(newClientQueue, id)
		} else {
			db.Write([]byte("new_client"), []byte(id), []byte("denied | reason: " + reason))
			delete(newClientQueue, id)
		}
	}
}

func verifyClientInfo(c global.ClientSettings) (bool, string, error) {
	_, err := wallet.Block(c.RPCURL)
	if err != nil {
		return false, "invalid rpc", err
	}
	_, err = wallet.Block(c.WSURL)
	if err != nil {
		return false, "invalid websocket url", err
	}
	if c.PuffinGeoAddress != "" {
		err = blockchain.GetTier("0x56A52b69179fB4BF0d0Bc9aefC340E63c36d3895", c.PuffinGeoAddress, c.RPCURL)
	} else {
		err = blockchain.GetTier("0x56A52b69179fB4BF0d0Bc9aefC340E63c36d3895", c.PuffinKYCAddress, c.RPCURL)
	}
	if err != nil {
		return false, "geo or kyc contract could not be verified", err
	}
	err = blockchain.GetEpoch(c.PuffinClientAddress, c.RPCURL)
	if err != nil {
		return false, "PuffinClient contract could not be verified", err
	}
	return true, "", nil
}

