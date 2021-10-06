package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jelias2/go-eth-hello-world/src/store"
)

func main() {

	endpoint := os.Getenv("RINKEBY_ENDPOINT")
	privKey := os.Getenv("RINKEBY_PRIVATE_KEY")

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// Get Wallet Nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction nonce: %d \n", nonce)

	// Get estimated gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Suggested gasPrice: %d \n", gasPrice)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(10000000) // in units
	auth.GasPrice = gasPrice

	// Address of the contract we wish to query

	contract_addr := "0x5C12097027a35803f8f75B8c1434F092841ca4B2"
	address := common.HexToAddress(contract_addr)
	// using the client and the address create an instance of the contact with in our code
	store_instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")

	// Create a 32 byte array
	key := [32]byte{}
	value := [32]byte{}

	// Copy the strings into the sized bytes arrays
	copy(key[:], []byte("Jacobs"))
	copy(value[:], []byte("1st-contract"))

	// Set the items array within the contract
	fmt.Println("Sending transaction")
	// send with auth bind options
	tx, err := store_instance.SetItem(auth, key, value)

	// Print the transaction hash
	fmt.Printf("Sent transaction https://rinkeby.etherscan.io/tx/%v \n", tx.Hash().Hex())

	fmt.Printf("Sleeping while transaction procresses \n")
	for i := 0; i < 30; i++ {
		time.Sleep(1 * time.Second)
		if i%10 == 0 {
			fmt.Printf("%d seconds elapsed...\n", i)
		}
	}
	fmt.Printf("Reading from contract %s \n", contract_addr)
	// Sleep for a little bit let the the transcaction process
	// Read the items array within the contract
	result, err := store_instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Reading contract at address: %s \n", contract_addr)
	fmt.Printf("Value of %s: %s \n", string(key[:]), string(result[:])) // 1st-contract
	fmt.Println("Done :)")

}
