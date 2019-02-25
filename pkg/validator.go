package pkg

import (
	"strconv"
	"strings"
)

var (
	ErrorNotNumber = "Input Value has to be a number"
	ErrorInvalid   = "Input Value is an invalid number"
)

type Validator struct {
	FloatValue float64
	ErrMsg     string
}

func ValidateInput(inputValue string) Validator {
	convertedValue, err := strconv.ParseFloat(inputValue, 64)
	if err != nil {
		validator := Validator{FloatValue: 0,
			ErrMsg: ErrorNotNumber}
		return validator
	}
	splitedArr := strings.Split(inputValue, ".")
	if len(splitedArr) == 2 && splitedArr[1] == "" {
		validator := Validator{FloatValue: 0,
			ErrMsg: ErrorInvalid}
		return validator
	}
	validator := Validator{FloatValue: convertedValue,
		ErrMsg: ""}
	return validator
}
