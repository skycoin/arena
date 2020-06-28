package parser

import (
	"errors"
	"strconv"
)

var ErrWrongArgumentsCount = errors.New("wrong count of arguments")
var ErrFirstArgumentIsNotNumber = errors.New("first argument is not a number")
var ErrSecondArgumentIsNotNumber = errors.New("second argument is not a number")

type argumentsParser struct{}

func NewArgumentsParser() ArgumentsParser {
	return &argumentsParser{}
}

func (ap *argumentsParser) Take2Numbers(args []string) (float64, float64, error) {
	if len(args) != 2 {
		return 0, 0, ErrWrongArgumentsCount
	}

	n1, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, 0, ErrFirstArgumentIsNotNumber
	}
	n2, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return 0, 0, ErrSecondArgumentIsNotNumber
	}

	return n1, n2, nil
}
