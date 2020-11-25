package main

import (
	"addxandy/pkg/addculate"
	"addxandy/pkg/arguments"
	"fmt"
	"log"
	"os"
)

func main() {
	inputArguments := os.Args[1:]
	addculateData, err := arguments.Handler(inputArguments)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addculate.AddFunction(addculateData))
}
