package main

import (
	"testing"
)

func TestNumbers(t *testing.T) {
	type data struct {
		first  float64
		second float64
	}
	tests := []struct {
		name   string
		data   data
		result float64
	}{
		{
			name: "Test negative numbers",
			data: data{
				first:  -11,
				second: -12,
			},
			result: -23,
		},
		{
			name: "Test float numbers",
			data: data{
				first:  11.2,
				second: 12,
			},
			result: 23.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := Inputs{tt.data.first, tt.data.second}

			if sum := input.Sum(); sum != tt.result {
				t.Errorf("Sum() = %v, result %v", sum, tt.result)
			}
		})
	}

}
