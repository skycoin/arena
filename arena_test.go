package main

import (
	"fmt"
	"testing"
)

var cases = []struct{ argx, argy, result string }{
	{"1", "1", "2"},
	{"-1", "1", "0"},
	{"10e6", "-Inf", "-Inf"},
	{"10e-10", "2.2e-10", "1.2200000000000001e-09"},
	{"10e120", "0", "1e+121"},
	{"Inf", "Inf", "+Inf"},
	{"Inf", "-Inf", "NaN"},
	{"A", "B", "strconv.ParseFloat: parsing \"A\": invalid syntax"},
	{"10", "B", "strconv.ParseFloat: parsing \"B\": invalid syntax"},
}

func TestArena(t *testing.T) {
	for _, tt := range cases {
		test := fmt.Sprintf("%v + %v = %v", tt.argx, tt.argy, tt.result)
		t.Run(test, func(t *testing.T) {
			result := sumargs(tt.argx, tt.argy)
			if result != tt.result {
				t.Errorf("got %q, want %q ", result, tt.result)
			}
			t.Log(test)
		})
	}
}
