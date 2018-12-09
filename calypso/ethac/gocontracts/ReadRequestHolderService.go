package gocontracts

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceDeployReadRequestHolder(privateKey *ecdsa.PrivateKey, client *ethclient.Client, nonce uint64) (common.Address, *types.Transaction, *ReadRequestHolder, error) {
	auth, e := GetAuth(privateKey, client, nonce)
	if e != nil {
		log.Fatal(e)
	}

	address, tx, instance, err := DeployReadRequestHolder(auth, client)
	if err != nil {
		fmt.Println("bjorn")
		log.Fatal(err)
	}
	return address, tx, instance, err
}
