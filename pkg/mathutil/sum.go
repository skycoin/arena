package mathutil

import "math/big"

// Sum calculates a sum of a given arguments of type `*big.Float`. Assumes that sum of
// no arguments is 0.
func Sum(args ...*big.Float) *big.Float {
	if len(args) == 0 {
		return big.NewFloat(0)
	}

	sum := new(big.Float)
	for _, v := range args {
		sum.Add(sum, v)
	}

	return sum
}
