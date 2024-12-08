package statistics

import (
	"runtime"
	"sync"
)

type DiscreteDistribtuion struct {
	outcomes []float64
}

func (dd *DiscreteDistribtuion) Add(amount float64) {
	length := len(dd.outcomes)
	if length == 0 {
		return
	}

	// Determine the number of Goroutines to use
	numCPU := runtime.NumCPU()
	chunkSize := (length + numCPU - 1) / numCPU // Ceiling division for chunk size

	var wg sync.WaitGroup

	// Precompute chunk boundaries and launch Goroutines
	for i := 0; i < numCPU; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > length {
			end = length
		}

		if start >= length {
			break
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				dd.outcomes[j] += amount
			}
		}(start, end)
	}

	// Wait for all Goroutines to complete
	wg.Wait()
}
