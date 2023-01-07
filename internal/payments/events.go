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
var signatures = map[string]string{
	logNewEpoch.Hex(): "NewEpoch",
	logNewPayment.Hex(): "NewPayment",
}

type LogNewEpoch struct {
	Epoch       *big.Int
	Timestamp   *big.Int
	Users       *big.Int
	USDCPerUser *big.Int
}

func findEvent(logs []types.Log, abi ethABI.ABI) (interface{}, string, error) {
	var data interface{}
	method := ""
	for _, vLog := range logs {
		method = signatures[vLog.Topics[0].Hex()]
		data, err := parseEvent(abi, method, vLog)
		if err != nil {
			return data, method, err
		}
	}
	return data, method, nil
}

func parseEvent(_abi ethABI.ABI, method string, vLog types.Log) (map[string]interface{}, error) {
	log.WithFields(log.Fields{
		"bridge_address": vLog.Address,
		"block":          vLog.BlockNumber,
		"tx_hash":        vLog.TxHash,
		"method":         method,
	}).Info("New epoch event")

	data := map[string]interface{}{}
	err := _abi.UnpackIntoMap(data, "NewEpoch", vLog.Data)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return data, nil
}
