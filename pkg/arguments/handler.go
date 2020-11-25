package arguments

import (
	"errors"
	"strconv"
)

func Handler(args []string) ([]float64, error) {
	var returnValue []float64
	if len(args) < 2 {
		return returnValue, errors.New("You must give me 2 numbers")
	}

	arg1, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return returnValue, errors.New("The first input is not a number")
	}

	arg2, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return returnValue, errors.New("The second input is not a number")
	}

	returnValue = []float64{arg1, arg2}
	return returnValue, nil
}
