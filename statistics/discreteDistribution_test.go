package statistics

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	n := 10_000_000
	dd := NewDiscreteDistribution(n, 100, .01)

	shouldStop := false
	for i := 0; i < n; i++ {
		shouldStop = dd.AddOutcome(2 * rand.NormFloat64())
		if shouldStop {
			fmt.Printf("Stopped at iteration: %d\n", i)
			break
		}
	}

	fmt.Println(Range(dd.meanWindow))
	fmt.Println(shouldStop)

	fmt.Println(dd.Compute().standardDeviation)

}
