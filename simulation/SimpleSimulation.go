package simulation

import (
	"InvestmentSimulator/models"
	"InvestmentSimulator/statistics"
)

// TODO: this is really SimpleInvestmentSimulator
// use a magic account that can take any combination of investments with any starting value and any yearly contribution
// precision target: .01, .1, 1, 10, 100, 1000
// send request, let data stabalize, return each outcome accumulators values
// multihread
// would be nice to somehow send a periodic snapshot to the frontend of the stablization progress

// TODO: refactor to simple simulator
func SimpleSimulation(precisionTarget float64, years int, startingBalance float64, investment string, additional float64) []statistics.OutcomeAggregator {
	outcomeAggregators := make([]statistics.OutcomeAggregator, years)
	for year := 0; year < years; year++ {
		outcomeAggregators[year] = *statistics.NewOutcomeAggregator(10, 500)
	}

	// never run more than a billion sims for now0
	for sim := 0; sim < 10_000; sim++ {

		magicAccount := models.NewMagic(&models.SandP500)
		switch investment {
		case "bonds":
			magicAccount = models.NewMagic(&models.TBonds)
		case "market":
			magicAccount = models.NewMagic(&models.SandP500)
		}
		magicAccount.Deposit(investment, startingBalance)
		for year := 0; year < years; year++ {

			magicAccount.Deposit(investment, additional)

			magicAccount.Accrue()

			outcomeAggregators[year].AddOutcome(magicAccount.Balances[investment].Total)

			switch investment {
			case "bonds":
				models.TBonds.Accrue()
			case "market":
				models.SandP500.Accrue()
			}

		}
		// if sim%10 == 0 {
		// 	stable := true
		// 	for year := 0; year < years; year++ {
		// 		if !outcomeAggregators[year].Stable {
		// 			stable = false
		// 		}
		// 	}
		// 	if stable {
		// 		fmt.Println(sim)
		// 		break
		// 	}
		// }
	}

	return outcomeAggregators
}
