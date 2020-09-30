package adder

import (
	"errors"
	"math"
)

var ErrFloatOverflow = errors.New("float overflow")

func AddTwoFloatNumbers(left, right float64) (float64, error) {
	if right > 0 {
		if left > math.MaxFloat64-right {
			return 0, ErrFloatOverflow
		}
	} else {
		if left < math.SmallestNonzeroFloat64-right {
			return 0, ErrFloatOverflow
		}
	}
	return left + right, nil
}

var ErrIntOverflow = errors.New("integer overflow")

func AddTwoInt64Numbers(left, right int64) (int64, error) {
	if right > 0 {
		if left > math.MaxInt64-right {
			return 0, ErrIntOverflow
		}
	} else {
		if left < math.MinInt64-right {
			return 0, ErrIntOverflow
		}
	}
	return left + right, nil
}
