package services

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const numToRead = 2

var tables = []struct {
	description string
	inputs      string
	testNumbers []float64
	sum         float64
}{
	{
		description: "Integer",
		inputs:      "120\n4",
		testNumbers: []float64{120, 4},
		sum:         124,
	},
	{
		description: "Floats",
		inputs:      "15.1\n2.3",
		testNumbers: []float64{15.1, 2.3},
		sum:         17.4,
	},
	{
		description: "Integers with Spaces",
		inputs:      "15 2",
		testNumbers: []float64{15, 2},
		sum:         17,
	},
}

var failTables = []struct {
	description string
	inputs      string
}{
	{
		description: "Strings",
		inputs:      "a\nb",
	},
	{
		description: "Strings with spaces",
		inputs:      "a b",
	},
}

func TestReadNumbers(t *testing.T) {
	for _, tt := range tables {
		t.Run(tt.description, func(t *testing.T) {

			assert := assert.New(t)
			numbers, err := ReadNumbers(strings.NewReader(tt.inputs), numToRead)

			assert.Nil(err)
			assert.Equal(tt.testNumbers[0], numbers[0], "They should be equal")
			assert.Equal(tt.testNumbers[1], numbers[1], "They should be equal")

		})
	}
}

func TestAddNumbers(t *testing.T) {
	for _, tt := range tables {
		t.Run(tt.description, func(t *testing.T) {

			numbers, _ := ReadNumbers(strings.NewReader(tt.inputs), numToRead)
			sum := AddNumbers(numbers)
			assert.Equal(t, tt.sum, sum, "They should be equal")

		})
	}
}

func TestFailInputs(t *testing.T) {
	for _, tt := range failTables {
		t.Run(tt.description, func(t *testing.T) {

			_, err := ReadNumbers(strings.NewReader(tt.inputs), numToRead)
			assert.Nil(t, err)

		})
	}
}
