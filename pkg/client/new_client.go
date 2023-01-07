package client

import (
	log "github.com/sirupsen/logrus"
	"puffin_client_manager/pkg/global"
	"puffin_client_manager/pkg/util"
)

var newClientQueue []global.ClientSettings

func QueueNewClient(c global.ClientSettings) {
	newClientQueue = append(newClientQueue, c)
}

func ListenForNewClients() {
	t := util.SecondsTicker(15)
	for {
		<-t.C
		t = util.SecondsTicker(15)
		for _, newClient := range newClientQueue {
			isVerified, err := verifyClientInfo(newClient)
			if err != nil {
				log.Error(err)
				continue
			}
			if isVerified {
				// give uuid
				// add to db
			}
		}
	}

}

func verifyClientInfo(c global.ClientSettings) (bool, error) {
	// verify website
	// verify rpc url
	// verify ws url
	// verify chain id
	// verify geo/kyc and client addresses are deployed and can be access with rpc
	return true, nil
}

