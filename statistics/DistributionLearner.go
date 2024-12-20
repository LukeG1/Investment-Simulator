package statistics

// https://github.com/cxxr/LiveStats/blob/master/livestats/livestats.py
// https://www.cs.wustl.edu/~jain/papers/ftp/psqr.pdf

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// TODO: document
// TODO: refactor stability check to struct that can be applied to multiple metrics using delta of ewma
// TODO: stability checker based on rate of change of ewma, what if your random values happen to be the
// same twice in a row, it would come as stable, wait for n iterations minimum to be stable? use a small
// window? other solution?

const quantileLength int = 5
const minimumIterations int = 30

func calcP2(qp1, q, qm1, d, np1, n, nm1 float64) float64 {
	outer := d / (np1 - nm1)
	innerLeft := (n - nm1 + d) * (qp1 - q) / (np1 - n)
	innerRight := (np1 - n - d) * (q - qm1) / (n - nm1)

	return q + outer*(innerLeft+innerRight)
}

type quantile struct {
	dn          []float64
	npos        []float64
	pos         []float64
	heights     []float64
	initialized bool
	p           float64
}

func newQuantile(p float64) *quantile {
	posRange := make([]float64, quantileLength)
	for i := 0; i < quantileLength; i++ {
		posRange[i] = float64(i + 1)
	}
	return &quantile{
		dn:          []float64{0, p / 2, p, (1 + p) / 2, 1},
		npos:        []float64{1, 1 + 2*p, 1 + 4*p, 3 + 2*p, 5},
		pos:         posRange,
		heights:     []float64{},
		initialized: false,
		p:           p,
	}
}

func (q *quantile) add(outcome float64) {
	if len(q.heights) < quantileLength {
		q.heights = append(q.heights, outcome)
		return
	}

	if !q.initialized {
		sort.Float64s(q.heights)
		q.initialized = true
	}

	k := -1
	if outcome < q.heights[0] {
		q.heights[0] = outcome
		k = 0
	} else {
		for i := 1; i < quantileLength; i++ {
			if q.heights[i-1] <= outcome && outcome < q.heights[i] {
				k = i
				break
			}
		}
		if k == -1 {
			k = quantileLength - 1
			if outcome > q.heights[k] {
				q.heights[k] = outcome
			}
		}
	}

	for i := range q.pos {
		if i >= k {
			q.pos[i]++
		}
	}

	for i := range q.npos {
		q.npos[i] += q.dn[i]
	}

	q.adjust()
}

func (q *quantile) adjust() {
	for i := 1; i < quantileLength-1; i++ {
		n := q.pos[i]
		d := q.npos[i] - n

		if (d >= 1 && q.pos[i+1]-n > 1) || (d <= -1 && q.pos[i-1]-n < -1) {
			dd := math.Copysign(1, d)
			newHeight := calcP2(q.heights[i+1], q.heights[i], q.heights[i-1], dd, q.pos[i+1], q.pos[i], q.pos[i-1])

			if q.heights[i-1] < newHeight && newHeight < q.heights[i+1] {
				q.heights[i] = newHeight
			} else {
				q.heights[i] += dd * (q.heights[i+int(dd)] - q.heights[i]) / (q.pos[i+int(dd)] - q.pos[i])
			}
			q.pos[i] += dd
		}
	}
}

func (q *quantile) quantile() float64 {
	if q.initialized {
		return q.heights[2]
	}

	sort.Float64s(q.heights)
	l := float64(len(q.heights))
	return q.heights[int(math.Min(math.Max(l-1, 0), l*q.p))]
}

type StabilityChecker struct {
	ewma            float64
	alpha           float64 // Smoothing factor
	precisionTarget float64
	stable          bool
	count           int
	lastDelta       float64
}

// NewStabilityChecker creates a new StabilityChecker with a specified precision target
// and an optional smoothing factor (alpha). If alpha is zero, it defaults to 2/(n+1).
func NewStabilityChecker(precisionTarget float64, alpha float64) *StabilityChecker {
	if alpha <= 0 || alpha > 1 {
		alpha = 0.0 // Default to dynamic alpha calculation
	}
	return &StabilityChecker{
		precisionTarget: precisionTarget,
		alpha:           alpha,
	}
}

