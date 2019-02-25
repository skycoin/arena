package arena

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	description string
	in          string
	out         string
}{
	{
		description: "ints",
		in:          "42 17",
		out:         "59",
	},
	{
		description: "floats",
		in:          "15.67 17.1",
		out:         "32.77",
	},
	{
		description: "newlines",
		in:          "7\n3.3",
		out:         "10.3",
	},
}

func TestSumNumbers(t *testing.T) {
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			writer := bytes.NewBufferString("")
			ReadSumAndPrint(strings.NewReader(test.in), writer, 2)
			assert.Equal(t, test.out, writer.String())
		})
	}
}
