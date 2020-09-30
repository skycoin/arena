package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	numberOfInputArguments = 2
	separator              = " "
)

func main() {
	for {
		sumVariablesFromCLI()
	}
}

// function tries to read arguments from command line, it'll return an error if input data format is incorrect
func readArgumentsFromCLI() (numbers []string, err error) {
	fmt.Printf("Input format with space is \"number_1 number_1_2\", for example, \"3 -53\"\nEnter %v numbers, please: ", numberOfInputArguments)

	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return strings.Split(inputString, separator), nil
}

// function tries to sum arguments from command line after successful reading
func sumVariablesFromCLI() float64 {
	arguments, err := readArgumentsFromCLI()
	if err != nil {
		fmt.Print(err)
		return 0
	}

	numbers, err := parseStringToFloat(arguments)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	sum := sumFloatVariables(numbers)
	fmt.Printf("Sum of your numbers is: %v\n\n", sum)
	return sum
}

// function tries to parse inputStrings to float, it'll return an error if input strings format is incorrect
func parseStringToFloat(inputStrings []string) (numbers []float64, err error) {
	arguments := make([]float64, len(inputStrings))

	if len(inputStrings) >= numberOfInputArguments {
		for i := 0; i < numberOfInputArguments; i++ {
			stringArg := strings.TrimSpace(inputStrings[i])
			number, err := strconv.ParseFloat(stringArg, 64)
			if err != nil {
				err := errors.New("Input data format is incorrect, please try again\n\n")
				return nil, err
			}
			arguments = append(arguments, number)
		}
	} else {
		err := errors.New("Wrong amount of arguments, please try again\n\n")
		return nil, err
	}

	return arguments, nil
}

// function sums all arguments from input slice
func sumFloatVariables(numbers []float64) (sum float64) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}
