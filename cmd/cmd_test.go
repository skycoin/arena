package cmd

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_, testErr = strconv.ParseInt("invalid input", 10, 64)
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		input          []string
		expectedResult int64
		expectedError  error
	}{
		"basic input with two valid numbers": {
			input:          []string{"add", "1", "2"},
			expectedResult: 3,
		},
		"input with more than two numbers": {
			input:          []string{"add", "1", "2", "3"},
			expectedResult: 6,
		},
		"input with negative numbers": {
			input:          []string{"add", "1", "-2"},
			expectedResult: -1,
		},
		"input with less than two numbers": {
			input:          []string{"add", "1"},
			expectedResult: -1,
			expectedError:  errors.New("At least two numbers are required for performing addition"),
		},
		"invalid input": {
			input:          []string{"add", "invalid input", "-2"},
			expectedResult: -1,
			expectedError:  testErr,
		},
	}

	for _, test := range tests {
		result, err := Add(test.input)
		assert.Equal(t, result, test.expectedResult)
		assert.Equal(t, err, test.expectedError)
	}
}
