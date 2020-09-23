package additionTask

import (
	"errors"
	"strconv"
)

// AdditionOfTwoNums calculates sum of two numbers and
//returns the sum and error
func AdditionOfTwoNums(args []string) (float64, error) {
	var sum float64
	argsLength := len(args)
	if argsLength < 2 {
		return 0, errors.New("Please input 2 numbers to know the sum")
	} else if argsLength > 2 {
		return 0, errors.New("Input size exceeded! Please give 2 numbers only")
	} else {
		num1, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return 0, errors.New("Error while parsing the number")
		}
		num2, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			return 0, errors.New("Error while parsing the number")
		}
		sum = num1 + num2
	}
	return sum, nil
}