// Update processes a new value and updates EWMA and stability status
func (sc *StabilityChecker) Update(value float64) {
	sc.count++

	// Calculate smoothing factor dynamically for the first few samples if alpha is not set
	if sc.alpha == 0 {
		sc.alpha = 2.0 / (float64(sc.count) + 1.0)
	}

	// Update EWMA using the smoothing factor
	oldEwma := sc.ewma
	sc.ewma = sc.alpha*value + (1-sc.alpha)*sc.ewma

	// Check stability: compare delta with the precision target
	delta := math.Abs(sc.ewma - oldEwma)
	sc.lastDelta = delta
	if sc.count > 1 && delta < sc.precisionTarget {
		sc.stable = true
	} else {
		sc.stable = false
	}
}

type DistributionLearner struct {
	minVal          float64
	maxVal          float64
	varM2           float64
	skewM3          float64
	kurtM4          float64
	mean            float64
	count           int
	failureCount    int
	quantiles       map[float64]*quantile
	randSrc         *rand.Rand
	precisionTarget float64
	// windowSize      int
	// meanWindow      []float64
	meanStability *StabilityChecker
	ppfStability  *StabilityChecker
}

// TODO: should I really be looking at a window of the varaince of the estimates
func NewDistributionLearner(precisionTarget float64) *DistributionLearner {
	return &DistributionLearner{
		minVal:          math.Inf(1),
		maxVal:          math.Inf(-1),
		quantiles:       map[float64]*quantile{0.25: newQuantile(0.25), 0.5: newQuantile(0.5), 0.75: newQuantile(0.75)},
		randSrc:         rand.New(rand.NewSource(time.Now().UnixNano())),
		precisionTarget: precisionTarget,
		// windowSize:      windowSize,
		// meanWindow:      make([]float64, windowSize),
		meanStability: NewStabilityChecker(precisionTarget, .05),
		ppfStability:  NewStabilityChecker(.01/100., 0),
	}
}

func (dl *DistributionLearner) AddOutcome(outcome float64) {
	delta := outcome - dl.mean
	dl.minVal = math.Min(dl.minVal, outcome)
	dl.maxVal = math.Max(dl.maxVal, outcome)

	dl.mean += delta / float64(dl.count+1)
	dl.count++

	// index := dl.count % dl.windowSize
	// dl.meanWindow[index] = dl.mean

	dl.varM2 += delta * (outcome - dl.mean)

	if outcome < 0 {
		dl.failureCount++
	}

	dl.meanStability.Update(outcome)
	dl.ppfStability.Update(float64(dl.failureCount) / float64(dl.count))

	for _, q := range dl.quantiles {
		q.add(outcome)
	}

	dl.kurtM4 += math.Pow(outcome-dl.mean, 4)
	dl.skewM3 += math.Pow(outcome-dl.mean, 3)
}

type LearnedSummary struct {
	Stable   bool    `json:"Stable"`
	Count    int     `json:"Count"`
	PPF      float64 `json:"PPF"`
	Mean     float64 `json:"Mean"`
	Variance float64 `json:"Variance"`
	Kurtosis float64 `json:"Kurtosis"`
	Skewness float64 `json:"Skewness"`
	Min      float64 `json:"Min"`
	Q1       float64 `json:"Q1"`
	Q2       float64 `json:"Q2"`
	Q3       float64 `json:"Q3"`
	Max      float64 `json:"Max"`
}

func (dl *DistributionLearner) Summarize() *LearnedSummary {
	variance := dl.varM2 / float64(dl.count)
	return &LearnedSummary{
		Stable:   dl.count > minimumIterations && dl.meanStability.stable,
		Count:    dl.count,
		PPF:      float64(dl.failureCount) / float64(dl.count),
		Mean:     dl.mean,
		Variance: variance,
		Kurtosis: dl.kurtM4/(float64(dl.count)*math.Pow(variance, 2)) - 3.0,
		Skewness: dl.skewM3 / (float64(dl.count) * math.Pow(variance, 1.5)),
		Min:      dl.minVal,
		Q1:       dl.quantiles[0.25].quantile(),
		Q2:       dl.quantiles[0.5].quantile(),
		Q3:       dl.quantiles[0.75].quantile(),
		Max:      dl.maxVal,
	}
}
