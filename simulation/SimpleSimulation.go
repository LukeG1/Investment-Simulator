package simulation

import (
	"InvestmentSimulator/models"
	"InvestmentSimulator/statistics"
	"time"
)

// TODO: multihread
func SimpleSimulation(results *SimulationResult, precisionTarget float64, years int, startingBalance float64, investment string, additional float64) {
	distributionLearners := make([]statistics.DistributionLearner, years)

	for year := 0; year < years; year++ {
		distributionLearners[year] = *statistics.NewDistributionLearner(precisionTarget)
		accountResults := *NewAccountResults("magic")
		results.YearlyResults[year] = accountResults
	}

	startTime := time.Now().Unix()

	// never run more than a 10 million sims for now
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

		results.TotalSims++
		results.SimulationDuration = time.Now().Unix() - startTime

		if sim%2_000 == 0 {
			stable := true
			for year := 0; year < years; year++ {
				summary := *distributionLearners[year].Summarize()
				results.YearlyResults[year].InvestmentResults[investment] = summary
				if !summary.Stable {
					stable = false
				}
			}
			if stable {
				break
			}
		}
	}

	for year := 0; year < years; year++ {
		summary := *distributionLearners[year].Summarize()
		results.YearlyResults[year].InvestmentResults[investment] = summary
	}
}
