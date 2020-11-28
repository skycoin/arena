package service

import "math/big"

func SumTwoNumbers(n1, n2 *big.Float) *big.Float {
	return new(big.Float).Add(n1, n2)
}
