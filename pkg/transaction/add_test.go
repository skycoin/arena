package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		X        int
		Y        int
		expected int
	}{
		{1, 2, 3},
		{3, 4, 7},
		{5, 6, 11},
		{7, 11, 18},
		{8, 21, 29},
		{76, 31, 107},
		{5, 4, 9},
	}

	for _, test := range tests {
		result, _ := AddTwoNumber(test.X, test.Y)
		assert.Equal(result, test.expected)
	}
}
