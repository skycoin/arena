package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleInput(t *testing.T) {
	testCase := []struct {
		name     string
		args     []string
		is_valid bool
	}{
		{
			name:     "valid int",
			args:     []string{"", "5", "5"},
			is_valid: true,
		},
		{
			name:	"valid float",
			args: []string{"","5.15","5.15"},
			is_valid: true,
		},
		{
			name:     "invalid",
			args:     []string{},
			is_valid: false,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			_, err := doubleInput(tc.args)

			if tc.is_valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

		})
	}
}
