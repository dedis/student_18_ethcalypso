package calypso

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/dedis/cothority"
	"github.com/dedis/onet"
	"github.com/dedis/student_18_ethcalypso/calypso/ethac/gocontracts"
	"github.com/dedis/student_18_ethcalypso/calypso/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", "Workflow", elapsed)
}

func Test_LogAddress(t *testing.T) {
	a := common.HexToAddress("0xa4971b2c3b6f14cEAF991dF67161220586D657b1")
	//p := common.HexToAddress("0xFb18214babBEFeA6E2c3B21C5C77eD2516e30474")
	privateKey, err := crypto.HexToECDSA("e9a7601bfa9e9129095663c259370c1477c641c92abfdee34e6575d009e7d4f7")
	require.Nil(t, err)
	roster := GetRoster()
	writeClient := NewClient(*roster, privateKey)
	writeClient.LogAddress(a)
}
func Test_workflow(t *testing.T) {

	defer timeTrack(time.Now())
	client, e := ethereum.GetClient()
	require.Nil(t, e)
	fmt.Println("Deploying calypso")
	cal, e := gocontracts.GetStaticCalypso()
	require.Nil(t, e)
	a := *cal
	fmt.Println("Address of Calypso: ", a.String())
	privateKey, err := crypto.HexToECDSA("e9a7601bfa9e9129095663c259370c1477c641c92abfdee34e6575d009e7d4f7")
	require.Nil(t, err)
	roster := GetRoster()
	writeClient := NewClient(*roster, privateKey)
	ltsReply, err := writeClient.CreateLTS(a)
	writeClient.ltsReply = ltsReply
	policy := make([]common.Address, 2)
	writer := common.HexToAddress("6B0F3f0E19cA38063A94D5825789DED79F757c43")
	reader := common.HexToAddress("583DB01D879E80DF911059b59d71209Ca3cfC16f")
	policy[0] = writer
	policy[1] = reader
	fmt.Println("policy")
	p, tx, _, e := gocontracts.ServiceDeployPolicy(privateKey, client, policy)
	gocontracts.WaitForTransAction(tx, client, 1)
	isMined, e := gocontracts.IsMined(client, tx.Hash())
	require.Nil(t, e)
	require.True(t, isMined)
	wr, err, txw := writeClient.AddWrite(ltsReply.LTSID, ltsReply.X, []byte("Sabrina"), a, p)
	require.Nil(t, err)
	readPK, err := crypto.HexToECDSA("229aa299066b691a9ddc1bb5dea861622ec8110315565c48924bbafe79eba5b6")
	require.Nil(t, err)
	readClient := NewClient(*roster, readPK)
	wrhash := txw.Hash().Bytes()
	secret, rr, err, txr := readClient.AddRead(*wr, a, wrhash)
	require.Nil(t, err)
	fmt.Println("RR: ", (*rr).Hex())
	fmt.Println("Secret: ", secret.String())
	//Sleeping waiting to be able to add to calypso
	dk := &DecryptKey{
		Read:     *rr,
		Write:    *wr,
		ReadHash: txr.Hash(),
	}
	dkr, err := readClient.DecryptKey(dk)
	require.Nil(t, err)
	fmt.Println("decrypt key reply ", dkr)
}

func GetRoster() *onet.Roster {
	l := onet.NewTCPTest(cothority.Suite)
	servers, roster, _ := l.GenTree(3, true)
	l.GetServices(servers, calypsoID)
	return roster
}

func Test__MaxTransactions(t *testing.T) {
	defer timeTrack(time.Now())
	client, e := ethereum.GetClient()
	require.Nil(t, e)
	cal := common.HexToAddress("7c52BE2b11935957bE14E0373d1c46f20Df81390")
	privateKey, err := crypto.HexToECDSA("e9a7601bfa9e9129095663c259370c1477c641c92abfdee34e6575d009e7d4f7")
	require.Nil(t, err)
	ctx := context.Background()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		require.True(t, ok)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	roster := GetRoster()
	writeClient := NewClient(*roster, privateKey)
	ltsReply, err := writeClient.CreateLTS(cal)
	writeClient.ltsReply = ltsReply
	policy := make([]common.Address, 2)
	writer := common.HexToAddress("6B0F3f0E19cA38063A94D5825789DED79F757c43")
	reader := common.HexToAddress("583DB01D879E80DF911059b59d71209Ca3cfC16f")
	policy[0] = writer
	policy[1] = reader
	fmt.Println("policy")
	p, tx, _, e := gocontracts.ServiceDeployPolicy(privateKey, client, policy)
	require.Nil(t, e)
	gocontracts.WaitForTransAction(tx, client, 1)
	write := NewWrite(cothority.Suite, ltsReply.LTSID, p, ltsReply.X, []byte("Sabrina"))
	temp := make([][]byte, 0)
	for _, point := range write.Cs {
		marshalled, e := point.MarshalBinary()
		if e != nil {
			require.Nil(t, e)
		}
		temp = append(temp, marshalled)
	}
	wrData, e := write.U.MarshalBinary()
	if e != nil {
		require.Nil(t, e)
	}
	Ubar, e := write.Ubar.MarshalBinary()
	if e != nil {
		require.Nil(t, e)
	}
	F, e := write.F.MarshalBinary()
	if e != nil {
		require.Nil(t, e)
	}
	E, e := write.E.MarshalBinary()
	if e != nil {
		require.Nil(t, e)
	}
	var iterations uint64
	oldNonce, e := client.PendingNonceAt(ctx, fromAddress)
	fmt.Println("oldNonce: ", oldNonce)
	nonce, e := gocontracts.GetNonce(privateKey, client)
	require.Nil(t, e)
	transactions := make([]*types.Transaction, 0)
	for {
		addr, tx, _, e := gocontracts.ServiceDeployWriteRequest(privateKey, client, write.Data, write.ExtraData, write.LTSID, p, wrData, temp, Ubar, F, E, nonce)
		if e != nil {
			break
		}
		transactions = append(transactions, tx)
		newNonce, e := client.PendingNonceAt(ctx, fromAddress)
		if e != nil {
			break
		}
		if newNonce != oldNonce {
			break
		}
		iterations++
		nonce = nonce + 1
		tx2, e := gocontracts.ServiceAddWriteRequest(privateKey, cal, addr, client, nonce)
		if e != nil {
			break
		}
		newNonce, e = client.PendingNonceAt(ctx, fromAddress)
		if e != nil {
			break
		}
		if newNonce != oldNonce {
			break
		}
		transactions = append(transactions, tx2)
		iterations++
		nonce = nonce + 1
	}
	var counter uint64
	for _, t := range transactions {
		fmt.Println("Hash: ", t.Hash())
		if ok, e := gocontracts.IsMined(client, t.Hash()); ok {
			if e != nil {
				fmt.Println(e)
			}
			counter++
		}
	}
	fmt.Println("This many transactions were no longer pending: ", counter)
	fmt.Println("I performed this many transactions: ", iterations)
}
