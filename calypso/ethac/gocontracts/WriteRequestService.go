package gocontracts

import (
	"context"
	"crypto/ecdsa"
	"log"

	"github.com/dedis/cothority"
	"github.com/dedis/kyber"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ServiceDeployWriteRequest(privateKey *ecdsa.PrivateKey, client *ethclient.Client, d []byte, ed []byte, ltsid []byte, p common.Address, U []byte, cs [][]byte, Ubar []byte, F []byte, E []byte, nonce uint64) (common.Address, *types.Transaction, *WriteRequest, error) {
	auth, e := GetAuth(privateKey, client, nonce)
	if e != nil {
		log.Fatal(e)
	}
	temp := make([]byte, 0)
	slice := make([]int64, 1)
	slice[0] = 0
	for j := 0; j < len(cs); j++ {
		temp = append(temp, cs[j]...)
		slice = append(slice, int64(len(cs[j])))
	}
	address, tx, instance, err := DeployWriteRequest(auth, client, d, ed, ltsid, p, U, temp, slice, F, E, Ubar)
	if err != nil {
		log.Fatal(err)
	}
	return address, tx, instance, err
}

//Write is a struct that keeps track
//of all the data in a WriteRequest
//store on an ethereum blockchain
type Write struct {
	Data      []byte
	ExtraData []byte
	U         kyber.Point
	LTSID     []byte
	Cs        []kyber.Point
	Ubar      kyber.Point
	E         kyber.Scalar
	F         kyber.Scalar
}

func ServiceGetWriteRequest(privateKey *ecdsa.PrivateKey, client *ethclient.Client, a common.Address) (*Write, error) {
	call, e := CreateCallOpts(context.Background(), privateKey, client)
	if e != nil {
		return nil, e
	}
	wrCaller, e := NewWriteRequestCaller(a, client)
	if e != nil {
		return nil, e
	}
	ltsid, e := wrCaller.LTSID(call)
	if e != nil {
		return nil, e
	}
	data, e := wrCaller.Data(call)
	if e != nil {
		return nil, e
	}
	ed, e := wrCaller.ExtraData(call)
	if e != nil {
		return nil, e
	}
	U, e := wrCaller.U(call)
	if e != nil {
		return nil, e
	}
	Cs, e := wrCaller.Cs(call)
	if e != nil {
		return nil, e
	}
	slice, e := wrCaller.GetSlice(call)
	if e != nil {
		return nil, e
	}
	ubData, e := wrCaller.Ubar(call)
	if e != nil {
		return nil, e
	}
	fData, e := wrCaller.F(call)
	if e != nil {
		return nil, e
	}
	eData, e := wrCaller.E(call)
	if e != nil {
		return nil, e
	}
	temp := make([][]byte, 0)
	for i := 0; i < len(slice)-1; i++ {
		temp = append(temp, Cs[slice[i]:slice[i+1]])
	}
	points := make([]kyber.Point, 0)
	for i := 0; i < len(temp); i++ {
		p := cothority.Suite.Point()
		e := p.UnmarshalBinary(temp[i])
		if e != nil {
			return nil, e
		}
		points = append(points, p)
	}
	point := cothority.Suite.Point()
	e = point.UnmarshalBinary(U)
	if e != nil {
		return nil, e
	}
	Ubar := cothority.Suite.Point()
	e = Ubar.UnmarshalBinary(ubData)
	F := cothority.Suite.Scalar()
	e = F.UnmarshalBinary(fData)

	if e != nil {
		return nil, e
	}
	E := cothority.Suite.Scalar()
	e = E.UnmarshalBinary(eData)
	if e != nil {
		return nil, e
	}
	write := &Write{
		Data:      data,
		ExtraData: ed,
		U:         point,
		LTSID:     ltsid,
		Cs:        points,
		Ubar:      Ubar,
		E:         E,
		F:         F,
	}
	return write, nil
}
