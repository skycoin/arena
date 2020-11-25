package addculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const float64Epsilon = 1e-9

func TestAddFunction(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input  []float64
		output float64
	}{
		{[]float64{2, 4}, 6},
		{[]float64{2, -1}, 1},
		{[]float64{2, -2}, 0},
		{[]float64{0, -1.0}, -1},
		{[]float64{2.3, -1.7}, 0.6},
	}

	for _, test := range tests {
		if test.output == 0 {
			assert.InDelta(AddFunction(test.input), test.output, float64Epsilon)
		} else {
			assert.InEpsilon(AddFunction(test.input), test.output, float64Epsilon)
		}
	}
}
