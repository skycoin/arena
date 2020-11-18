package pkg

import "testing"

var SumTests = []struct {
	name string
	in   []float64
	out  float64
}{
	{"first", []float64{3, 4.5}, 7.5},
	{"second", []float64{3.1, 10.5}, 13.6},
	{"negative", []float64{3.1, -10.5}, -7.4},
}

func TestFlagParser(t *testing.T) {
	for _, tt := range SumTests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sum(tt.in[0], tt.in[1])
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}
