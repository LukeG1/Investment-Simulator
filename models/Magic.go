package models

import "math"

type Magic struct {
	*AbstractAccount
}

// newHYSA creates a new instance of a High-Yield Savings Account (HYSA).
func NewMagic(economicFactors ...*EconomicFactor) *Magic {
	abstractAccount := &AbstractAccount{
		Investments: NewBalanceMap(economicFactors...),
	}
	a := &Magic{
		AbstractAccount: abstractAccount,
	}
	abstractAccount.Account = a
	return a
}

func (h *Magic) AllowedContribution() float64 {
	return math.MaxFloat64
}
