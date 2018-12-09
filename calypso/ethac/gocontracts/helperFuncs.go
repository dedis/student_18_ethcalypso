package gocontracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

//GetNonce returns the nonce for a given privatekey. i.e. the the next pending nonce.
func GetNonce(privateKey *ecdsa.PrivateKey, client *ethclient.Client) (uint64, error) {
	ctx := context.Background()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return 0, errors.New("Could not cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func DeployEmptyCalypso(privateKey *ecdsa.PrivateKey, client *ethclient.Client, owners []common.Address) (*common.Address, error) {
	fmt.Println("RR holder")
	nonce, e := GetNonce(privateKey, client)
	if e != nil {
		return nil, e
	}
	rrHolderAddr, _, _, e := ServiceDeployReadRequestHolder(privateKey, client, nonce)
	if e != nil {
		return nil, e
	}
	//queuing transactions
	nonce = nonce + 1
	fmt.Println("WR holder")
	wrHolderAddr, _, _, e := ServiceDeployWriteRequestHolder(privateKey, client, nonce)
	if e != nil {
		return nil, e
	}
	//another transaction
	nonce = nonce + 1
	fmt.Println("OWners holder")
	ownersAddress, _, _, e := ServiceDeployOwners(privateKey, client, owners, nonce)
	if e != nil {
		return nil, e
	}
	//another transaction
	nonce = nonce + 1
	fmt.Println("Calypso")
	calAddr, txc, _, e := ServiceCalypso(privateKey, client, ownersAddress, wrHolderAddr, rrHolderAddr, nonce)
	if e != nil {
		return nil, e
	}
	WaitForTransAction(txc, client, 1)
	return &calAddr, e
}
