package main

import (
	"arena/calc"
	"math/big"
	"testing"
)

type TestStruct struct {
	Result *big.Float
	Input  calc.AddStruct
}

var tests = []struct {
	Result *big.Float
	Input  calc.AddStruct
}{
	{
		Result: big.NewFloat(10999999.9998),
		Input: calc.AddStruct{
			X: big.NewFloat(999999.9999),
			Y: big.NewFloat(9999999.9999),
		},
	}, {
		Result: big.NewFloat(20.20),
		Input: calc.AddStruct{
			X: big.NewFloat(10.10),
			Y: big.NewFloat(10.10),
		},
	}, {
		Result: big.NewFloat(9000000),
		Input: calc.AddStruct{
			X: big.NewFloat(-999999.9999),
			Y: big.NewFloat(9999999.9999),
		},
	}, {
		Result: big.NewFloat(-8999999.9999),
		Input: calc.AddStruct{
			X: big.NewFloat(999999.9999999),
			Y: big.NewFloat(-9999999.9999),
		},
	},
	{
		Result: big.NewFloat(-10000099.999),
		Input: calc.AddStruct{
			X: big.NewFloat(-99.9999999),
			Y: big.NewFloat(-9999999.9999),
		},
	},
	{
		Result: big.NewFloat(99.99999999999999),
		Input: calc.AddStruct{
			X: big.NewFloat(99.99999999999999),
			Y: big.NewFloat(0),
		},
	},
	{
		Result: big.NewFloat(0),
		Input: calc.AddStruct{
			X: big.NewFloat(0),
			Y: big.NewFloat(0),
		},
	},
	{
		Result: big.NewFloat(1000000009.999999999999),
		Input: calc.AddStruct{
			X: big.NewFloat(999999999),
			Y: big.NewFloat(10.999999999999),
		},
	},
	{
		Result: big.NewFloat(0.000000000000000000324),
		Input: calc.AddStruct{
			X: big.NewFloat(0.000000000000000000321),
			Y: big.NewFloat(0.000000000000000000003),
		},
	},
}

func TestAdd(t *testing.T) {

	for _, tt := range tests {
		result := calc.Add(&tt.Input)
		if tt.Result.String() != result {
			t.Errorf("got %v, want %v", result, tt.Result)
		}
	}
}
