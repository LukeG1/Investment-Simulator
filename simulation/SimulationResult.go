package simulation

import (
	"InvestmentSimulator/statistics"
)

// years = [
//     accounts = [
//         account = {
//             investment: summary
//             ...,
//         }
//         ...,
//     ]
//
//     ...,
// ]

type SimulationResults struct {
	YearlyResults      []AccountResults `json:"YearlyResults"`
	TotalSims          int              `json:"TotalSims"`
	SimulationDuration float64          `json:"SimulationDuration"`
	// more metadata needed?
}

type AccountResults struct {
	InvestmentResults map[string]*statistics.LearnedSummary `InvestmentResults:"InvestmentResults"`
}

func (sr *SimulationResults) toCSV(exportPath string) {
	// year | account | investment | stats...
}
