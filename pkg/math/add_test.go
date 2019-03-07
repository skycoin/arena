package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	first    float64
	second   float64
	expected float64
}{
	{
		first:    11,
		second:   22,
		expected: 33,
	},
	{
		first:    121.5,
		second:   167,
		expected: 288.5,
	},
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	for _, test := range tests {
		result := Add(test.first, test.second)
		assert.Equal(test.expected, result, "%v + %v should be %v", test.first, test.second, test.expected)
	}
}
