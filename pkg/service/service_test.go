package service

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddTwoNumbers(t *testing.T) {
	tt := []struct {
		n1     *big.Float
		n2     *big.Float
		result *big.Float
	}{
		{
			n1:     big.NewFloat(1.0),
			n2:     big.NewFloat(2.0),
			result: big.NewFloat(3.0),
		},
		{
			n1:     big.NewFloat(-1.0),
			n2:     big.NewFloat(-2.0),
			result: big.NewFloat(-3.0),
		},
		{
			n1:     big.NewFloat(0.0),
			n2:     big.NewFloat(-1.0),
			result: big.NewFloat(-1.0),
		},
	}

	for _, tc := range tt {
		result := AddTwoNumbers(tc.n1, tc.n2)
		require.Equal(t, result, tc.result, "The sum of n1 and n2 should be equal to result")
	}
}
