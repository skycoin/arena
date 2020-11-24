package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	testData := []struct {
		name   string
		num1   float64
		num2   float64
		result float64
	}{
		{
			name:   "Both positive",
			num1:   1.2,
			num2:   2,
			result: 3.2,
		},
		{
			name:   "One positive one negative",
			num1:   1.5,
			num2:   -1,
			result: 5,
		},
		{
			name:   "Both negative",
			num1:   -1.2,
			num2:   -1.5,
			result: -2.7,
		},
	}

	for _, tc := range testData {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.result, add(tc.num1, tc.num2))
		})
	}
}
