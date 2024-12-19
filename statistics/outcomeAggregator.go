package statistics

// TODO: delete soon

import (
	"sync"
)

type OutcomeAggregator struct {
	// raw
	Count           int
	FailureCount    float64
	Sum             float64
	SumOfSquares    float64
	PrecisionTarget float64
	WindowSize      int
	MeanWindow      []float64
	Mu              sync.Mutex
	// calculated
	PPF          float64 `json:"ppf"`
	Mean         float64 `json:"mean"`
	Variance     float64 `json:"variance"`
	Min          float64 `json:"min"`
	Q1           float64 `json:"q1"`
	Q2           float64 `json:"q2"`
	Q3           float64 `json:"q3"`
	Max          float64 `json:"max"`
	LearningRate float64
	Stable       bool `json:"stable"`
}

func NewOutcomeAggregator(precisionTarget float64, windowSize int) *OutcomeAggregator {
	return &OutcomeAggregator{
		PrecisionTarget: precisionTarget,
		WindowSize:      windowSize,
		MeanWindow:      make([]float64, windowSize),
		Stable:          false,
	}
}

func (oa *OutcomeAggregator) AddOutcome(outcome float64) bool {

	oa.Mu.Lock()
	defer oa.Mu.Unlock()

	if outcome <= 1000 {
		oa.FailureCount++
	}

	oa.PPF = oa.FailureCount / float64(oa.Count)

	oa.Sum += outcome
	oa.SumOfSquares += outcome * outcome
	oa.Count++

	newMean := oa.Sum / float64(oa.Count)
	oa.Mean = newMean
	newVar := (oa.SumOfSquares / float64(oa.Count)) - (oa.Mean * oa.Mean)
	oa.Variance = newVar

	index := oa.Count % oa.WindowSize
	oa.MeanWindow[index] = oa.Mean

	if oa.Count%10 == 0 {
		meanRange := Range(oa.MeanWindow)
		// TODO: this method of median estimation is not great
		oa.LearningRate = meanRange
		oa.Stable = oa.Count > oa.WindowSize && meanRange <= float64(oa.PrecisionTarget)
	}

	// initalize guesses
	if oa.Count == 1 {
		oa.Q1 = oa.Mean
		oa.Q2 = 0 //oa.Mean
		oa.Q3 = oa.Mean
		oa.Min = outcome
		oa.Max = outcome
	}

	// increment guesses
	// https://citeseerx.ist.psu.edu/document?repid=rep1&type=pdf&doi=825891dc5d461d93a06a9db5c4c7075cb738767f
	oa.Q1 += oa.LearningRate * (Sgn(outcome-oa.Q1) + 2.0*(.25) - 1.0)
	oa.Q2 += oa.LearningRate * (Sgn(outcome - oa.Q2))
	oa.Q3 += oa.LearningRate * (Sgn(outcome-oa.Q3) + 2.0*(.75) - 1.0)

	// update min and max
	if outcome < oa.Min {
		oa.Min = outcome
	}
	if outcome > oa.Max {
		oa.Max = outcome
	}

	// TODO: could be cool to return stable and a measure of how close to stable it is?
	return oa.Stable
}
