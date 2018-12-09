package gocontracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceCalypso(privateKey *ecdsa.PrivateKey, client *ethclient.Client, owners common.Address, wrHolder common.Address, rrHolder common.Address, nonce uint64) (common.Address, *types.Transaction, *Calypso, error) {
	auth, e := GetAuth(privateKey, client, nonce)
	if e != nil {
		log.Fatal(e)
	}

	address, tx, instance, err := DeployCalypso(auth, client, wrHolder, rrHolder, owners)
	if err != nil {
		log.Fatal(err)
	}
	return address, tx, instance, err
}

func ServiceAddWriteRequest(privateKey *ecdsa.PrivateKey, cal common.Address, wr common.Address, client *ethclient.Client, nonce uint64) (*types.Transaction, error) {
	calTransActor, e := NewCalypsoTransactor(cal, client)
	if e != nil {
		log.Fatal(e)
	}
	auth, e := GetAuth(privateKey, client, nonce)
	if e != nil {
		return nil, e
	}
	tx, e := calTransActor.AddWriteRequest(auth, wr)
	return tx, e

}

func ServiceAddReadRequest(privateKey *ecdsa.PrivateKey, cal common.Address, rr common.Address, client *ethclient.Client, nonce uint64) (*types.Transaction, error) {
	calTransActor, e := NewCalypsoTransactor(cal, client)
	if e != nil {
		return nil, e
	}
	auth, e := GetAuth(privateKey, client, nonce)
	if e != nil {
		return nil, e
	}
	tx, e := calTransActor.AddReadRequest(auth, rr)
	return tx, e
}

func ServiceCheckIfCanRead(privateKey *ecdsa.PrivateKey, cal common.Address, rr common.Address, client *ethclient.Client) bool {
	call, e := CreateCallOpts(context.Background(), privateKey, client)
	if e != nil {
		return false
	}
	calypsoCaller, e := NewCalypsoCaller(cal, client)
	canRead, e := calypsoCaller.IsValidReadRequest(call, rr)
	if e != nil {
		fmt.Println("Can read: ", canRead)
		fmt.Println("I can't read because: ", e)
		return false
	}
	return canRead
}
