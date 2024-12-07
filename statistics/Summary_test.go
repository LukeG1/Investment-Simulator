package statistics

import (
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3},
		{[]float64{-1, -2, -3, -4, -5}, -3},
		{[]float64{}, math.NaN()},
		{[]float64{10}, 10},
	}

	for _, test := range tests {
		result := Mean(test.input)
		if math.IsNaN(test.expected) {
			if !math.IsNaN(result) {
				t.Errorf("Mean(%v) = %v, expected NaN", test.input, result)
			}
		} else if result != test.expected {
			t.Errorf("Mean(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, math.Sqrt(2)},
		{[]float64{10, 10, 10}, 0},
		{[]float64{}, math.NaN()},
		{[]float64{-1, -1, 1, 1}, 1},
	}

	for _, test := range tests {
		result := StandardDeviation(test.input)
		if math.IsNaN(test.expected) {
			if !math.IsNaN(result) {
				t.Errorf("StandardDeviation(%v) = %v, expected NaN", test.input, result)
			}
		} else if math.Abs(result-test.expected) > 1e-9 {
			t.Errorf("StandardDeviation(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestPercentile(t *testing.T) {
	tests := []struct {
		input    []float64
		p        float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 50, 3},
		{[]float64{1, 2, 3, 4, 5}, 25, 2},
		{[]float64{1, 2, 3, 4, 5}, 75, 4},
		{[]float64{}, 50, math.NaN()},
		{[]float64{1, 2, 3, 4, 5}, -1, math.NaN()},
		{[]float64{1, 2, 3, 4, 5}, 101, math.NaN()},
		{[]float64{10, 20, 30, 40}, 50, 25},
	}

	for _, test := range tests {
		result := Percentile(test.input, test.p)
		if math.IsNaN(test.expected) {
			if !math.IsNaN(result) {
				t.Errorf("Percentile(%v, %v) = %v, expected NaN", test.input, test.p, result)
			}
		} else if math.Abs(result-test.expected) > 1e-9 {
			t.Errorf("Percentile(%v, %v) = %v, expected %v", test.input, test.p, result, test.expected)
		}
	}
}

func TestQ1(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := 2.0
	result := Q1(input)
	if result != expected {
		t.Errorf("Q1(%v) = %v, expected %v", input, result, expected)
	}
}

func TestQ2(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := 3.0
	result := Q2(input)
	if result != expected {
		t.Errorf("Q2(%v) = %v, expected %v", input, result, expected)
	}
}

func TestQ3(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := 4.0
	result := Q3(input)
	if result != expected {
		t.Errorf("Q3(%v) = %v, expected %v", input, result, expected)
	}
}

func TestIQR(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := 2.0 // Q3 - Q1 = 4 - 2
	result := IQR(input)
	if result != expected {
		t.Errorf("IQR(%v) = %v, expected %v", input, result, expected)
	}
}
