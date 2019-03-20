package main

import "testing"
import sumNumbers "github.com/offgridsito/skycoin-arena-cli/pkg/sum"

var sumTest = []struct {
	addendOne float64
	addendTwo float64
	expected  float64
}{
	{0, 4, 4},
	{11, 33, 44},
	{-11, 2, -9},
	{2, 2, 4},
	{3, -1, 2},
	{2.33, 10.11, 12.44},
	{0, 10, 10},
	{8, 9, 17},
	{1, 8, 9},
}

func TestSum(t *testing.T) {
	for _, tt := range sumTest {
		actual := sumNumbers.SumNumbers(tt.addendOne, tt.addendTwo)
		if actual != tt.expected {
			t.Errorf("Sum(%f, %f): expected %f, actual %f", tt.addendOne, tt.addendTwo, tt.expected, actual)
		}
	}
}
