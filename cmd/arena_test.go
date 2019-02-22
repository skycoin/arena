package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var flagtests = []struct {
	in  string
	out string
}{
	{"1 2\n", "3\n"},
	{"1 2.2\n", "3.2\n"},
	{"abc 2\n", "0\n"},
	{"abc def\n", "0\n"},
}

func captureStdout(reader io.Reader) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Add2Number(reader)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestAdd2Number(t *testing.T) {
	for _, tt := range flagtests {
		reader := struct{ io.Reader }{strings.NewReader(tt.in)}
		message := captureStdout(reader)
		require.Equal(t, message, tt.out, "Should be the same.")
	}
}
