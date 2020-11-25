package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	res, err := doubleInput(args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func doubleInput(input []string) (float64, error) {
	if len(input) < 3 {
		return 0, errors.New("not enough cmdline args")
	}

	first, err := strconv.ParseFloat(input[1],32)
	if err != nil {
		return 0, err
	}

	second, err := strconv.ParseFloat(input[2],32)
	if err != nil {
		return 0, err
	}
	res := first + second
	return res, nil
}
