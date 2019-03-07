package cmd

import (
	"flag"
	"fmt"
	"log"

	"github.com/skycoin/arena/pkg/math"
)

// Parse parses the command-line flags
func Parse() {

	// Accept two values and store them into variables
	var first, second float64
	flag.Float64Var(&first, "first", 1.0, "accepts first number")
	flag.Float64Var(&second, "second", 1.0, "accepts second number")
	flag.Parse()

	// Add input values
	result := math.Add(first, second)

	// Print result
	_, err := fmt.Printf("The result is %v\n", result)
	if err != nil {
		log.Printf("Print error: %s", err)
	}

}
