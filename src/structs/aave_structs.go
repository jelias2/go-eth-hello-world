package structs

import (
	"fmt"
	"github.com/jelias2/go-eth-hello-world/src/decimal"
	"math/big"
)

var (
	// No decimals in solditiy everything needs to be moved by 18 decimal places
	eighteen = decimal.NewDec(1e18)
)

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
