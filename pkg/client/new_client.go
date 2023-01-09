package client

import (
	log "github.com/sirupsen/logrus"
	"puffin_client_manager/internal/database"
	"puffin_client_manager/pkg/blockchain"
	"puffin_client_manager/pkg/db"
	"puffin_client_manager/pkg/global"
	"puffin_client_manager/pkg/util"
	"puffin_client_manager/pkg/wallet"
)

var newClientQueue = map[string]global.ClientSettings{}
var clients = []global.ClientSettings{}

func GetClients() []global.ClientSettings {
	return clients
}

func SetClients(c []global.ClientSettings) {
	clients = c
}

func QueueNewClient(c global.ClientSettings, id string) {
	db.Write([]byte("new_client"), []byte(id), []byte("pending"))
	newClientQueue[id] = c
}

func ListenForNewClients(d database.Database) {
	t := util.SecondsTicker(15)
	for {
		<-t.C
		t = util.SecondsTicker(15)
		for id, newClient := range newClientQueue {
			db.Write([]byte("new_client"), []byte(id), []byte("verifying"))
			isVerified, reason, err := verifyClientInfo(newClient)
			if err != nil {
				log.Error(err)
				continue
			}
			if isVerified {
				val, err := d.GetCurrentUUID()
				if err != nil {
					db.Write([]byte("new_client"), []byte(id), []byte("denied | reason: Internal error"))
					continue
				}
				newClient.UUID = val + 1
				err = d.AddClient(newClient)
				if err != nil {
					db.Write([]byte("new_client"), []byte(id), []byte("denied | reason: Internal error"))
					continue
				}
				db.Write([]byte("new_client"), []byte(id), []byte("approved"))
			} else {
				db.Write([]byte("new_client"), []byte(id), []byte("denied | reason: " + reason))
			}
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

