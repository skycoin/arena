package skytask

import (
	"errors"
	"math"
	"strconv"
)

// Add2Num - Take variadic arguments and two numbers
func Add2Num(args ...string) (sum float64, err error) {
	if len(args) != 3 {
		err = errors.New("can only add two number neither less nor more")
		return
	}

	x, err := strconv.ParseFloat(args[1], 64)
	if nil != err {
		err = errors.New("Ayyy mate can't read the first number")
		return
	}

	y, err := strconv.ParseFloat(args[2], 64)
	if nil != err {
		err = errors.New("Ayyy mate can't read the second number")
		return
	}

	if x == math.MaxFloat64 || y == math.MaxFloat64 {
		err = errors.New("These numbers are too damn high")
		return
	}

	sum = x + y
	return
}
