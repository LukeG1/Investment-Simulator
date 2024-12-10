package models

type Balance struct {
	total          float64
	yearDeposited  float64
	yearWithdrawn  float64
	economicFactor *EconomicFactor
}

func NewBalanceMap(economicFactors ...EconomicFactor) map[string]*Balance {
	balances := make(map[string]*Balance)
	for _, economicFactor := range economicFactors {
		balances[economicFactor.Name] = &Balance{
			economicFactor: &economicFactor,
		}
	}
	return balances
}

func (balance *Balance) Accrue() {
	balance.total *= 1 + balance.economicFactor.rate - Inflation.rate
	balance.yearDeposited = 0
	balance.yearWithdrawn = 0
}

type Account interface {
	AllowedContribution() float64
	GetBalance(string) *Balance
	Withwithdrawal(string, float64)
}

func (accountStatus *Balance) deposit(amount float64) {
	accountStatus.total += amount
	accountStatus.yearDeposited += amount
}

func Deposit(account Account, economicFactor string, amount float64) {
	account.GetBalance(economicFactor).deposit(min(amount, account.AllowedContribution()))
}
