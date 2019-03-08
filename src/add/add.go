package add

import (
	"math/big"
)

// Data is to scan any number of numbers for addition
type Data struct {
	Numbers []*big.Float
}

// Add function will return the sum of two numbers, reusable function to add more than 2 numbers at a time.
func (d *Data) Add() *big.Float {
	sum := big.NewFloat(0)
	for _, number := range d.Numbers {
		sum.Add(sum, number)
	}
	return sum
}
