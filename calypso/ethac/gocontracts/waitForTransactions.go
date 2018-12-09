package gocontracts

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WaitForTransAction(tx *types.Transaction, client *ethclient.Client, duration int) {
	fmt.Println("Starting to wait")
	for {
		_, isPending, e := client.TransactionByHash(context.Background(), tx.Hash())
		if e != nil {
			fmt.Println("Error: ", e)
			break
		}
		if !isPending {
			return
		}
		time.Sleep(time.Duration(duration) * time.Second)
	}
}

func IsMined(client *ethclient.Client, txHash common.Hash) (bool, error) {
	_, isPending, e := client.TransactionByHash(context.Background(), txHash)
	if e != nil {
		return false, e
	}
	return !isPending, e
}
