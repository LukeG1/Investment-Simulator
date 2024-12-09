package models

import (
	"math"
)

type HYSA struct {
	balance AccountStatus
}

// getBalance implements Account.
func (hysa HYSA) GetBalance() *AccountStatus {
	return &hysa.balance
}

// allowedContribution implements Account.
func (hysa HYSA) AllowedContribution() float64 {
	return math.MaxFloat64
}
