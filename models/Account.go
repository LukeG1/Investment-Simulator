package models

import (
	"InvestmentSimulator/statistics"
)

// TODO: Something somewhere should probably be called an investment
type Balance struct {
	Total          float64
	yearDeposited  float64
	yearWithdrawn  float64
	economicFactor *EconomicFactor
	accumulator    *statistics.OutcomeAggregator
}

func (balance *Balance) deposit(amount float64) {
	balance.Total += amount
	balance.yearDeposited += amount
}

func (balance *Balance) accrue() {
	balance.Total *= 1 + balance.economicFactor.Rate - Inflation.Rate
	balance.yearDeposited = 0
	balance.yearWithdrawn = 0
}

// Abstract account implemented like here: https://stackoverflow.com/questions/30261032/how-to-implement-an-abstract-class-in-go
type Account interface {
	AllowedContribution() float64
}

type AbstractAccount struct {
	Account
	Balances map[string]*Balance
}

func NewBalanceMap(economicFactors ...*EconomicFactor) map[string]*Balance {
	balances := make(map[string]*Balance)
	for _, economicFactor := range economicFactors {
		balances[economicFactor.Name] = &Balance{
			economicFactor: economicFactor,
		}
	}
	return balances
}

func (a *AbstractAccount) Deposit(economicFactor string, amount float64) {
	a.Balances[economicFactor].deposit(min(amount, a.AllowedContribution()))
}

func (a *AbstractAccount) Accrue() {
	for _, balance := range a.Balances {
		balance.accrue()
	}
}
