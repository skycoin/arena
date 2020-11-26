package main

import (
	"errors"
	"os"

	"github.com/bobbiprathama/arena/src/helpers"
	"github.com/bobbiprathama/arena/src/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var flagNumber1 int
var flagNumber2 int

func SumTwoNumber() *cobra.Command {
	cliCommand := &cobra.Command{
		Use:   "sum",
		Short: "Sum two number and print the result when given two number as input from cli argument",
		Run:   sumTwoNumber,
	}

	cliCommand.Flags().IntVarP(&flagNumber1, "first-number", "", 0, "First number for sum operation")
	cliCommand.Flags().IntVarP(&flagNumber2, "second-number", "", 0, "Second number for sum operation")

	return cliCommand
}

func sumTwoNumber(cmd *cobra.Command, args []string) {

	// -- init usecase
	uc := usecase.NewNumberOperation()

	// -- execute operation
	resultSum, err := uc.Sum(helpers.IntToInt64(&flagNumber1), helpers.IntToInt64(&flagNumber2))
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(1)
	}
	if resultSum == nil {
		logrus.Error(errors.New("Sum operation has empty result"))
		os.Exit(1)
	}
	logrus.Infof("Sum operation result from %d + %d is %d", flagNumber1, flagNumber2, *resultSum)
}
