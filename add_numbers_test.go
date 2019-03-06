package arena

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var additionTests = []struct {
	description  string
	input        string
	firstNumber  float64
	secondNumber float64
	sum          float64
}{
	{
		description:  "Integers",
		input:        "90\n2",
		firstNumber:  90,
		secondNumber: 2,
		sum:          92,
	},
	{
		description:  "Floats",
		input:        "19.75\n1.25",
		firstNumber:  19.2,
		secondNumber: 1.2,
		sum:          20.4,
	},
	{
		description:  "Integers (with space)",
		input:        "90 2",
		firstNumber:  90,
		secondNumber: 2,
		sum:          92,
	},
}

func TestReadNumbers(t *testing.T) {
	for _, test := range additionTests {
		t.Run(test.description, func(t *testing.T) {
			firstNumber, secondNumber, err := ReadNumbers(strings.NewReader(test.input))

			assert.Nil(t, err)
			assert.Equal(t, test.firstNumber, firstNumber)
			assert.Equal(t, test.secondNumber, secondNumber)
		})
	}
}

func TestAddNumbers(t *testing.T) {
	for _, test := range additionTests {
		t.Run(test.description, func(t *testing.T) {
			firstNumber, secondNumber, _ := ReadNumbers(strings.NewReader(test.input))
			sum := AddNumbers(firstNumber, secondNumber)
			assert.Equal(t, test.sum, sum)
		})
	}
}
