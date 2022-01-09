package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jelias2/go-eth-hello-world/src/decimal"
	"github.com/jelias2/go-eth-hello-world/src/external-contracts/aave/AaveProtocolDataProvider"
	"github.com/jelias2/go-eth-hello-world/src/external-contracts/aave/ILendingPool"
	"github.com/jelias2/go-eth-hello-world/src/external-contracts/aave/ILendingPoolAddressProvider"
	"github.com/jelias2/go-eth-hello-world/src/structs"
)

const (
	secPerYear                        float64 = 31536000
	LendingPoolAddressProviderAddress         = "0xB53C1a33016B2DC2fF3653530bfF1848a515c8c5"
	USDT_MAINNET_ADDRESS                      = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	TUSD_MAINNET_ASSET_ADDRESS                = "0x0000000000085d4780B73119b644AE5ecd22b376"
)

func logAccountStruct(input structs.AaveAccount) {
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

func rayValueToPercentage(rayValue *big.Int) string {
	var ray big.Int
	// This ray is only 25 positions, 27 give decimals APR, 25 give percentage -> 0.02 = 2%
	ray.SetString("10000000000000000000000000", 10)
	twentyfive := decimal.NewDecFromBigInt(&ray)
	return decimal.NewDecFromBigInt(rayValue).Quo(twentyfive).String()
}

// Yes ugly, first get it to work, then get it pretty, premature optimzation is the root of all evil
func stringPercentageToFloat(percentage string) float64 {
	val, err := strconv.ParseFloat(percentage, 64)
	if err != nil {
		fmt.Printf("Error converting string percentage to float: %s \n", err.Error())
	}
	//fmt.Printf("Succuessfully converted %s -> %f \n", percentage, val)
	return val
}

/*
https://docs.aave.com/developers/guides/apy-and-apr

  depositAPR = liquidityRate/RAY
  variableBorrowAPR = variableBorrowRate/RAY
  stableBorrowAPR = variableBorrowRate/ray

  depositAPY = ((1 + (depositAPR / SECONDS_PER_YEAR)) ^ SECONDS_PER_YEAR) - 1
  variableBorrowAPY = ((1 + (variableBorrowAPR / SECONDS_PER_YEAR)) ^ SECONDS_PER_YEAR) - 1
  stableBorrowAPY = ((1 + (stableBorrowAPR / SECONDS_PER_YEAR)) ^ SECONDS_PER_YEAR) - 1

  APY = ((1 + (APR / SECONDS_PER_YEAR)) ^ SECONDS_PER_YEAR) - 1
*/

func aprToAPY(inputAPR *big.Int) float64 {
	aprString := rayValueToPercentage(inputAPR)
	aprFloat := stringPercentageToFloat(aprString) / 100
	apy := math.Pow((1+(aprFloat/secPerYear)), secPerYear) - 1
	return apy

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
	eighteen := decimal.NewDec(1e18)
	//RAY := math.Pow(10, 27) // 10 to the power 27

	fmt.Printf("\n-----\tReserve Data Dump\t-----\n\t"+
		"AvailableLiquidity: %v \n\t"+
		"TotalStableDebt: %v \n\t"+
		"TotalVariableDebt: %v \n\t"+
		"LiquidityRate: %v \n\t"+
		"AverageStableBorrowRate: %v \n\n\t"+
		"Stable Borrow APR: %v \n\t"+
		"Stable Borrow APY: %.2f \n\n\t"+
		"Variable Borrow APR: %v \n\t"+
		"Variable Borrow APY: %.2f \n\n\t"+
		"Deposit APR: %v \n\t"+
		"Deposit APY: %.2f \n\t"+
		"LastUpdateTimestamp: %v \n"+
		"---------------------------\n",
		input.AvailableLiquidity,
		decimal.NewDecFromBigInt(input.TotalStableDebt).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(input.TotalVariableDebt).Quo(eighteen).String(),
		input.LiquidityRate,
		input.AverageStableBorrowRate,
		rayValueToPercentage(input.StableBorrowRate),
		aprToAPY(input.StableBorrowRate)*100,
		rayValueToPercentage(input.VariableBorrowRate),
		aprToAPY(input.VariableBorrowRate)*100,
		rayValueToPercentage(input.LiquidityRate),
		aprToAPY(input.LiquidityRate)*100,
		decimal.NewDecFromBigInt(input.LastUpdateTimestamp).Quo(eighteen).String(),
	)
}

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

	rawAccountStruct, err := lendingPoolCaller.GetUserAccountData(nil, aave_account_addr)
	if err != nil {
		log.Print("Error creating account_struct!")
		log.Fatal(err)
	}
	aaveAccount := structs.NewAaveAccount(rawAccountStruct)

	// log the account data
	aaveAccount.LogAccountStruct()

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
	reserve_data, err := aave_data_provider_caller.GetReserveData(nil, common.HexToAddress(TUSD_MAINNET_ASSET_ADDRESS))
	if err != nil {
		log.Print("Error getting token data!")
		log.Fatal(err)
	}

	logLendingPoolStruct(reserve_data)
}

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
