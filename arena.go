package main

import (
	"log"
	"os"
	"strconv"
)

func main()  {
	args := os.Args
	if len(args) != 3 {
		log.Printf("Incorrect number of params, %v instead of 2", len(args))
		return
	}

	sum(args[1], args[2])

}

func sum(first, second string) {
	firstFloat, err := strconv.ParseFloat(first, 64)
	if err != nil {
		log.Printf("Error parsing first param: %v", err)
		return
	}

	secondFloat, err := strconv.ParseFloat(second, 64)
	if err != nil {
		log.Printf("Error parsing second param: %v", err)
		return
	}

	log.Printf("Sum = %v", firstFloat + secondFloat)
}