package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jelias2/go-eth-hello-world/src/external-contracts/aave/AaveProtocolDataProvider"
	"github.com/jelias2/go-eth-hello-world/src/external-contracts/aave/ILendingPool"
	"github.com/jelias2/go-eth-hello-world/src/external-contracts/aave/ILendingPoolAddressProvider"
	"github.com/jelias2/go-eth-hello-world/src/structs"
	"log"
	"os"
)

const (
	LendingPoolAddressProviderAddress = "0xB53C1a33016B2DC2fF3653530bfF1848a515c8c5"
	USDT_MAINNET_ADDRESS              = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	TUSD_MAINNET_ASSET_ADDRESS        = "0x0000000000085d4780B73119b644AE5ecd22b376"
	DAI_MAINNET_ASSET_ADDRESS         = "0x6b175474e89094c44da98b954eedeac495271d0f"
	USDC_MAINNET_ASSET_ADDRESS        = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
)

func main() {

	fmt.Println("Starting load aave program!")
	endpoint := os.Getenv("MAINNET_ENDPOINT")

	if endpoint == "" {
		fmt.Println("MAINNET_ENDPOINT environment variable is empty, please source and try again")
		fmt.Println("exiting...")
		os.Exit(0)
	}

	var addressMap = map[string]string{
		"USDT": USDT_MAINNET_ADDRESS,
		"TUSD": TUSD_MAINNET_ASSET_ADDRESS,
		"DAI":  DAI_MAINNET_ASSET_ADDRESS,
		"USDC": USDC_MAINNET_ASSET_ADDRESS,
	}

	lendingPoolMap := make(map[string]structs.WrappedLendingPool)

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Load an Address which is using Aave
	my_aave_account_addr := os.Getenv("AAVE_PUBLIC_ADDRESS")
	fmt.Println()
	fmt.Println("Starting.....")
	fmt.Println("Using Aave Public Address: ", my_aave_account_addr)
	aave_account_addr := common.HexToAddress(my_aave_account_addr)

	// Address of the Mainnet LendingPoolAddressProvider contract we wish to query
	lending_pool_address := common.HexToAddress(LendingPoolAddressProviderAddress)

	// using the client and the address create an instance of the contact with in our code
	aave_instance, err := ILendingPoolAddressProvider.NewAave(lending_pool_address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get the Address of a LendingPool
	lendingPoolAddress, err := aave_instance.GetLendingPool(nil)
	if err != nil {
		log.Print("Error calling GetLendingPool contract function!")
		log.Fatal(err)
	}
	price_oracle, err := aave_instance.GetPriceOracle(nil)
	if err != nil {
		log.Print("Error calling GetPriceOracle contract function!")
		log.Fatal(err)
	}

	fmt.Printf("Current LendingPool address is: %v \n", lendingPoolAddress)
	fmt.Printf("Current Price Oracle address is: %v \n", price_oracle)

	// Create a read-only caller to the aave lending pool
	lendingPoolCaller, err := ILendingPool.NewAaveCaller(lendingPoolAddress, client)
	if err != nil {
		log.Print("Error Creating LendingPoolCaller Struct")
		log.Fatal(err)
	}

	rawAccountStruct, err := lendingPoolCaller.GetUserAccountData(nil, aave_account_addr)
	if err != nil {
		log.Print("Error creating account_struct!")
		log.Fatal(err)
	}
	aaveAccount := structs.NewAaveAccount(rawAccountStruct)

	// log the account data
	aaveAccount.LogAccountStruct()

	// addressProvider, err := lendingPoolCaller.GetAddressesProvider(nil)
	// if err != nil {
	// 	log.Print("Error getting LendingPool Address!")
	// 	log.Fatal(err)
	// }
	//fmt.Printf("Address Provider: %v \n", addressProvider)

	//Create the Aave Protocol Data Provider, not sure if the aave_account_address is the proper one to pass in
	aave_data_provider_caller, err := AaveProtocolDataProvider.NewAaveProtocolDataProviderCaller(lendingPoolAddress, client)
	if err != nil {
		log.Print("Error creating aave_data_provider_caller_object!")
		log.Fatal(err)
	}

	for ticker, addr := range addressMap {
		rawReserveData, err := aave_data_provider_caller.GetReserveData(nil, common.HexToAddress(addr))
		if err != nil {
			log.Print("Error getting token data!")
			log.Fatal(err)
		}
		wrappedLendingPool := structs.NewWrappedLendingPool(rawReserveData, ticker, addr)
		lendingPoolMap[ticker] = wrappedLendingPool
	}
	structs.SortedDepositRates(lendingPoolMap)
}
