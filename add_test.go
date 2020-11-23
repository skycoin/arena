package main

import (
	"math"
	"testing"
)

var tests = []struct {
	first  float64
	second float64
	sum    float64
}{
	{1, 1, 2},
	{1.1, 2.2, 3.3},
	{0.2, 0.1, 0.3},
	{-1, -2.1, -3.1},
	{-0.1, 1, 0.9},
}

func TestAdd(t *testing.T) {
	for _, testCase := range tests {
		const float64EqualityThreshold = 1e-9
		actualSum := testCase.first + testCase.second
		if math.Abs(actualSum-testCase.sum) > float64EqualityThreshold {
			t.Fatalf("Result is not correct: %f + %f = %f, got %f", testCase.first, testCase.second, actualSum, testCase.sum)
		}
	}
}
