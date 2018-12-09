package gocontracts

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//ServiceDeployPolicy is just an abstraction of the process of deployment process of creating a policy and deploying it on ethereum.
func ServiceDeployPolicy(privateKey *ecdsa.PrivateKey, client *ethclient.Client, policy []common.Address) (common.Address, *types.Transaction, *Policy, error) {
	nonce, e := GetNonce(privateKey, client)
	if e != nil {
		log.Fatal(e)
	}
	auth, e := GetAuth(privateKey, client, nonce)
	if e != nil {
		log.Fatal(e)
	}
	address, tx, instance, err := DeployPolicy(auth, client, policy)
	return address, tx, instance, err
}
