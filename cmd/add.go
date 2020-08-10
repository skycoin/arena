package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add values passed to function",
	Long: `Print the output of the addition of two numbers. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		flagStatus, _ := cmd.Flags().GetBool("float")
		if flagStatus { // if status is true, call addFloat
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

func addInt(args []string) {
	var sum int
	for _, value := range args {
		num, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + num
	}
	fmt.Printf("Addition of numbers is %d \n", sum)
}


func addFloat(args []string) {
	var sum float64
	for _, value := range args {
		num, err := strconv.ParseFloat(value, 64)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + num
	}
	fmt.Printf("Addition of numbers is %f \n", sum)
}

