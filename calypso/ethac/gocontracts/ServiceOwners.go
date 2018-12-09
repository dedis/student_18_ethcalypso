package gocontracts

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceDeployOwners(privateKey *ecdsa.PrivateKey, client *ethclient.Client, a []common.Address, nonce uint64) (common.Address, *types.Transaction, *Owners, error) {
	auth, e := GetAuth(privateKey, client, nonce)
	if e != nil {
		log.Fatal(e)
	}

	address, tx, instance, err := DeployOwners(auth, client, a)
	if err != nil {
		log.Fatal(err)
	}
	return address, tx, instance, err
}

func ServiceUpdatOwners(a common.Address, deployed common.Address, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	calypsoTransActor, e := NewCalypsoTransactor(deployed, client)
	if e != nil {
		return nil, e
	}
	nonce, e := GetNonce(privateKey, client)
	if e != nil {
		return nil, e
	}
	auth, e := GetAuth(privateKey, client, nonce)
	if e != nil {
		return nil, e
	}
	tx, e := calypsoTransActor.AddOwner(auth, a)
	return tx, e
}
