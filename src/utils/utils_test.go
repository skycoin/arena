package utils

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBigParseFloat(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		bigFloat, _ := ParseBigFloat("123")
		assert.Equal(t, big.NewFloat(123).Cmp(bigFloat), 0, "The two values should be the same.")
	})
}
