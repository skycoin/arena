package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func init() {

	addCommand.Flags().String("firstNumber", "", "firstNumber is required for addition")
	addCommand.Flags().String("secondNumber", "", "secondNumber is required for addition")
	addCommand.MarkFlagRequired("firstNumber")
	addCommand.MarkFlagRequired("secondNumber")
}
//Command for addition
var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Addition of two numbers",
	Long:  "Takes a two numbers from command line and add them and print the results",
	Run: func(cmd *cobra.Command, args []string) {
		firstNumber := cmd.Flag("firstNumber")
		if firstNumber.Value.String() == "" {
			log.Fatal(errors.New("firstNumber is empty"))
			return
		}
		secondNumber := cmd.Flag("secondNumber")
		if secondNumber.Value.String() == "" {
			log.Fatal(errors.New("secondNumber is empty"))
			return
		}
		numOne, _ := strconv.Atoi(firstNumber.Value.String())
		numTwo, _ := strconv.Atoi(secondNumber.Value.String())
		fmt.Printf("The Sum of %v and %v is %v\n", numOne, numTwo, Add(numOne, numTwo))
	},
}

func Execute() {
	if err := addCommand.Execute(); err != nil {
		log.Println(err)
	}
}

func Add(numOne, numTwo int) int {
	return numOne + numTwo
}
