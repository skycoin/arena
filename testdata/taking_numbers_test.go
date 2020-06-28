package testdata

import (
	"testing"

	"github.com/skycoin/arena/src/parser"
	"github.com/stretchr/testify/assert"
)

var argstests = []struct {
	in  []string
	out []interface{}
}{
	{
		in:  []string{"4", "5"},
		out: []interface{}{4.0, 5.0, nil},
	},
	{
		in:  []string{"4", "5", "6"},
		out: []interface{}{0.0, 0.0, parser.ErrWrongArgumentsCount},
	},
	{
		in:  []string{},
		out: []interface{}{0.0, 0.0, parser.ErrWrongArgumentsCount},
	},
	{
		in:  []string{"4"},
		out: []interface{}{0.0, 0.0, parser.ErrWrongArgumentsCount},
	},
	{
		in:  []string{"4", "5", "6"},
		out: []interface{}{0.0, 0.0, parser.ErrWrongArgumentsCount},
	},
	{
		in:  []string{"4r", "5"},
		out: []interface{}{0.0, 0.0, parser.ErrFirstArgumentIsNotNumber},
	},
	{
		in:  []string{"4", "5t"},
		out: []interface{}{0.0, 0.0, parser.ErrSecondArgumentIsNotNumber},
	},
}

func TestGettingNumbers(t *testing.T) {
	argParser := parser.NewArgumentsParser()
	for _, test := range argstests {
		n1, n2, err := argParser.Take2Numbers(test.in)
		resp := []interface{}{n1, n2, err}
		assert.Equal(t, test.out, resp, "Taking numbers is failed.")
	}
}
