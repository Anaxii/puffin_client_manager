package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"puffin_client_manager/pkg/abi"
)

func GetTier(walletAddress string, contractAddress string, rpcurl string) error {
	conn, err := ethclient.Dial(rpcurl)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to connect to the Ethereum client")
		return err
	}

	core, err := abi.NewPuffinStatus(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to instantiate PuffinApprovedAccounts contract")
		return err
	}

	_, _, err = core.Status(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to check if user is approved")
		return err
	}

	return err
}

func GetEpoch(contractAddress string, rpcurl string) error {
	conn, err := ethclient.Dial(rpcurl)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to connect to the Ethereum client")
		return err
	}

	core, err := abi.NewPuffinClient(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to instantiate PuffinApprovedAccounts contract")
		return err
	}

	_, err = core.Epoch(nil)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to check if user is approved")
		return err
	}

	return err
}