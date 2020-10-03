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
		fmt.Println("\nPress CTRL + C to exit \n OR")
		sumVariablesFromCLI()
	}
}

// reads input from the CLI, notifies about any error
func readArgumentsFromCLI() (numbers []float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input format (space between 2 nums) is \"num1 num2\"\n")
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
				fmt.Printf("Input format incorrect, please try again\n")
				return nil, err
			}
			arguments = append(arguments, number)
		}
	} else {
		fmt.Printf("Input format incorrect, please try again\n")
		return nil, err
	}

	return arguments, nil
}

// calls summing function (sumNums) if the input data read is succesfull
func sumVariablesFromCLI() {
	numbers, err := readArgumentsFromCLI()
	if err != nil {
		return
	}
	sum := sumNums(numbers)
	fmt.Printf("Sum of your numbers is: %v\n\n", sum)
}

// sums all values in input slice
func sumNums(numbers []float64) (sum float64) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}
