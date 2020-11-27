// Adds two numbers from command line
package main

import (
	"fmt"
	"os"

	"github.com/skycoin/arena/pkg"
)

func parseTwo(s1, s2 string) (float64, float64, error) {
	output1, err := pkg.Parse(s1)
	if err != nil {
		return 0, 0, err
	}
	output2, err := pkg.Parse(s2)
	if err != nil {
		return 0, 0, err
	}
	return output1, output2, nil
}

func useArgs(args []string) (float64, error) {
	if (len(args)) < 3 {
		return 0, fmt.Errorf("Expected at least two arguments, got %v.", len(args))
	}
	n1, n2, err := parseTwo(args[1], args[2])
	if err != nil {
		return 0, err
	}
	return pkg.Add(n1, n2)
}

func main() {
	output, err := useArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		// No nice way of not killing the process but still providing the status.
		os.Exit(1)
		return
	}
	fmt.Printf("Your result is %v.\n", output)
}
