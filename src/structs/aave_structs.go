package structs

import (
	"fmt"
	"github.com/jelias2/go-eth-hello-world/src/decimal"
	"math"
	"math/big"
	"sort"
	"strconv"
)

var (
	// No decimals in solditiy everything needs to be moved by 18 decimal places
	eighteen           = decimal.NewDec(1e18)
	secPerYear float64 = 31536000
)

type AssetAddress struct {
	Ticker  string
	Address string
}

type AaveAccount struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}

func NewAaveAccount(input struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}) AaveAccount {
	return AaveAccount{
		TotalCollateralETH:          input.TotalCollateralETH,
		TotalDebtETH:                input.TotalDebtETH,
		AvailableBorrowsETH:         input.AvailableBorrowsETH,
		CurrentLiquidationThreshold: input.CurrentLiquidationThreshold,
		Ltv:                         input.Ltv,
		HealthFactor:                input.HealthFactor,
	}
}

func (a *AaveAccount) LogAccountStruct() {

	fmt.Printf("\n-----\tAccount Dump\t-----\n\t TotalCollateralETH: %v \n\t TotalDebtETH: %v \n\t AvailableBorrowsETH: %v \n\t CurrentLiquidationThreshold: %v \n\t Loan To Value: %v \n\t HealthFactor: %v \n ---------------------------\n",
		decimal.NewDecFromBigInt(a.TotalCollateralETH).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(a.TotalDebtETH).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(a.AvailableBorrowsETH).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(a.CurrentLiquidationThreshold).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(a.Ltv).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(a.HealthFactor).Quo(eighteen).String(),
	)
}

type Rates struct {
	StableBorrowAPR   float64
	StableBorrowAPY   float64
	VariableBorrowAPR float64
	VariableBorrowAPY float64
	DepositAPR        float64
	DepositAPY        float64
}

type WrappedLendingPool struct {
	Ticker                  string
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
	URL                     string
	Rates                   Rates
}

func NewWrappedLendingPool(input struct {
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
}, ticker, address string) WrappedLendingPool {
	return WrappedLendingPool{
		Ticker:                  ticker,
		AvailableLiquidity:      input.AvailableLiquidity,
		TotalStableDebt:         input.TotalStableDebt,
		TotalVariableDebt:       input.TotalVariableDebt,
		LiquidityRate:           input.LiquidityRate,
		VariableBorrowRate:      input.VariableBorrowRate,
		StableBorrowRate:        input.StableBorrowRate,
		AverageStableBorrowRate: input.AverageStableBorrowRate,
		LiquidityIndex:          input.LiquidityIndex,
		VariableBorrowIndex:     input.VariableBorrowIndex,
		LastUpdateTimestamp:     input.LastUpdateTimestamp,
		URL:                     fmt.Sprintf("https://app.aave.com/#/reserve-overview/%s", address),
		Rates: Rates{
			StableBorrowAPR:   rayValueToPercentage(input.StableBorrowRate),
			StableBorrowAPY:   aprToAPY(input.StableBorrowRate),
			VariableBorrowAPR: rayValueToPercentage(input.VariableBorrowRate),
			VariableBorrowAPY: aprToAPY(input.VariableBorrowRate),
			DepositAPR:        rayValueToPercentage(input.LiquidityRate),
			DepositAPY:        aprToAPY(input.LiquidityRate),
		},
	}
}

// Create a list type of WrappedLendingPools with length and custom sorting function (by deposit APR) for that list
type WrappedLendingPoolList []WrappedLendingPool

func (wllpl WrappedLendingPoolList) Len() int      { return len(wllpl) }
func (wllpl WrappedLendingPoolList) Swap(i, j int) { wllpl[i], wllpl[j] = wllpl[j], wllpl[i] }
func (wllpl WrappedLendingPoolList) Less(i, j int) bool {
	return wllpl[i].Rates.DepositAPR < wllpl[j].Rates.DepositAPR
}

// SortedDepositRates will take in the LP map and output a list ordered by the best deposit intrest rates
func SortedDepositRates(lpMap map[string]WrappedLendingPool) {

	var wllpList []WrappedLendingPool

	for _, val := range lpMap {
		wllpList = append(wllpList, val)
	}

	sort.Sort(sort.Reverse(WrappedLendingPoolList(wllpList)))
	fmt.Println("\n----- Sorted Deposit Rates -----")
	for _, lp := range wllpList {
		fmt.Printf("\tTicker: %s DepositAPR: %.3f \n", lp.Ticker, lp.Rates.DepositAPR)
	}
	fmt.Println("\n-------------------------------")
}

func (w *WrappedLendingPool) LogLendingPoolStruct() {

	fmt.Printf("\n-----\t %s Reserve Data Dump\t-----\n\t"+
		"AvailableLiquidity: %v \n\t"+
		"TotalStableDebt: %v \n\t"+
		"TotalVariableDebt: %v \n\t"+
		"LiquidityRate: %v \n\t"+
		"AverageStableBorrowRate: %v \n\n\t"+
		"Stable Borrow APY: %.3f \n\t"+
		"Stable Borrow APR: %.3f \n\n\t"+
		"Variable Borrow APY: %.3f \n\t"+
		"Variable Borrow APR: %.3f \n\n\t"+
		"Deposit APY: %.3f \n\t"+
		"Deposit APR: %.3f \n\n\t"+
		"LastUpdateTimestamp: %v \n\t"+
		"URL: %s \n"+
		"---------------------------\n",
		w.Ticker,
		w.AvailableLiquidity,
		decimal.NewDecFromBigInt(w.TotalStableDebt).Quo(eighteen).String(),
		decimal.NewDecFromBigInt(w.TotalVariableDebt).Quo(eighteen).String(),
		w.LiquidityRate,
		w.AverageStableBorrowRate,
		w.Rates.StableBorrowAPY,
		w.Rates.StableBorrowAPR,
		w.Rates.VariableBorrowAPY,
		w.Rates.VariableBorrowAPR,
		w.Rates.DepositAPY,
		w.Rates.DepositAPR,
		decimal.NewDecFromBigInt(w.LastUpdateTimestamp).Quo(eighteen).String(),
		w.URL,
	)
}

func rayValueToPercentage(rayValue *big.Int) float64 {
	var ray big.Int
	// This ray is only 25 positions, 27 give decimals APR, 25 give percentage -> 0.02 = 2%
	ray.SetString("10000000000000000000000000", 10)
	twentyfive := decimal.NewDecFromBigInt(&ray)
	return stringPercentageToFloat(decimal.NewDecFromBigInt(rayValue).Quo(twentyfive).String())
}

// Yes ugly, first get it to work, then get it pretty, premature optimzation is the root of all evil
func stringPercentageToFloat(percentage string) float64 {
	val, err := strconv.ParseFloat(percentage, 64)
	if err != nil {
		fmt.Printf("Error converting string percentage to float: %s \n", err.Error())
	}
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
	apr := rayValueToPercentage(inputAPR) / 100
	apy := math.Pow((1+(apr/secPerYear)), secPerYear) - 1
	return apy * 100
}
