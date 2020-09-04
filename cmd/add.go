/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var helpText = `
This command adds the arguments passed following it, Please remember to leave space between arguments!
Usage: arena add arg1 arg2 arg3 ...
	   arena add arg1 arg2 arg3 -f

`

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Commnad used to perform addition",
	Long:  helpText,
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		flag, _ := cmd.Flags().GetBool("float")
		if flag {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

// function to add all integers parsed from the command line
func addInt(args []string) {
	var sum int
	var added []int

	if len(args) == 0 {
		fmt.Println("Please pass some arguments to add")
		fmt.Printf(helpText)
		return
	}

	if len(args) == 1 {
		fmt.Println("WARNING: Please remember to add space between arguments or they will be parsed as one!")
	}

	for _, iVal := range args {

		// check if the arguments are all strings, and print an error
		// if any non-strings are encountered
		tmpInt, err := strconv.Atoi(iVal)

		if err != nil {
			fmt.Println("error:")
			fmt.Printf("%s while parsing integers, omiting value : %s \n", err.Error(), iVal)
			continue
		}

		added = append(added, tmpInt)
		sum = sum + tmpInt
	}

	if len(added) == 0 {
		fmt.Printf("\n No valid integer arguments found, please check the input \n use the -f flag to add float")
		fmt.Printf(helpText)
		return
	}

	fmt.Printf("Sum of given numbers %s is %d \n", args, sum)

}

func addFloat(args []string) {
	var sum float64
	var added []float64

	if len(args) == 0 {
		fmt.Println("Please pass some arguments to add")
		fmt.Printf(helpText)
		return
	}

	if len(args) == 1 {
		fmt.Println("WARNING: Please remember to add space between arguments or they will be parsed as one!")
	}

	for _, fVal := range args {

		// check if the arguments are all strings, and print an error
		// if any non-strings are encountered
		tmpFlt, err := strconv.ParseFloat(fVal, 64)

		if err != nil {
			fmt.Println("error:")
			fmt.Printf("%s while parsing floating point, omiting value : %s \n", err.Error(), fVal)
			continue
		}

		added = append(added, tmpFlt)
		sum = sum + tmpFlt
	}

	fmt.Printf("Sum of given float numbers %s is %f \n", args, sum)
}
