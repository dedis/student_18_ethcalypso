package gocontracts

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceDeployWriteRequestHolder(privateKey *ecdsa.PrivateKey, client *ethclient.Client, nonce uint64) (common.Address, *types.Transaction, *WriteRequestHolder, error) {
	auth, e := GetAuth(privateKey, client, nonce)
	if e != nil {
		log.Fatal(e)
	}
	address, tx, instance, err := DeployWriteRequestHolder(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	return address, tx, instance, err
}
