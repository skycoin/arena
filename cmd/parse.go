package cmd

import (
	"arena/pkg/math"
	"flag"
	"fmt"
	"log"
)

// Parse parses the command-line flags
func Parse() {

	// Accept two values and store them into variables
	var firstPtr, secondPtr float64
	flag.Float64Var(&firstPtr, "first", 1.0, "accepts first number")
	flag.Float64Var(&secondPtr, "second", 1.0, "accepts second number")
	flag.Parse()

	// Add input values
	result := math.Add(firstPtr, secondPtr)

	// Print result
	_, err := fmt.Printf("The result is %v\n", result)
	if err != nil {
		log.Printf("Print error: %s", err)
	}

}
