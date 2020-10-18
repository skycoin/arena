package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTwoArgs(t *testing.T) {
	assert.Equal(t, 12, AddTwoNumbers(7, 5))
}
