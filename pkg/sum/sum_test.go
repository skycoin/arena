package sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var int64tests = []struct {
	a   int64
	b   int64
	sum int64
}{
	{-1, -1, -2},
	{0, 0, 0},
	{1, 1, 2},
}

func TestInt64(t *testing.T) {
	for _, tt := range int64tests {
		t.Run(fmt.Sprintf("%d %d", tt.a, tt.b), func(t *testing.T) {
			require.Equal(t, tt.sum, Int64(tt.a, tt.b))
		})
	}
}
