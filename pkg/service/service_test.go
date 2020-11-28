package service

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumTwoNumbers(t *testing.T) {
	tt := []struct {
		n1     *big.Float
		n2     *big.Float
		result *big.Float
	}{
		{
			n1:     big.NewFloat(1111111111111111111111111111111.1),
			n2:     big.NewFloat(1111111111111111111111111111111.1),
			result: big.NewFloat(2222222222222222222222222222222.2),
		},
		{
			n1:     big.NewFloat(-1111111111111111111111111111111.1),
			n2:     big.NewFloat(-1111111111111111111111111111111.1),
			result: big.NewFloat(-2222222222222222222222222222222.2),
		},
	}

	for _, tc := range tt {
		result := SumTwoNumbers(tc.n1, tc.n2)
		require.Equal(t, result, tc.result)
	}
}
