package simulation

import (
	"InvestmentSimulator/models"
	"InvestmentSimulator/statistics"
	"fmt"
	"time"
)

// TODO: this is really SimpleInvestmentSimulator
// use a magic account that can take any combination of investments with any starting value and any yearly contribution
// precision target: .01, .1, 1, 10, 100, 1000
// send request, let data stabalize, return each outcome accumulators values
// multihread
// would be nice to somehow send a periodic snapshot to the frontend of the stablization progress

func SimpleSimulation(results *SimulationResult, precisionTarget float64, years int, startingBalance float64, investment string, additional float64) []statistics.LearnedSummary {
	distributionLearners := make([]statistics.DistributionLearner, years)
	learnedSummaries := make([]statistics.LearnedSummary, years)
	for year := 0; year < years; year++ {
		distributionLearners[year] = *statistics.NewDistributionLearner(precisionTarget)
	}

	startTime := time.Now().Unix()

	// never run more than a billion sims for now0
	for sim := 0; sim < 10_000_000; sim++ {

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

			distributionLearners[year].AddOutcome(magicAccount.Investments[investment].Balance)

			switch investment {
			case "bonds":
				models.TBonds.Accrue()
			case "market":
				models.SandP500.Accrue()
			}

		}

		// TODO: consider moving this to the incremental checker
		results.TotalSims++
		results.SimulationDuration = time.Now().Unix() - startTime

		if sim%500 == 0 {
			stable := true
			for year := 0; year < years; year++ {
				learnedSummaries[year] = *distributionLearners[year].Summarize()
				if !learnedSummaries[year].Stable {
					stable = false
				}
			}
			if stable {
				break
			}
		}
	}

	for year := 0; year < years; year++ {
		learnedSummaries[year] = *distributionLearners[year].Summarize()
	}

	fmt.Println(learnedSummaries[years-1].Mean)

	return learnedSummaries
}
