package statistics

import "sort"

type DiscreteDistribution struct {
	outcomes        []float64
	meanWindow      []float64
	windowSize      int
	precisionTarget float64
	maxOutcomes     int
	sum             float64
	count           int
	failureCount    int
}

// TODO: consider if this is the best place, could perhaps be named more aptly here and have a substruct for 5 num summary that lives in the summary file
type Summary struct {
	count             int
	PPF               float64
	mean              float64
	standardDeviation float64
	min               float64
	q1                float64
	q2                float64
	q3                float64
	max               float64
}

func NewDiscreteDistribution(maxOutcomes int, windowSize int, precisionTarget float64) *DiscreteDistribution {
	return &DiscreteDistribution{
		maxOutcomes:     maxOutcomes,
		outcomes:        make([]float64, maxOutcomes),
		windowSize:      windowSize,
		meanWindow:      make([]float64, windowSize),
		precisionTarget: precisionTarget,
	}
}

func (dd *DiscreteDistribution) AddOutcome(outcome float64) bool {
	// update the aggregators
	dd.outcomes[dd.count] = outcome
	dd.sum += outcome
	dd.count++

	// Update the mean window with the new value
	mean := dd.sum / float64(dd.count)
	// TODO: consider switching to float64(dd.failureCount) / float64(dd.count)
	index := dd.count % dd.windowSize
	dd.meanWindow[index] = mean

	// update failure count
	if outcome <= 0 {
		dd.failureCount++
	}

	precisionTargetMet := dd.count%10 == 0 && dd.count > dd.windowSize &&
		Range(dd.meanWindow) <= dd.precisionTarget
	dataLimitReached := dd.count == dd.maxOutcomes
	return precisionTargetMet || dataLimitReached
}

// TODO: each account / asset balance can be parrellized in go routines?

// TODO: think of this more as a snapshot, consider how to increase the count values and such this whole idea may actually be bad
func (dd *DiscreteDistribution) Compute() Summary {

	results := Summary{
		count: dd.count,
		PPF:   float64(dd.failureCount) / float64(dd.count),
		mean:  dd.sum / float64(dd.count),
	}

	// sort the list
	dd.outcomes = dd.outcomes[:dd.count]
	sort.Float64s(dd.outcomes)

	// populate results from the sorted list
	results.min = dd.outcomes[0]
	results.q1 = dd.outcomes[int(.25*float64(dd.count-1))]
	results.q2 = dd.outcomes[int(.50*float64(dd.count-1))]
	results.q3 = dd.outcomes[int(.75*float64(dd.count-1))]
	results.max = dd.outcomes[dd.count-1]

	// calculate standard deviation
	results.standardDeviation = StandardDeviation(dd.outcomes)

	// reset the values for the next year
	dd.outcomes = make([]float64, dd.maxOutcomes)
	dd.meanWindow = make([]float64, dd.windowSize)
	dd.count = 0
	dd.sum = 0
	dd.failureCount = 0

	return results
}
