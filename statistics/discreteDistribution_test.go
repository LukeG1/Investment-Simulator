package statistics

// // import (
// // 	"InvestmentSimulator/models"
// // 	"fmt"
// // 	"testing"
// // )

// // func TestAdd(t *testing.T) {
// // 	const n int = 10_000_000
// // 	principal := 1000.

// // 	dd := NewDiscreteDistribution(n, 200, 1)

// // 	matketSampler := models.SandP500.Sampler
// // 	inflationSampler := models.Inflation.Sampler

// // 	shouldStop := false
// // 	for i := 0; i < n; i++ {

// // 		balance := principal
// // 		for year := 0; year < 40; year++ {
// // 			balance += 100
// // 			balance *= 1 + matketSampler.Sample() - inflationSampler.Sample()
// // 		}
// // 		shouldStop = dd.AddOutcome(balance)
// // 		if shouldStop {
// // 			fmt.Printf("Stopped at iteration: %d\n", i+1)
// // 			break
// // 		}
// // 	}

// // 	fmt.Println(Range(dd.meanWindow))
// // 	fmt.Println(shouldStop)

// // 	res := dd.Compute()
// // 	fmt.Println(res.PPF)
// // 	fmt.Println(res.q2)
// // 	fmt.Println(res.standardDeviation)
// // }

// package statistics

// import (
// 	"math"
// 	"testing"
// )

// func TestNewDiscreteDistribution(t *testing.T) {
// 	maxOutcomes := 100
// 	windowSize := 10
// 	precisionTarget := 0.01

// 	dd := NewDiscreteDistribution(maxOutcomes, windowSize, precisionTarget)

// 	if dd.maxOutcomes != maxOutcomes {
// 		t.Errorf("expected maxOutcomes %d, got %d", maxOutcomes, dd.maxOutcomes)
// 	}

// 	if dd.windowSize != windowSize {
// 		t.Errorf("expected windowSize %d, got %d", windowSize, dd.windowSize)
// 	}

// 	if dd.precisionTarget != precisionTarget {
// 		t.Errorf("expected precisionTarget %f, got %f", precisionTarget, dd.precisionTarget)
// 	}
// }

// func TestAddOutcome(t *testing.T) {
// 	maxOutcomes := 10
// 	windowSize := 5
// 	precisionTarget := 0.01

// 	dd := NewDiscreteDistribution(maxOutcomes, windowSize, precisionTarget)

// 	outcomes := []float64{1.0, 2.0, 3.0, 4.0, 5.0, -1.0, 6.0, 7.0, 8.0, 9.0}
// 	for i, outcome := range outcomes {
// 		reached := dd.AddOutcome(outcome)
// 		if i < maxOutcomes-1 && reached {
// 			t.Errorf("unexpected precisionTargetMet or dataLimitReached before maxOutcomes, iteration %d", i)
// 		}
// 	}

// 	if dd.count != maxOutcomes {
// 		t.Errorf("expected count %d, got %d", maxOutcomes, dd.count)
// 	}

// 	if dd.failureCount != 1 {
// 		t.Errorf("expected failureCount 1, got %d", dd.failureCount)
// 	}
// }

// func TestComputeSummary(t *testing.T) {
// 	maxOutcomes := 10
// 	windowSize := 5
// 	precisionTarget := 0.01

// 	dd := NewDiscreteDistribution(maxOutcomes, windowSize, precisionTarget)

// 	outcomes := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
// 	for _, outcome := range outcomes {
// 		dd.AddOutcome(outcome)
// 	}

// 	summary := dd.Compute()

// 	if summary.count != maxOutcomes {
// 		t.Errorf("expected count %d, got %d", maxOutcomes, summary.count)
// 	}

// 	if math.Abs(summary.mean-5.5) > 1e-6 {
// 		t.Errorf("expected mean 5.5, got %f", summary.mean)
// 	}

// 	if math.Abs(summary.PPF-0) > 1e-6 {
// 		t.Errorf("expected PPF 0, got %f", summary.PPF)
// 	}

// 	if summary.min != 1.0 {
// 		t.Errorf("expected min 1.0, got %f", summary.min)
// 	}

// 	if summary.max != 10.0 {
// 		t.Errorf("expected max 10.0, got %f", summary.max)
// 	}

// 	expectedStandardDeviation := math.Sqrt(8.25)
// 	if math.Abs(summary.standardDeviation-expectedStandardDeviation) > 1e-6 {
// 		t.Errorf("expected standard deviation %f, got %f", expectedStandardDeviation, summary.standardDeviation)
// 	}

// 	expectedQ1 := 3.25
// 	if math.Abs(summary.q1-expectedQ1) > 1e-6 {
// 		t.Errorf("expected Q1 %f, got %f", expectedQ1, summary.q1)
// 	}

// 	expectedQ2 := 5.5
// 	if math.Abs(summary.q2-expectedQ2) > 1e-6 {
// 		t.Errorf("expected Q2 %f, got %f", expectedQ2, summary.q2)
// 	}

// 	expectedQ3 := 7.75
// 	if math.Abs(summary.q3-expectedQ3) > 1e-6 {
// 		t.Errorf("expected Q3 %f, got %f", expectedQ3, summary.q3)
// 	}
// }

// func TestAddOutcomePrecisionTarget(t *testing.T) {
// 	maxOutcomes := 100
// 	windowSize := 10
// 	precisionTarget := 0.5

// 	dd := NewDiscreteDistribution(maxOutcomes, windowSize, precisionTarget)

// 	outcomes := []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0} // Constant outcomes
// 	for i, outcome := range outcomes {
// 		reached := dd.AddOutcome(outcome)
// 		if i == len(outcomes)-1 && !reached {
// 			t.Errorf("expected precisionTargetMet or dataLimitReached at iteration %d", i)
// 		}
// 	}
// }
