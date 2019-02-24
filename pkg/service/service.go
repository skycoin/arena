package service

import (
	"math/big"
)

// AddTwoNumbers performs an addition of two numbers, passed as arguments of type `*big.Float` named `n1` and `n2`.
// It returns a value of type `*big.Float`.
func AddTwoNumbers(n1, n2 *big.Float) *big.Float {
	return new(big.Float).Add(n1, n2)
}
