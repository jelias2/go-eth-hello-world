package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"

	"github.com/jelias2/go-eth-hello-world/src/decimal"
	"github.com/jelias2/go-eth-hello-world/src/external-contracts/aave/ILendingPool"
	"github.com/jelias2/go-eth-hello-world/src/external-contracts/aave/ILendingPoolAddressProvider"
)

func logAccountStruct(input struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}) {

	// No decimals in solditiy everything needs to be moved by 18 decimal places
	eighteen := decimal.NewDec(1e18)

	fmt.Printf("\n-----\tAccount Dump\t-----\n\t TotalCollateralETH: %v \n\t TotalDebtETH: %v \n\t AvailableBorrowsETH: %v \n\t CurrentLiquidationThreshold: %v \n\t Loan To Value: %v \n\t HealthFactor: %v \n ---------------------------\n",
		decimal.NewDecFromBigInt(input.TotalCollateralETH).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(input.TotalDebtETH).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(input.AvailableBorrowsETH).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(input.CurrentLiquidationThreshold).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(input.Ltv).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(input.HealthFactor).Quo(eighteen).String(),
	)
}

func main() {

	endpoint := os.Getenv("MAINNET_ENDPOINT")
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
	contract_addr := "0xB53C1a33016B2DC2fF3653530bfF1848a515c8c5"
	address := common.HexToAddress(contract_addr)

	// using the client and the address create an instance of the contact with in our code
	aave_instance, err := ILendingPoolAddressProvider.NewAave(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get the Address of a LendingPool
	lendingPoolAddress, err := aave_instance.GetLendingPool(nil)
	if err != nil {
		log.Print("Error calling GetLendingPool contract function!")
		log.Fatal(err)
	}
	fmt.Printf("Current LendingPool address is: %v \n", lendingPoolAddress)

	// Create a read-only caller to the aave lending pool
	lendingPoolCaller, err := ILendingPool.NewAaveCaller(lendingPoolAddress, client)
	if err != nil {
		log.Print("Error Creating LendingPoolCaller Struct")
		log.Fatal(err)
	}

	account_struct, err := lendingPoolCaller.GetUserAccountData(nil, aave_account_addr)
	if err != nil {
		log.Print("Error creating account_struct!")
		log.Fatal(err)
	}
	// log the account data
	logAccountStruct(account_struct)

	marketID, err := aave_instance.GetMarketId(nil)
	if err != nil {
		log.Print("Error calling getMarketID!")
		log.Fatal(err)
	}

	fmt.Printf("Current MarketID: %v \n", marketID)

	// 	marketID, err := aave_instance.G(nil)
	//   if err != nil {
	//     log.Print("Error calling TotalBorrows!")
	//     log.Fatal(err)
	//   }

	//   fmt.Printf("Current Compound is CDai Total Supply: %v \n", TotalBorrows)

}
