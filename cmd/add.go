package main

import (
	"flag"

	"github.com/skycoin/arena"
)

var (
	parseFlags = true
	var1       float64
	var2       float64
)

func main() {
	flag.Float64Var(&var1, "v1", 0, "Variable 1(type=float64) to be added")
	flag.Float64Var(&var2, "v2", 0, "Variable 2(type=float64) to be added")

	if parseFlags {
		flag.Parse()
	}

	arena.PrintSum(var1, var2)
}
