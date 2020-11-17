package adder

import (
	"github.com/shopspring/decimal"
)

const maxArgsCount = 2

// Adder interface
type Adder interface {
	Add() decimal.Decimal
}

type adder struct {
	nums []decimal.Decimal
}

// New init adder from args
func New(args []string) (Adder, error) {
	argsLen := len(args)
	if argsLen != maxArgsCount {
		return nil, ErrArgCount
	}
	nums := make([]decimal.Decimal, argsLen)
	for i, s := range args {
		num, err := decimal.NewFromString(s)
		if err != nil {
			return nil, NewArgError(i, err)
		}
		nums[i] = num
	}
	return adder{
		nums: nums,
	}, nil
}

// Return adds of initialized values
func (a adder) Add() decimal.Decimal {
	var res decimal.Decimal
	for _, num := range a.nums {
		res = res.Add(num)
	}
	return res
}
