package app

import "flag"

var (
	FirstAddend  float64
	SecondAddend float64
)

func ParseFlags() {
	flag.Float64Var(&FirstAddend, "firstAddend", 0, "Input 1st (number) to sum")
	flag.Float64Var(&SecondAddend, "secondAddend", 0, "Input 2nd (number) to sum")
	flag.Parse()
}
