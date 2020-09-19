package skycoinArenaTask

import (
	"fmt"
	"math"
	"testing"
)

type Args struct {
	args []string
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		name string
		args Args
		want []interface{}
	}{
		{
			"no arguments",
			Args{[]string{"main"}},
			[]interface{}{
				float64(-1),
				"add function accepts 2 or more numbers",
			},
		},
		{
			"sum of 3 numbers",
			Args{[]string{"main", "2", "1.3", "3"}},
			[]interface{}{
				float64(6.3),
				"",
			},
		},
		{
			"malformed input",
			Args{[]string{"main", "22", "2.4 1"}},
			[]interface{}{
				float64(-1),
				"error parsing numbers provided",
			},
		},
		{
			"very big number",
			Args{[]string{"main", fmt.Sprintf("%f", math.MaxFloat64), "23", "2"}},
			[]interface{}{
				float64(-1),
				"provided number too high for addition",
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sum, err := Add(testCase.args.args...)
			if sum != testCase.want[0].(float64) {
				t.Fail()
			}
			if err != nil && err.Error() != testCase.want[1].(string) {
				t.Fail()
			}
		})
	}
}
