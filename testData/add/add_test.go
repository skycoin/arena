package testing

import (
	"testing"

	"github.com/skycoin/Addition/pkg/add"
)

var flagtests = []struct {
	input  add.Data
	result float64
}{
	{add.Data{
		Numbers: []float64{45.5, 23.2},
	}, 68.7},
	{add.Data{
		Numbers: []float64{4.5, 2.2},
	}, 6.7},
}

func TestSums(t *testing.T) {
	for _, tt := range flagtests {
		t.Run("test1", func(t *testing.T) {
			s := tt.input.Add()
			if s != tt.result {
				t.Errorf("got %.2f, want %.2f", s, tt.result)
			}
		})
	}
}
