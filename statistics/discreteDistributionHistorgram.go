package statistics

// import (
// 	"errors"
// 	"math"
// )

// type DiscreteDistribution struct {
// 	outcomeBuckets []float64
// 	bucketSize     int
// 	bucketCount    int
// 	meanWindow     []float64
// 	windowSize     int
// 	sum            float64
// 	sumOfSquares   float64
// 	count          int
// 	failureCount   int
// }

// type Summary struct {
// 	count    int
// 	PPF      float64
// 	mean     float64
// 	variance float64
// 	min      float64
// 	q1       float64
// 	q2       float64
// 	q3       float64
// 	max      float64
// }

// func NewDiscreteDistribution(bucketCount int, bucketSize int, windowSize int) *DiscreteDistribution {
// 	return &DiscreteDistribution{
// 		bucketCount:    bucketCount,
// 		bucketSize:     bucketSize,
// 		outcomeBuckets: make([]float64, bucketCount),
// 		windowSize:     windowSize,
// 		meanWindow:     make([]float64, windowSize),
// 	}
// }

// func (dd *DiscreteDistribution) AddOutcome(outcome float64) (bool, error) {

// 	if outcome > float64(dd.bucketCount)*float64(dd.bucketSize) {
// 		return false, errors.New("should have allocted more bucket spacde")
// 	}

// 	// should this be an error
// 	if outcome < 0 {
// 		dd.failureCount++
// 		outcome = 0
// 	}
// 	// update the aggregators
// 	dd.outcomeBuckets[int(math.Floor(outcome/float64(dd.bucketSize)))]++
// 	dd.sum += outcome
// 	dd.sumOfSquares += outcome * outcome
// 	dd.count++

// 	// Update the mean window with the new value
// 	mean := dd.sum / float64(dd.count)
// 	index := dd.count % dd.windowSize
// 	dd.meanWindow[index] = mean

// 	// the distribution is stable if every mean in the window is within the bucket size
// 	stable := dd.count%10 == 0 && dd.count > dd.windowSize &&
// 		Range(dd.meanWindow) <= float64(dd.bucketSize)

// 	return stable, nil
// }

// func (dd *DiscreteDistribution) Compute() Summary {

// 	minIndex := dd.bucketCount - 1
// 	maxIndex := 0

// 	// TODO: double check this algorithm
// 	for i, count := range dd.outcomeBuckets {
// 		if count > 0 && i < minIndex {
// 			minIndex = i
// 		}
// 		if count > 0 && i > maxIndex {
// 			maxIndex = i
// 		}
// 	}

// 	results := Summary{
// 		count:    dd.count,
// 		PPF:      float64(dd.failureCount) / float64(dd.count),
// 		mean:     dd.sum / float64(dd.count),
// 		variance: dd.sumOfSquares / float64(dd.count),
// 		min:      float64(minIndex * dd.bucketSize),
// 		max:      float64(maxIndex * dd.bucketSize),
// 		q1:       math.Round(.25*dd.sum) * float64(dd.bucketSize),
// 		q2:       math.Round(.50*dd.sum) * float64(dd.bucketSize),
// 		q3:       math.Round(.75*dd.sum) * float64(dd.bucketSize),
// 	}

// 	return results
// }
