package arguments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input []string
		err   string
	}{
		{[]string{}, "You must give me 2 numbers"},
		{[]string{"1"}, "You must give me 2 numbers"},
		{[]string{"a", "2"}, "The first input is not a number"},
		{[]string{"1", "a"}, "The second input is not a number"},
		{[]string{"1", "3"}, ""},
	}

	for _, test := range tests {
		_, err := Handler(test.input)
		if err != nil {
			assert.EqualError(err, test.err)
		}
	}
}
