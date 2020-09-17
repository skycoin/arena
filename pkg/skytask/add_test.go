package skytask

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test decorator to decode slice output to match test case
func testDecoder(a []interface{}) []interface{} {
	return []interface{}{
		a[0].(float64),
		a[1].(error).Error(),
	}
}

// Test decorator for multi value return type
func toSlice(a ...interface{}) []interface{} {
	return a
}

func TestAdd2Num(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			"nil input",
			args{[]string{"main_file"}},
			[]interface{}{
				float64(0),
				"can only add two number neither less nor more",
			},
		},
		{
			"only one input",
			args{[]string{"main_file", "1"}},
			[]interface{}{
				float64(0),
				"can only add two number neither less nor more",
			},
		},
		{
			"more than two input",
			args{[]string{"main_file", "1", "2.4", "3"}},
			[]interface{}{
				float64(0),
				"can only add two number neither less nor more",
			},
		},
		{
			"first input malformed",
			args{[]string{"main_file", "xr", "2.4"}},
			[]interface{}{
				float64(0),
				"Ayyy mate can't read the first number",
			},
		},
		{
			"second input malformed",
			args{[]string{"main_file", "2", "xr"}},
			[]interface{}{
				float64(0),
				"Ayyy mate can't read the second number",
			},
		},
		{
			"float max limit",
			args{[]string{"main_file", fmt.Sprint(math.MaxFloat64), fmt.Sprint(math.MaxFloat64)}},
			[]interface{}{
				float64(0),
				"These numbers are too damn high",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, testDecoder(toSlice(Add2Num(tt.args.args...))))
		})
	}
}
