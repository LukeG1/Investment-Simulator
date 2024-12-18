package models

import (
	"fmt"
	"testing"
)

// Helper function to create a test EconomicFactor
func newTestEconomicFactor(name string, rate float64) *EconomicFactor {
	return &EconomicFactor{
		Name: name,
		Rate: rate,
	}
}

func TestNewMagic(t *testing.T) {
	magicAccount := NewMagic(&SandP500)

	magicAccount.Deposit("market", 100)
	magicAccount.Accrue()

	fmt.Println(magicAccount.Balances["market"].Total)

	// Assert
	// if magicAccount == nil {
	// 	t.Fatalf("Expected new Magic account to be non-nil")
	// }
	// if len(magicAccount.Balances) != 1 {
	// 	t.Fatalf("Expected 1 balance for economic factor but got %d", len(magicAccount.Balances))
	// }
	// if magicAccount.Balances["TestEconomy"].economicFactor != economicFactor {
	// 	t.Fatalf("Expected economic factor to be %v but got %v", economicFactor, magicAccount.Balances["TestEconomy"].economicFactor)
	// }
}

// func TestAllowedContribution(t *testing.T) {
// 	// Arrange
// 	economicFactor := newTestEconomicFactor("TestEconomy", 0.05)
// 	magicAccount := newMagic(economicFactor)

// 	// Act
// 	allowedContribution := magicAccount.AllowedContribution()

// 	// Assert
// 	expected := 10000.0
// 	if allowedContribution != expected {
// 		t.Fatalf("Expected AllowedContribution to be %f but got %f", expected, allowedContribution)
// 	}
// }

// func TestDeposit(t *testing.T) {
// 	// Arrange
// 	economicFactor := newTestEconomicFactor("TestEconomy", 0.05)
// 	magicAccount := newMagic(economicFactor)

// 	// Act 1: Deposit an amount less than the allowed contribution
// 	magicAccount.Deposit("TestEconomy", 5000.0)

// 	// Assert 1: Check that the deposit is applied correctly
// 	if magicAccount.Balances["TestEconomy"].Total != 5000.0 {
// 		t.Fatalf("Expected balance to be 5000.0 but got %f", magicAccount.Balances["TestEconomy"].Total)
// 	}

// 	// Act 2: Deposit an amount greater than the allowed contribution
// 	magicAccount.Deposit("TestEconomy", 15000.0)

// 	// Assert 2: Ensure the deposit is capped at the allowed contribution
// 	if magicAccount.Balances["TestEconomy"].Total != 10000.0 {
// 		t.Fatalf("Expected balance to be capped at 10000.0 but got %f", magicAccount.Balances["TestEconomy"].Total)
// 	}
// }

// func TestAccrue(t *testing.T) {
// 	// Arrange
// 	economicFactor := newTestEconomicFactor("TestEconomy", 0.05)
// 	magicAccount := newMagic(economicFactor)
// 	magicAccount.Deposit("TestEconomy", 1000.0)

// 	// Act
// 	magicAccount.Balances["TestEconomy"].Accrue()

// 	// Assert
// 	expectedBalance := 1000.0 * (1 + 0.05 - Inflation.Rate)
// 	if magicAccount.Balances["TestEconomy"].Total != expectedBalance {
// 		t.Fatalf("Expected balance to accrue to %f but got %f", expectedBalance, magicAccount.Balances["TestEconomy"].Total)
// 	}
// }
