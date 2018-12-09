package ethereum

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var addresses []common.Address

const localPrivateKey string = "1944dae12efeb1b1107dc1f3c7a459a01d865fff1c4b43a26f1755876aa1b820"
const privateKey string = "e9a7601bfa9e9129095663c259370c1477c641c92abfdee34e6575d009e7d4f7"

//GetPrivateKey returns a privatekey where it is possible
//to send transactions from
func GetPrivateKey() (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(privateKey)
}

//GetAnotherPrivateKey is a function that returns an
//alternative privateKey to its compliment function
//GetPrivateKey
func GetAnotherPrivateKey() (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA("a05b7b4580376959940f3bbdb84dab4780c49e97f47c1e8792c12963552931b3")
}

//GetTestListOfAddresses returns the array of addresses
//that can be used for testing purposes. The addresses
//will be the addresses in order and can be found with ganache
func GetTestListOfAddresses() []common.Address {
	address1 := common.HexToAddress("8A749D6B91C35b8a5Ce812278C73C988a97790aA")
	address2 := common.HexToAddress("F28B17a7D4Ab334584dF6ebfD70FA59c3527CD1e")
	address3 := common.HexToAddress("09E360FeD8580641CE129252417a8646709196f5")
	arr := []common.Address{address1, address2, address3}
	return arr
}
