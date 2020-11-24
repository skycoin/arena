package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Sum recieves a range of values and returns their sum
// this method can accept variadic number of arguments if we eventually want to do so
func Sum(values ...string) (float64, error) {
	sum := 0.0
	for _, arg := range values {
		argFloat64, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return 0.0, fmt.Errorf("cannot convert %s to number", arg)
		}
		sum += argFloat64
	}
	return sum, nil
}

func main() {
	argsFromCommand := os.Args[1:]
	if len(argsFromCommand) < 2 {
		log.Println("To use this command, you must pass at least two number as argument.")
		return
	}
	// "Take two numbers from the command line"
	twoNumbers := argsFromCommand[:2]
	sum, err := Sum(twoNumbers...)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Printf("sum of %v: %f\n", strings.Join(twoNumbers, " and "), sum)
}
