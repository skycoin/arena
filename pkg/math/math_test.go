package math

import (
	"testing"
)

type testpair struct {
	values [2]string
	res    float64
}

var tests = []testpair{
	{[2]string{"1", "2"}, 3},
	{[2]string{"1.2", "2"}, 3.2},
	{[2]string{"-1", "2"}, 1},
	{[2]string{"0", "2"}, 2},
	{[2]string{"-1", "-2.1"}, -3.1},
}

func TestAddValues(t *testing.T) {
	for _, pair := range tests {
		v, err := AddNumbers(pair.values[0], pair.values[1])
		if err != nil {
			t.Error("unexpected err", err)
		}
		if v != pair.res {
			t.Error(
				"For", pair.values,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}
