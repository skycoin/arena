package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	for _, test := range []struct {
		Args     []string
		Expected string
	}{
		{
			Args:     []string{"arena", "1", "2"},
			Expected: "3",
		},
		{
			Args:     []string{"arena", "0", "0"},
			Expected: "0",
		},
		{
			Args:     []string{"arena"},
			Expected: fmt.Sprintf(errExitMessage),
		},
		{
			Args:     []string{"arena", "1"},
			Expected: fmt.Sprintf(errExitMessage),
		},
		{
			Args:     []string{"arena", "a", "1"},
			Expected: fmt.Sprintf(errMessage),
		},
		{
			Args:     []string{"arena", "2", "c"},
			Expected: fmt.Sprintf(errMessage),
		},
	} {
		expected := test.Expected
		os.Args = test.Args

		actual := Add(os.Args)

		if actual != expected {
			t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
		}

	}

}
