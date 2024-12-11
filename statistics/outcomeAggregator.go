package statistics

import "sync"

type OutcomeAggregator struct {
	// raw
	count           int
	failureCount    float64
	sum             float64
	sumOfSquares    float64
	precisionTarget float64
	windowSize      int
	meanWindow      []float64
	mu              sync.Mutex
	// calculated
	PPF          float64
	mean         float64
	variance     float64
	min          float64
	q1           float64
	q2           float64
	q3           float64
	max          float64
	learningRate float64
	stable       bool
}

func NewOutcomeAggregator(precisionTarget float64, windowSize int) *OutcomeAggregator {
	return &OutcomeAggregator{
		precisionTarget: precisionTarget,
		windowSize:      windowSize,
		meanWindow:      make([]float64, windowSize),
		stable:          false,
	}
}

func (oa *OutcomeAggregator) AddOutcome(outcome float64) bool {

	oa.mu.Lock()
	defer oa.mu.Unlock()

	if outcome <= 0 {
		oa.failureCount++
	}

	oa.PPF = oa.failureCount / float64(oa.count)

	oa.sum += outcome
	oa.sumOfSquares += outcome * outcome
	oa.count++

	newMean := oa.sum / float64(oa.count)
	oa.mean = newMean
	newVar := (oa.sumOfSquares / float64(oa.count)) - (oa.mean * oa.mean)
	oa.variance = newVar

	index := oa.count % oa.windowSize
	oa.meanWindow[index] = oa.mean

	if oa.count%10 == 0 {
		meanRange := Range(oa.meanWindow)
		// TODO: this method of median estimation is not great
		oa.learningRate = meanRange
		oa.stable = oa.count > oa.windowSize && meanRange <= float64(oa.precisionTarget)
	}

	// initalize guesses
	if oa.count == 1 {
		oa.q1 = oa.mean
		oa.q2 = 0 //oa.mean
		oa.q3 = oa.mean
		oa.min = outcome
		oa.max = outcome
	}

	// increment guesses
	// https://citeseerx.ist.psu.edu/document?repid=rep1&type=pdf&doi=825891dc5d461d93a06a9db5c4c7075cb738767f
	oa.q1 += oa.learningRate * (Sgn(outcome-oa.q1) + 2.0*(.25) - 1.0)
	oa.q2 += oa.learningRate * (Sgn(outcome - oa.q2))
	oa.q3 += oa.learningRate * (Sgn(outcome-oa.q3) + 2.0*(.75) - 1.0)

	// update min and max
	if outcome < oa.min {
		oa.min = outcome
	}
	if outcome > oa.max {
		oa.max = outcome
	}

	return oa.stable
}
