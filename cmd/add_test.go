// +build testrunmain

// This file allows us to run the entire program with test coverage enabled, useful for integration tests
package main

import (
	"testing"
)

func init() {
	parseFlags = false
}

func TestRunMain(t *testing.T) {
	main()
}
