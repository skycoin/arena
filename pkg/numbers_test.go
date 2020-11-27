package pkg

import (
	"fmt"
	"math"
	"testing"
)

var addTestsGreen = []struct {
	a   float64
	b   float64
	out float64
}{
	{10, 13, 23},
	{1120, 1121, 2241},
	{20.5, 20.5, 41},
	{512.5, 512.5, 1025},
}

func TestAddGreen(t *testing.T) {
	for _, tt := range addTestsGreen {
		t.Run(fmt.Sprint(tt.a), func(t *testing.T) {
			output, err := Add(tt.a, tt.b)
			if err != nil {
				t.Errorf("got %v when adding %v to %v", err, tt.a, tt.b)
			}
			if output != tt.out {
				t.Errorf("got %v, want %v", output, tt.out)
			}
		})
	}
}

var addTestsFailing = []struct {
	a   float64
	b   float64
	out float64
}{
	{math.MaxFloat64, math.MaxFloat64, math.Inf(1)},
	{-math.MaxFloat64, -math.MaxFloat64, math.Inf(-1)},
}

func TestAdd(t *testing.T) {
	for _, tt := range addTestsFailing {
		t.Run(fmt.Sprint(tt.a), func(t *testing.T) {
			output, err := Add(tt.a, tt.b)
			if err == nil {
				t.Errorf("got no error when adding %v to %v, returned %v", tt.a, tt.b, output)
			}
			if output != tt.out {
				t.Errorf("got %v, want %v", output, tt.out)
			}
		})
	}
}

var parseTests = []struct {
	in  string
	out float64
}{
	{"10", 10},
	{"1120", 1120},
	{"20.5", 20.5},
	{"512.5", 512.5},
}

func TestParseGreen(t *testing.T) {
	for _, tt := range parseTests {
		t.Run(tt.in, func(t *testing.T) {
			output, err := Parse(tt.in)
			if err != nil {
				t.Errorf("got %v when parsing %s", err, tt.in)
			}
			if output != tt.out {
				t.Errorf("got %v, want %v", output, tt.out)
			}
		})
	}
}

var parseTestsFailing = []struct {
	in  string
	out float64
}{
	{"hi there", 0},
	{"1.1.54", 0},
}

func TestParseFailing(t *testing.T) {
	for _, tt := range parseTestsFailing {
		t.Run(tt.in, func(t *testing.T) {
			output, err := Parse(tt.in)
			if err == nil {
				t.Errorf("got no error when parsing %v, returned %v", tt.in, output)
			}
			if output != tt.out {
				t.Errorf("got %v, want %v", output, tt.out)
			}
		})
	}
}
