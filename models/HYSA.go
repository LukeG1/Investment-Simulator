package models

import (
	"math"
)

type HYSA struct {
	balances map[string]*Balance
}

// TODO: error handliing on invalid economicFactors

// allowedContribution implements Account.
func (hysa HYSA) AllowedContribution() float64 {
	// there is no limit to how much you can withdrawl from a high yield savings account
	return math.MaxFloat64
}

// withdrawal implements Account.
func (hysa *HYSA) withdrawal(economicFactor string, amount float64) {
	hysa.balances[economicFactor].total -= amount
	hysa.balances[economicFactor].yearWithdrawn += amount
}

// ------------  can be pretty much copy pasted to other accounts
// TODO: Find a nice way to generalize these

func newHYSA(economicFactors ...EconomicFactor) *HYSA {
	return &HYSA{
		NewBalanceMap(economicFactors...),
	}
}

// getBalance implements Account.
func (hysa *HYSA) GetBalance(economicFactor string) *Balance {
	return hysa.balances[economicFactor]
}

func (hysa *HYSA) Accrue() {
	for _, balance := range hysa.balances {
		balance.Accrue()
	}
}
