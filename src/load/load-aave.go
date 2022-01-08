package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/jelias2/go-eth-hello-world/src/decimal"
	"github.com/jelias2/go-eth-hello-world/src/external-contracts/aave/AaveProtocolDataProvider"
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

func logLendingPoolStruct(input struct {
	AvailableLiquidity      *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}) {
	// https://docs.aave.com/developers/the-core-protocol/lendingpool#getreservedata
	// data is returned in ray units = 10^27
	// No decimals in solditiy everything needs to be moved by 18 decimal places
	six := decimal.NewDec(1e6)
	//RAY := math.Pow(10, 27) // 10 to the power 27
	RAY := new(big.Float).SetUint64(uint64(math.Pow(10, 25)))

	f := new(big.Float).SetInt(input.VariableBorrowRate)

	fmt.Printf("\n-----\tReserve Data Dump\t-----\n\t"+
		"AvailableLiquidity: %v \n\t"+
		"TotalStableDebt: %v \n\t"+
		"TotalVariableDebt: %v \n\t"+
		"LiquidityRate: %v \n\t"+
		"StableBorrowRate: %v \n\t"+
		"AverageStableBorrowRate: %v \n\t"+
		"VariableBorrowRate: %v \n\t"+
		"LastUpdateTimestamp: %v \n"+
		"---------------------------\n",
		decimal.NewDecFromBigInt(input.AvailableLiquidity).Quo(six).String(),
		decimal.NewDecFromBigInt(input.TotalStableDebt).Quo(six).String(),
		decimal.NewDecFromBigInt(input.TotalVariableDebt).Quo(six).String(),
		decimal.NewDecFromBigInt(input.LiquidityRate).Quo(six).String(),
		decimal.NewDecFromBigInt(input.StableBorrowRate).Quo(six).String(),
		decimal.NewDecFromBigInt(input.AverageStableBorrowRate).Quo(six).String(),
		//decimal.NewDecFromBigInt(input.VariableBorrowRate).Quo(six).String(),
		new(big.Float).Quo(f, RAY),
		decimal.NewDecFromBigInt(input.LastUpdateTimestamp).Quo(six).String(),
	)
}

const (
	LendingPoolAddressProviderAddress = "0xB53C1a33016B2DC2fF3653530bfF1848a515c8c5"
	USDT_MAINNET_ADDRESS              = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
)

func main() {

	fmt.Println("Starting!")
	endpoint := os.Getenv("MAINNET_ENDPOINT")

	if endpoint == "" {
		fmt.Println("MAINNET_ENDPOINT environment variable is empty, please source and try again")
		fmt.Println("exiting...")
		os.Exit(0)
	}

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

	account_struct, err := lendingPoolCaller.GetUserAccountData(nil, aave_account_addr)
	if err != nil {
		log.Print("Error creating account_struct!")
		log.Fatal(err)
	}
	// log the account data
	logAccountStruct(account_struct)

	addressProvider, err := lendingPoolCaller.GetAddressesProvider(nil)
	if err != nil {
		log.Print("Error getting LendingPool Address!")
		log.Fatal(err)
	}
	fmt.Printf("Address Provider: %v \n", addressProvider)

	//Create the Aave Protocol Data Provider, not sure if the aave_account_address is the proper one to pass in
	aave_data_provider_caller, err := AaveProtocolDataProvider.NewAaveProtocolDataProviderCaller(lendingPoolAddress, client)
	if err != nil {
		log.Print("Error creating aave_data_provider_caller_object!")
		log.Fatal(err)
	}

	// https://docs.aave.com/developers/the-core-protocol/lendingpool#getreservedata
	// data is returned in ray units = 10^27
	reserve_data, err := aave_data_provider_caller.GetReserveData(nil, common.HexToAddress(USDT_MAINNET_ADDRESS))
	if err != nil {
		log.Print("Error getting token data!")
		log.Fatal(err)
	}

	logLendingPoolStruct(reserve_data)

	// 	for index, token := range token_data {

	// 		fmt.Print("\n %d. Name: %s Address: %s", index, token.Symbol, token.TokenAddress)
	// 	}

	// 	marketID, err := aave_instance.GetMarketId(nil)
	// 	if err != nil {
	// 		log.Print("Error calling getMarketID!")
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("Current MarketID: %v \n", marketID)

	// 	marketID, err := aave_instance.G(nil)
	//   if err != nil {
	//     log.Print("Error calling TotalBorrows!")
	//     log.Fatal(err)
	//   }

	//   fmt.Printf("Current Compound is CDai Total Supply: %v \n", TotalBorrows)

}
