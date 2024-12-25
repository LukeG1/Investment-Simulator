package simulation

import (
	"InvestmentSimulator/statistics"
)

type SimulationResult struct {
	YearlyResults      []AccountResults `json:"YearlyResults"`
	TotalSims          int              `json:"TotalSims"`
	SimulationDuration int64            `json:"SimulationDuration"`
	// Cancel             bool             `json:"Cancel"`
	// more metadata needed?
	// TODO: consider a total accumulator
}

type AccountResults struct {
	Name string
	// wails wont let me make this a pointer even though I think it should be
	InvestmentResults map[string]statistics.LearnedSummary `json:"InvestmentResults"`
	// TODO: consider a total accumulator
}

func NewSimulationResult(years int) *SimulationResult {
	return &SimulationResult{
		YearlyResults: make([]AccountResults, years),
	}
}

func NewAccountResults(name string) *AccountResults {
	return &AccountResults{
		Name:              name,
		InvestmentResults: make(map[string]statistics.LearnedSummary),
	}
}

func (sr *SimulationResult) ToCSV() string {
	// year | account | investment | stats...
	return ""
}

func (sr *SimulationResult) ExportCSV(exportPath string) {
	// year | account | investment | stats...
}

// next steps:
// unit test existing code
// cancel button
// export csv
// on end of sim or cancel send the result.csv file to the frontend for possible download
