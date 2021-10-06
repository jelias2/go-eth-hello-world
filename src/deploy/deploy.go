package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"

	"github.com/jelias2/go-eth-hello-world/src/store"
)

func main() {

	endpoint := os.Getenv("RINKEBY_ENDPOINT")
	// pubKey := os.Getenv("RINKEBY_PUB_KEY")
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

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployed to address: https://rinkeby.etherscan.io/address/%v \n", address.Hex())
	fmt.Printf("Deployed at transaction https://rinkeby.etherscan.io/tx/%v \n", tx.Hash().Hex())

	_ = instance
}
