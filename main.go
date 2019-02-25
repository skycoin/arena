package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sysdevguru/skycoin/pkg"
)

func main() {
	var float1, float2 float64
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter First Number: ")
		inputString, _ := reader.ReadString('\n')
		validator := pkg.ValidateInput(strings.TrimSpace(inputString))
		if validator.ErrMsg != "" {
			fmt.Println(validator.ErrMsg)
		} else {
			float1 = validator.FloatValue
			break
		}
	}

	for {
		fmt.Print("Enter Second Number: ")
		inputString, _ := reader.ReadString('\n')
		validator := pkg.ValidateInput(strings.TrimSpace(inputString))
		if validator.ErrMsg != "" {
			fmt.Println(validator.ErrMsg)
		} else {
			float2 = validator.FloatValue
			break
		}
	}

	fmt.Printf("Result: %f\n", float1+float2)
	os.Exit(3)
}
