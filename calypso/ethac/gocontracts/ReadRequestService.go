package gocontracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/dedis/cothority"
	"github.com/dedis/kyber"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceDeployReadRequest(privateKey *ecdsa.PrivateKey, client *ethclient.Client, writeAddress common.Address, xc []byte, nonce uint64, hash []byte) (common.Address, *types.Transaction, *ReadRequest, error) {
	auth, err := GetAuth(privateKey, client, nonce)
	address, tx, instance, err := DeployReadRequest(auth, client, writeAddress, xc, hash)
	return address, tx, instance, err
}

func GetAuth(privateKey *ecdsa.PrivateKey, client *ethclient.Client, nonce uint64) (*bind.TransactOpts, error) {
	ctx := context.Background()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Could not cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.From = fromAddress
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4712388) // in units
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))
	return auth, nil
}

type Read struct {
	Write     common.Address
	Xc        kyber.Point
	WriteHash []byte
}

func CreateCallOpts(ctx context.Context, privateKey *ecdsa.PrivateKey, client *ethclient.Client) (*bind.CallOpts, error) {
	call := &bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	return call, nil
}

func ServiceGetRead(a common.Address, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*Read, error) {
	suite := cothority.Suite
	call, e := CreateCallOpts(context.Background(), privateKey, client)
	if e != nil {
		return nil, e
	}
	rrCaller, e := NewReadRequestCaller(a, client)
	if e != nil {
		return nil, e
	}
	write, e := rrCaller.WriteRequest(call)
	if e != nil {
		fmt.Println(e.Error())
		return nil, e
	}
	Xc, e := rrCaller.Xc(call)
	if e != nil {
		return nil, e
	}
	point := suite.Point()
	e = point.UnmarshalBinary(Xc)
	if e != nil {
		return nil, e
	}
	hash, e := rrCaller.WriteHash(call)
	if e != nil {
		return nil, e
	}
	read := &Read{Write: write, Xc: point, WriteHash: hash}
	return read, nil
}
