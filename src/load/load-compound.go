package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jelias2/go-eth-hello-world/src/external-contracts/compound"
)

func main() {

	endpoint := os.Getenv("RINKEBY_ENDPOINT")

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Address of the contract we wish to query
	contract_addr := "0x6d7f0754ffeb405d23c51ce938289d4835be3b14"
	address := common.HexToAddress(contract_addr)
	fmt.Printf("contract is loaded from %s \n", contract_addr)

	// using the client and the address create an instance of the contact with in our code
	//store_instance, err := store.NewStore(address, client)
	compound_instance, err := compound.NewCompound(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Call the contract (*store.StoreCaller)Version(opts *bind.CallOpts)
	// CallOpts is the collection of options to fine tune a contract call request, such as pending, from, and blocknumber
	is_Ctoken, err := compound_instance.IsCToken(nil)
	if err != nil {
		log.Print("Error calling isCToken contract function!")
		log.Fatal(err)
	}

	fmt.Printf("Current Compound is Ctoken instance: %v \n", is_Ctoken)

	totalSupply, err := compound_instance.TotalSupply(nil)
	if err != nil {
		log.Print("Error calling TotalSupply!")
		log.Fatal(err)
	}

	fmt.Printf("Current Compound is CDai Total Supply: %v \n", totalSupply)

	TotalBorrows, err := compound_instance.TotalBorrows(nil)
	if err != nil {
		log.Print("Error calling TotalBorrows!")
		log.Fatal(err)
	}

	fmt.Printf("Current Compound is CDai Total Supply: %v \n", TotalBorrows)

}
