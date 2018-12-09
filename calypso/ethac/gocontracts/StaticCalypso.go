package gocontracts

import (
	"fmt"

	"github.com/dedis/student_18_ethcalypso/calypso/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

var calypsoAddr *common.Address
var calypsoError error

func deploy() {
	/*c := common.HexToAddress("0x86D3Fb6137e53664Eb7ff0e272fCB6eeb2BBC846")
	calypsoAddr = &c*/
	if calypsoAddr != nil {
		fmt.Println("whaddup")
		return
	}
	policy := ethereum.GetTestListOfAddresses()
	privateKey, e := ethereum.GetPrivateKey()
	if e != nil {
		return
	}
	client, e := ethereum.GetClient()
	if e != nil {
		return
	}
	calypsoAddr, calypsoError = DeployEmptyCalypso(privateKey, client, policy)
	if calypsoAddr == nil && calypsoError == nil {
		fmt.Println("what is going on")
	}
}

func GetStaticCalypso() (*common.Address, error) {
	deploy()
	return calypsoAddr, calypsoError
}
