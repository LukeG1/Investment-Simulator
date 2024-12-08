package statistics

import (
	"math"
	"testing"
)

func floatsAreClose(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

// Take in the stock market values, calculate sample statistics, sample, check that they are similar
func TestSample(t *testing.T) {
	const samples int = 1_000_000

	inputData := SandP500
	dataSampler := GenerateKernelSampler(&inputData)
	outputData := [samples]float64{}

	for i := 0; i < samples; i++ {
		outputData[i] = dataSampler.Sample()
	}

	if !floatsAreClose(Mean(inputData), Mean(outputData[:]), .01) {
		t.Fatalf(`Means do not match`)
	}

	if !floatsAreClose(StandardDeviation(inputData), StandardDeviation(outputData[:]), .01) {
		t.Fatalf(`StandardDeviations do not match`)
	}

	if !floatsAreClose(Q2(inputData), Q2(outputData[:]), .01) {
		t.Fatalf(`Medians do not match`)
	}

}
