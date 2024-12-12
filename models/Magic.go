package models

import (
	"math"
)

type Magic struct {
	balances map[string]*Balance
}

// TODO: error handliing on invalid economicFactors

// allowedContribution implements Account.
func (magic Magic) AllowedContribution() float64 {
	// there is no limit to how much you can withdrawl from a high yield savings account
	return math.MaxFloat64
}

// withdrawal implements Account.
func (magic *Magic) withdrawal(economicFactor string, amount float64) {
	magic.balances[economicFactor].Total -= amount
	magic.balances[economicFactor].yearWithdrawn += amount
}

// ------------  can be pretty much copy pasted to other accounts
// TODO: Find a nice way to generalize these

func NewMagic(economicFactors ...*EconomicFactor) *Magic {
	return &Magic{
		NewBalanceMap(economicFactors...),
	}
}

// getBalance implements Account.
func (magic *Magic) GetBalance(economicFactor string) *Balance {
	return magic.balances[economicFactor]
}

func (magic *Magic) Accrue() {
	for _, balance := range magic.balances {
		balance.Accrue()
	}
}
