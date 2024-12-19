package models

import (
	historicdata "InvestmentSimulator/historicData"
	"InvestmentSimulator/statistics"
)

type EconomicFactor struct {
	Name    string
	Sampler statistics.NaieveDataSampler
	Rate    float64
}

func NewEconomicFactor(name string, data *[]float64) *EconomicFactor {
	sampler := statistics.GenerateKernelSampler(data)
	return &EconomicFactor{
		name,
		sampler,
		sampler.Sample(),
	}
}

func (economicFactor *EconomicFactor) Accrue() {
	economicFactor.Rate = economicFactor.Sampler.Sample()
}

var SandP500 EconomicFactor = *NewEconomicFactor(
	"market",
	historicdata.RawSandP500,
)

var Inflation EconomicFactor = *NewEconomicFactor(
	"inflation",
	historicdata.RawInflation,
)

var TBonds EconomicFactor = *NewEconomicFactor(
	"bonds",
	historicdata.RawTBonds,
)
