package models

import (
	"InvestmentSimulator/statistics"
)

type Account interface {
	AllowedContribution() float64
	GetBalance() *statistics.DiscreteDistribtuion
}

func deposit(account Account, amount float64) {
	account.GetBalance().Add(min(amount, account.AllowedContribution()))
}
