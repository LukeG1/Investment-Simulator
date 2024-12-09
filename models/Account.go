package models

type AccountStatus struct {
	balance       float64
	yearDeposited float64
	yearWithdrawn float64
}

type Account interface {
	AllowedContribution() float64
	GetBalance() *AccountStatus
}

func (accountStatus *AccountStatus) deposit(amount float64) {
	accountStatus.balance += amount
	accountStatus.yearDeposited += amount
}

func Deposit(account Account, amount float64) {
	account.GetBalance().deposit(min(amount, account.AllowedContribution()))
}

// func withdrawal(account Account, amount float64) {
// 	// deal with taxable vs nontaxable accounts
// 	//account.GetBalance().Subtract(amount)
// }

// TODO: dont use the discrete distrubution, the only option here is to parrelize the completion of a year
// steps:
// -1. replace discrete distribution balances with regular balances, include year in/out
// 2. work out the simulation structure, as it is now a requirement to do this math
// 	2a. I would like for it to work agnostic of things like a household for now
// 	2b. ensure the year is parrellized
// -3. store that data in a new discrete distribution, this time with the goal of
// accumulating values and calculating their statistics at the end
// -3a. consider a way to end simulating prematurely if the simulation has 'stabalized'
// 3b. accumulate data into statistics and clear results
// 4. move on from sufficent data stop logic
