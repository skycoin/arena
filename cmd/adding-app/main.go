package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatal("ERROR: the first and/or the second arguments are not specified")
	}

	a, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("ERROR: first argument is not an integer")
	}

	b, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal("ERROR: second argument is not an integer")
	}

	log.Println("Result: first argument + second argument == ", a+b)
}
