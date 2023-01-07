package payments

import (
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"math/big"
)

var logNewEpoch = crypto.Keccak256Hash([]byte("NewEpoch(uint256,uint256,uint256,uint256)"))
var logNewPayment = crypto.Keccak256Hash([]byte("NewPayment(uint256,uint256,uint256,uint256,uint256,uint256)"))

type LogNewEpoch struct {
	Timestamp   int64
	Users       *big.Int
	Epoch       *big.Int
	CostPerUser *big.Int
	Block       *big.Int
}

func findEvent(logs []types.Log, abi ethABI.ABI) (interface{}, string, error) {
	var err error
	var data interface{}
	method := ""
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case logNewEpoch.Hex():
			data, err = handleNewEpoch(abi, "BridgeIn", vLog)
			method = "NewEpoch"
		//case logNewPayment.Hex():
		//	data, err = logBridge(abi, "BridgeOut", vLog)
		//	method = "NewPayment"
		}
	}
	return data, method, err
}

func handleNewEpoch(_abi ethABI.ABI, method string, vLog types.Log) (LogNewEpoch, error) {
	log.WithFields(log.Fields{
		"bridge_address": vLog.Address,
		"block":          vLog.BlockNumber,
		"tx_hash":        vLog.TxHash,
		"method":         method,
	}).Info("New bridge event")

	var data LogNewEpoch
	err := _abi.UnpackIntoInterface(&data, method, vLog.Data)
	if err != nil {
		log.WithFields(log.Fields{
			"bridge_address": vLog.Address,
			"block":          vLog.BlockNumber,
			"tx_hash":        vLog.TxHash,
			"method":         method,
		}).Error("Could not unpack event")
		return LogNewEpoch{}, err
	}

	log.Println(vLog.Data, vLog.Topics)
	//data.Epoch = vLog.Data[1]
	//data.Asset = common.HexToAddress(vLog.Topics[2].Hex())
	//data.Amount = vLog.Topics[3].Big()
	//data.Block = int64(vLog.BlockNumber)

	return data, nil
}
