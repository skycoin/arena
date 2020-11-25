package main

import (
	"os"
	"testing"
)

func TestRunMain(t *testing.T) {
	os.Args = []string{"", "13.98", "-1.241"}
	main()
}
