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
//     ...,
// ]

type SimulationResult struct {
	YearlyResults      []AccountResults `json:"YearlyResults"`
	TotalSims          int              `json:"TotalSims"`
	SimulationDuration int64            `json:"SimulationDuration"`
	// more metadata needed?
}

type AccountResults struct {
	// wails wont let me make this a pointer even though I think it should be
	InvestmentResults map[string]statistics.LearnedSummary `json:"InvestmentResults"`
}

func NewSimulationResult() *SimulationResult {
	return &SimulationResult{
		YearlyResults: []AccountResults{},
	}
}

func (sr *SimulationResult) ExportCSV(exportPath string) {
	// year | account | investment | stats...
}
