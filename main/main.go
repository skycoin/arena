package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	numberOfInputArguments = 2
)

func main() {
	for {
		sumVariablesFromCLI()
	}
}

// function tries to read arguments from command line, it'll print error if input data format is wrong
func readArgumentsFromCLI() (numbers []float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input format with space is \"number_1 number_1_2\", for example, \"3 -53\"\nEnter %v numbers, please: ", numberOfInputArguments)
	inputString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("%v\n\n", err)
		return nil, err
	}

	stringArgs := strings.Split(inputString, " ")
	arguments := make([]float64, len(stringArgs))

	if len(stringArgs) >= numberOfInputArguments {
		for i := 0; i < numberOfInputArguments; i++ {
			stringArg := strings.TrimSpace(stringArgs[i])
			number, err := strconv.ParseFloat(stringArg, 64)
			if err != nil {
				fmt.Printf("Input data format is incorrect, please try again\n\n")
				return nil, err
			}
			arguments = append(arguments, number)
		}
	} else {
		fmt.Printf("Input data format is incorrect, please try again\n\n")
		return nil, err
	}

	return arguments, nil
}

// function tries to sum arguments from command line after successful reading
func sumVariablesFromCLI() {
	numbers, err := readArgumentsFromCLI()
	if err != nil {
		return
	}
	sum := sumFloatVariables(numbers)
	fmt.Printf("Sum of your numbers is: %v\n\n", sum)
}

// function sums all arguments from input slice
func sumFloatVariables(numbers []float64) (sum float64) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}
