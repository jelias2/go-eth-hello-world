package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jelias2/go-eth-hello-world/src/store"
	"log"
	"os"
)

func main() {

	endpoint := os.Getenv("RINKEBY_ENDPOINT")

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Address of the contract we wish to query
	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	// using the client and the address create an instance of the contact with in our code
	store_instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")

	// Call the contract (*store.StoreCaller).Version(opts *bind.CallOpts)
	// CallOpts is the collection of options to fine tune a contract call request, such as pending, from, and blocknumber
	version, err := store_instance.Version(nil)
	if err != nil {
		log.Print("Error calling instance contract!")
		log.Fatal(err)
	}

	fmt.Printf("Contract version: %v \n", version)

}
