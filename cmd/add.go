package cmd

import (
	"strconv"
)

var errMessage = "Cannot convert  number"
var errExitMessage = "please input two number"

func Add(args []string) string {
	if len(args) <= 2 {
		return errExitMessage

	}
	firstNumber, err := strconv.Atoi(args[1])
	if err != nil {
		return errMessage
	}
	secondNumber, err := strconv.Atoi(args[2])
	if err != nil {
		return errMessage
	}

	return strconv.Itoa(firstNumber + secondNumber)

}
