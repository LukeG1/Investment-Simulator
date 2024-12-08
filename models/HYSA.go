package models

import (
	"InvestmentSimulator/statistics"
)

type HYSA struct {
	balance       statistics.DiscreteDistribtuion
	yearDeposited float64
	yearWithdrawn float64
}

// getBalance implements Account.
func (hysa HYSA) GetBalance() *statistics.DiscreteDistribtuion {
	return &hysa.balance
}

// allowedContribution implements Account.
func (hysa HYSA) AllowedContribution() float64 {
	return 10_000.0
}
