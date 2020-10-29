package pkg

import (
	"fmt"
	"os"
	"strconv"
)

// AddTwoNums : adding numbers from CLI args
func AddTwoNums() (int, error) {

	// Handling panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()

	// Avoiding 0 index as it is the program name
	//taking CLI args
	argsLength := len(os.Args)

	// Converting string to int
	num1, err := strconv.Atoi(os.Args[argsLength-2])
	if err != nil {
		return 0, err
	}
	num2, err := strconv.Atoi(os.Args[argsLength-1])
	if err != nil {
		return 0, err
	}

	// Returning sum
	return num1 + num2, nil

}
