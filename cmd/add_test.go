package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 3.0, Add(1, 2), "they should be equal")
	assert.Equal(t, 43.5, Add(11.5, 32), "they should be equal")
}
