package statistics

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// TODO: document and rename

func calcP2(qp1, q, qm1, d, np1, n, nm1 float64) float64 {
	// TODO: asses float conversions
	outer := d / (np1 - nm1)
	inner_left := (n - nm1 + d) * (qp1 - q) / (np1 - n)
	inner_right := (np1 - n - d) * (q - qm1) / (n - nm1)

	return q + outer*(inner_left+inner_right)
}

const quantile_length int = 5

type quantile struct {
	dn         []float64
	npos       []float64
	pos        []float64
	heights    []float64
	initalized bool
	p          float64
}

func newQuantile(p float64) *quantile {
	posRange := [quantile_length + 1]float64{}
	for i := 0; i < quantile_length+1; i++ {
		posRange[i] = float64(i + 1)
	}
	return &quantile{
		dn:         []float64{0, p / 2, p, (1 + p) / 2, 1},
		npos:       []float64{1, 1 + 2*p, 1 + 4*p, 3 + 2*p, 5},
		pos:        posRange[:],
		heights:    []float64{},
		initalized: false,
		p:          p,
	}
}

func (q *quantile) add(outcome float64) {
	if len(q.heights) != quantile_length {
		q.heights = append(q.heights, outcome)
		return
	}

	if !q.initalized {
		sort.Float64s(q.heights)
		q.initalized = true
	}

	k := -1
	if outcome < q.heights[0] {
		q.heights[0] = outcome
		k = 1
	} else {
		for i := 1; i < quantile_length; i++ {
			if q.heights[i-1] <= outcome && outcome < q.heights[i] {
				k = i
				break
			}
		}
		// TODO: check if this condition functions like a for else
		if k <= 1 {
			k = 4
			if q.heights[len(q.heights)-1] < outcome {
				q.heights[len(q.heights)-1] = outcome
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

}

func (q *quantile) adjust() {
	for i := 1; i < 4; i++ {
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
	if q.initalized {
		return q.heights[2]
	}

	sort.Float64s(q.heights)
	l := float64(len(q.heights))
	return q.heights[int(min(max(l-1, 0), l*q.p))]
}

type DistributionLearner struct {
	minVal       float64
	maxVal       float64
	varM2        float64
	skewM3       float64
	kurtM4       float64
	mean         float64
	count        int
	failureCount int
	initalized   bool
	stable       bool
	quantiles    map[float64]*quantile
	randSrc      *rand.Rand
}

func NewDistributionLearner() *DistributionLearner {
	return &DistributionLearner{
		minVal:    math.Inf(1),
		maxVal:    math.Inf(-1),
		quantiles: map[float64]*quantile{0.25: newQuantile(0.25), 0.5: newQuantile(0.5), 0.75: newQuantile(0.75)},
		randSrc:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (dl *DistributionLearner) AddOutcome(outcome float64) {
	delta := outcome - dl.mean
	dl.minVal = math.Min(dl.minVal, outcome)
	dl.maxVal = math.Max(dl.maxVal, outcome)

	dl.mean += delta / float64(dl.count+1)
	dl.count++

	dl.varM2 += delta * (outcome - dl.mean)

	for _, q := range dl.quantiles {
		q.add(outcome)
	}

	dl.kurtM4 += math.Pow(outcome-dl.mean, 4)
	dl.skewM3 += math.Pow(outcome-dl.mean, 3)

	if outcome < 0 {
		dl.failureCount++
	}
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

func (dl *DistributionLearner) summaize() *LearnedSummary {
	// TODO: consider saftey checking some of these for being too new to have data

	variance := dl.varM2 / float64(dl.count)
	return &LearnedSummary{
		Stable:   false,
		Count:    dl.count,
		PPF:      float64(dl.failureCount) / float64(dl.count),
		Mean:     dl.mean,
		Variance: variance,
		Kurtosis: dl.kurtM4/(float64(dl.count)*math.Pow(variance, 2)) - 3.0,
		Skewness: dl.skewM3 / (float64(dl.count) * math.Pow(variance, 1.5)),
		Min:      dl.maxVal,
		Q1:       dl.quantiles[.25].quantile(),
		Q2:       dl.quantiles[.5].quantile(),
		Q3:       dl.quantiles[.75].quantile(),
		Max:      dl.maxVal,
	}
}
