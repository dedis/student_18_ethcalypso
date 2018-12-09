package ethereum

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client
var e error

const connectionString string = "https://ropsten.infura.io/"
const localConnectionString string = "http://127.0.0.1:7545"

func init() {
	client, e = ethclient.Dial(connectionString)
	if e != nil {
		fmt.Println("Bjo")
		log.Fatal(e)
		client.Close()
	}
}

//GetClient returns the client and the error
//if they were able to establish a connection
func GetClient() (*ethclient.Client, error) {
	return client, e
}
