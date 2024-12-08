package statistics

import (
	"fmt"
	"testing"
)

func TestAddConstant(t *testing.T) {

	const n int = 1_000_000_000

	values := [n]float64{}

	for i := 0; i < n; i++ {
		values[i] = 2
	}

	testDistribution := DiscreteDistribtuion{
		values[:],
	}

	testDistribution.Add(1)

	fmt.Println(testDistribution.outcomes[n/2])
}
