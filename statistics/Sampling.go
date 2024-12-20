package statistics

import (
	"math"
	"math/rand"
	"time"
)

// TODO: make my naming scheme for statistical concepts more consistent

type NaieveDataSampler struct {
	data  *[]float64 // the raw data
	h     float64    // bandwidth
	mu    float64    // mean
	sigma float64    // standard deviation
	n     int32      // the number of data points
	r     *rand.Rand // source of randomness
}

func GenerateKernelSampler(data *[]float64) NaieveDataSampler {
	sigma := StandardDeviation(*data)
	mu := Mean(*data)
	n := int32(len(*data))

	// https://en.wikipedia.org/wiki/Kernel_density_estimation#A_rule-of-thumb_bandwidth_estimator
	// first parameter slighly tuned for my preferences, may not technically be optimal but works well for my data
	h := .55 * min(sigma, IQR(*data)/1.34) * math.Pow(float64(n), -(1./5.))

	return NaieveDataSampler{
		data,
		h,
		mu,
		sigma,
		n,
		// TODO: consider adding varried seed for better randomness after testing phase
		rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Kernel Density based sampling: https://www.stat.cmu.edu/~cshalizi/350/lectures/28/lecture-28.pdf
func (data *NaieveDataSampler) Sample() float64 {
	randomSource := data.r
	observationSource := (*data.data)[randomSource.Int31n(data.n)]
	return observationSource + data.h*randomSource.NormFloat64()
}
