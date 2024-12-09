package statistics

import (
	"math"
	"sort"
)

func Mean(numbers []float64) float64 {
	var sum float64 = 0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func StandardDeviation(numbers []float64) float64 {
	mean := Mean(numbers)

	var sum float64 = 0
	for _, num := range numbers {
		sum += math.Pow(num-mean, 2)
	}

	return math.Sqrt(sum / float64(len(numbers)))
}

// TODO: refactor with a second version that takes a sorted list, accessible externally but also called internally
// Percentile calculates the p-th percentile of a slice of numbers
// TODO: this is painfully slow
func Percentile(numbers []float64, p float64) float64 {
	if len(numbers) == 0 {
		return math.NaN()
	}

	if p < 0 || p > 100 {
		return math.NaN() // Invalid percentile
	}

	// Sort the slice
	sortedNumbers := append([]float64{}, numbers...)
	sort.Float64s(sortedNumbers)

	// Calculate the index in the sorted array
	index := p / 100 * float64(len(sortedNumbers)-1)
	lowerIndex := int(math.Floor(index))
	upperIndex := int(math.Ceil(index))

	// Interpolate between the two bounding values
	if lowerIndex == upperIndex {
		return sortedNumbers[lowerIndex]
	}

	lowerValue := sortedNumbers[lowerIndex]
	upperValue := sortedNumbers[upperIndex]
	weight := index - float64(lowerIndex)

	return lowerValue + weight*(upperValue-lowerValue)
}

// Q1 calculates the first quartile (25th percentile)
func Q1(numbers []float64) float64 {
	return Percentile(numbers, 25)
}

// Q2 calculates the second quartile (median or 50th percentile)
func Q2(numbers []float64) float64 {
	return Percentile(numbers, 50)
}

// Q3 calculates the third quartile (75th percentile)
func Q3(numbers []float64) float64 {
	return Percentile(numbers, 75)
}

// IQR calculates the interquartile range (Q3 - Q1)
func IQR(numbers []float64) float64 {
	return Q3(numbers) - Q1(numbers)
}

func Min(slice []float64) float64 {
	// Initialize the min value with the first element
	minVal := slice[0]
	// Iterate through the slice to find the minimum
	for _, value := range slice {
		if value < minVal {
			minVal = value
		}
	}
	return minVal
}
func Max(slice []float64) float64 {
	// Initialize the min value with the first element
	maxVal := slice[0]
	// Iterate through the slice to find the minimum
	for _, value := range slice {
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

// Function to calculate the range (max - min) of a slice of float64
func Range(slice []float64) float64 {
	// Initialize min and max with the first element
	minVal, maxVal := slice[0], slice[0]
	// Iterate through the slice to find both the minimum and maximum
	for _, value := range slice {
		if value < minVal {
			minVal = value
		}
		if value > maxVal {
			maxVal = value
		}
	}
	// Return the range (max - min)
	return maxVal - minVal
}
