package main

import (
	"InvestmentSimulator/simulation"
	"InvestmentSimulator/statistics"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RunSimpleSimulation(precisionTarget float64, years int, startingBalance float64, investment string, additional float64) []statistics.OutcomeAggregator {
	return simulation.SimpleSimulation(precisionTarget, years, startingBalance, investment, additional)
}
