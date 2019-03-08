package utils

import "math/big"

const base = 10

// ParseBigFloat will parse big.Float from the given string
func ParseBigFloat(s string) (*big.Float, error) {
	number, _, err := new(big.Float).Parse(s, base)
	return number, err
}
