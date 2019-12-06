package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const usage = "usage:\n\t [number] [number]"

func main() {
	validateArgs()
	value, err := sum(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatalf("%v\n%v", err, usage)
	}

	fmt.Println(value)
}

func sum(a, b interface{}) (float64, error) {
	i, err := parseFloat(a)
	if err != nil {
		return 0, err
	}

	j, err := parseFloat(b)
	if err != nil {
		return 0, err
	}

	return i + j, nil
}

func validateArgs() {
	args := len(os.Args)
	switch args {
	case 1:
		log.Fatalf("Program requires 2 arguments, but none found\n%v", usage)
	case 2:
		log.Fatalf("Program requires 2 arguments, only one entered\n%v", usage)
	}
}

func parseFloat(number interface{}) (float64, error) {
	val, _ := number.(string)
	arg, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return arg, errors.New("Invalid argument " + val)
	}

	return arg, nil
}
