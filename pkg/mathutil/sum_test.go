package mathutil

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	type want struct {
		sum *big.Float
	}

	tt := []struct {
		name string
		args []*big.Float
		want want
	}{
		{
			name: "OK - sum 9",
			args: []*big.Float{
				big.NewFloat(1),
				big.NewFloat(3),
				big.NewFloat(5),
			},
			want: want{
				sum: big.NewFloat(9),
			},
		},
		{
			name: "OK - sum 0 (no args)",
			args: []*big.Float{},
			want: want{
				sum: big.NewFloat(0),
			},
		},
		{
			name: "OK - sum 1 (1 arg)",
			args: []*big.Float{
				big.NewFloat(1),
			},
			want: want{
				sum: big.NewFloat(1),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sum := Sum(tc.args...)
			require.Equal(t, tc.want.sum, sum)
		})
	}
}
