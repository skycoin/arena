package arena

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// captureOutput allows us to capture the string written to stdout
func captureOutput(f func(...float64), args []float64) (string, error) {
	reader, writer, err := os.Pipe()
	if err != nil {
		return "", err
	}

	stdout := os.Stdout
	stderr := os.Stderr

	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer

	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		str := buf.String()
		out <- str[0 : len(str)-1]
	}()

	wg.Wait()
	f(args...)
	writer.Close()
	return <-out, nil
}

var tests = []struct {
	testCase string
	in       []float64
	out      string
}{
	{
		testCase: "ints",
		in:       []float64{2, 4},
		out:      "6",
	}, {
		testCase: "floats",
		in:       []float64{3.14, 3.34},
		out:      "6.48",
	}, {
		testCase: "int-float",
		in:       []float64{3, 2.33},
		out:      "5.33",
	},
}

func TestPrintSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			actual, err := captureOutput(PrintSum, tt.in)
			if err != nil {
				panic(err)
			}
			expected := tt.out
			assert.Equal(t, actual, expected, "The two outputs should be the same")
		})
	}
}
