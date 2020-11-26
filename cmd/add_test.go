package cmd

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestCLI(t *testing.T) {
	for _, test := range []struct {
		Args   []string
		Output string
	}{
		{
			Args:   []string{"arena", "add", "5", "7"},
			Output: "12\n",
		},
		{
			Args:   []string{"arena", "add", "0", "7"},
			Output: "7\n",
		},
		{
			Args:   []string{"arena", "add"},
			Output: fmt.Sprintf("Cannot add, Err: %v", errTwoArguments),
		},
		{
			Args:   []string{"arena", "add", "1"},
			Output: fmt.Sprintf("Cannot add, Err: %v", errTwoArguments),
		},
		{
			Args:   []string{"arena", "add", "a", "1"},
			Output: fmt.Sprintf("Cannot add, Err: %v", errCannotConvert),
		},
		{
			Args:   []string{"arena", "add", "2", "c"},
			Output: fmt.Sprintf("Cannot add, Err: %v", errCannotConvert),
		},
	} {
		t.Run("", func(t *testing.T) {
			os.Args = test.Args
			out = bytes.NewBuffer(nil)
			AddCommand()

			if actual := out.(*bytes.Buffer).String(); actual != test.Output {
				fmt.Println(actual, test.Output)
				t.Errorf("expected %s, but got %s", test.Output, actual)
				fmt.Printf("expected %s, but got %s", test.Output, actual)
			}
		})
	}
}
