package adder

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoFloatNumbers(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		left, right, result float64
		err                 error
	}{
		{1.0, 1.0, 2.0, nil},
		{1.3, 1.3, 2.6, nil},
		{1.3, 0, 1.3, nil},
		{math.MaxFloat64, math.MaxFloat64 + 1.0, 0, ErrFloatOverflow},
		{0, 0, 0, ErrFloatOverflow},
	}
	for _, tt := range tests {
		tt := tt
		floatSum, floatErr := AddTwoFloatNumbers(tt.left, tt.right)
		assert.Equal(floatSum, tt.result, "Sum is incorect")
		assert.Equal(floatErr, tt.err, "Overflow")
	}
}
func TestAddTwoIntNumbers(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		left, right, result int64
		err                 error
	}{
		{0, 0, 0, nil},
		{1, 2, 3, nil},
		{math.MaxInt64, 1, 0, ErrIntOverflow},
		{1, math.MaxInt64, 0, ErrIntOverflow},
		{math.MaxInt64, math.MaxInt64, 0, ErrIntOverflow},
		{math.MinInt64, math.MinInt64, 0, ErrIntOverflow},
	}
	for _, tt := range tests {
		tt := tt
		floatSum, floatErr := AddTwoInt64Numbers(tt.left, tt.right)
		assert.Equal(floatErr, tt.err, "Overflow")
		assert.Equal(floatSum, tt.result, "Sum is incorect")
	}
}
