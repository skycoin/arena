package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSlicedInterface(t *testing.T) {
	payload := []string{
		"5",
		"3",
	}
	expects := []interface{}{
		"5",
		"3",
	}
	ui := NewUserInput(payload)
	assert.Equal(t, expects, ui.ToSlicedInterface())
}

func TestProcessableValuesToNumbers(t *testing.T) {
	payload := []string{
		"5",
		"3",
	}

	ui := NewUserInput(payload)

	assert.Equal(t, 5, ui.GetA())
	assert.Equal(t, 3, ui.GetB())
}