package models

type Investment struct {
	Balance        float64
	yearDeposited  float64
	yearWithdrawn  float64
	economicFactor *EconomicFactor
	// TODO: eventually I could link an outcome aggregator at construction that would addOutcome on accrue, could be cleaner
}

func (investment *Investment) deposit(amount float64) {
	investment.Balance += amount
	investment.yearDeposited += amount
}

func (investment *Investment) accrue() {
	investment.Balance *= 1 + investment.economicFactor.Rate - Inflation.Rate
	investment.yearDeposited = 0
	investment.yearWithdrawn = 0
}

// Abstract account implemented like so: https://stackoverflow.com/questions/30261032/how-to-implement-an-abstract-class-in-go
type Account interface {
	AllowedContribution() float64
}

type AbstractAccount struct {
	Account
	Investments map[string]*Investment
}

func NewBalanceMap(economicFactors ...*EconomicFactor) map[string]*Investment {
	balances := make(map[string]*Investment)
	for _, economicFactor := range economicFactors {
		balances[economicFactor.Name] = &Investment{
			economicFactor: economicFactor,
		}
	}
	return balances
}

func (a *AbstractAccount) Deposit(economicFactor string, amount float64) {
	a.Investments[economicFactor].deposit(min(amount, a.AllowedContribution()))
}

func (a *AbstractAccount) Accrue() {
	for _, balance := range a.Investments {
		balance.accrue()
	}
}
