package add

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// addtests are test cases for add program
var addtests = []struct {
	input  Data
	result *big.Float
}{
	{Data{
		Numbers: []*big.Float{big.NewFloat(45.5), big.NewFloat(23.2)},
	}, big.NewFloat(68.7)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(4), big.NewFloat(2.2)},
	}, big.NewFloat(6.2)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(22), big.NewFloat(23)},
	}, big.NewFloat(45)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(4.5), big.NewFloat(2)},
	}, big.NewFloat(6.5)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(-22), big.NewFloat(23)},
	}, big.NewFloat(1)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(4), big.NewFloat(-6)},
	}, big.NewFloat(-2)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(-22), big.NewFloat(-23)},
	}, big.NewFloat(-45)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(-4.5), big.NewFloat(2)},
	}, big.NewFloat(-2.5)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(-4.5), big.NewFloat(99)},
	}, big.NewFloat(94.5)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(0), big.NewFloat(0)},
	}, big.NewFloat(0)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(0), big.NewFloat(2)},
	}, big.NewFloat(2)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(0), big.NewFloat(-2)},
	}, big.NewFloat(-2)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(0), big.NewFloat(-2.5)},
	}, big.NewFloat(-2.5)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(0), big.NewFloat(2.5)},
	}, big.NewFloat(2.5)},
	{Data{
		Numbers: []*big.Float{big.NewFloat(0.000000000000000000003), big.NewFloat(0.000000000000000000321)},
	}, big.NewFloat(3.24e-19)},
}

func TestSums(t *testing.T) {
	for _, tt := range addtests {
		t.Run("test1", func(t *testing.T) {
			s := tt.input.Add()
			assert.Equal(t, s.Cmp(tt.result), 0, "The two values should be the same.")
		})
	}
}
