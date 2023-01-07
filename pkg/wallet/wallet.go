package wallet

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
)

func DialEthClient(rpcURL string) (*ethclient.Client, error) {
	conn, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.WithFields(log.Fields{
			"rpcURL": rpcURL,
			"err":    err,
		}).Error("Failed to connect to the Ethereum client")
		return nil, err
	}

	return conn, nil
}

func Block(rpcURL string) (*big.Int, error) {
	conn, err := DialEthClient(rpcURL)
	if err != nil {
		return big.NewInt(0), err
	}

	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.WithFields(log.Fields{
			"network": rpcURL,
			"err":     err,
		}).Error("Failed to get block header")
		return big.NewInt(0), err
	}

	return header.Number, nil
}
