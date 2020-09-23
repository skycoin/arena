package additionTask

import (
	"testing"
)

var testCases = []struct {
	name string
	args []string
	res  []interface{}
}{
	{
		"no arguments",
		[]string{"1"},
		[]interface{}{
			float64(0),
			"Please input 2 numbers to know the sum",
		},
	},
	{
		"passing 3 arguments",
		[]string{"1", "2", "5.9"},
		[]interface{}{
			float64(0),
			"Input size exceeded! Please give 2 numbers only",
		},
	},
	{
		"error in input conversion",
		[]string{"1.0", "1e"},
		[]interface{}{
			float64(0),
			"Error while parsing the number",
		},
	},
	{
		"sum of two numbers",
		[]string{"4", "5.5"},
		[]interface{}{
			float64(9.5),
			"nil",
		},
	},
}

func TestAddition(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sum, err := AdditionOfTwoNums(testCase.args)
			if sum != testCase.res[0].(float64) {
				t.Fail()
			}
			if err != nil && err.Error() != testCase.res[1].(string) {
				t.Fail()
			}
		})
	}
}
