package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Add returns the sum of two digits
func Add(a, b float64) float64 {
	return a + b
}

// rootCmd represents the base command when called without any subcommands
var addCommand = &cobra.Command{
	Use:   "add",
	Short: "test",
	Long:  "Type two numbers from the command line, adds them up and prints the result",
	Run: func(cmd *cobra.Command, args []string) {
		var a, b float64
		fmt.Println("Enter the two numbers")
		fmt.Scanf("%f", &a)
		fmt.Scanf("%f", &b)
		fmt.Printf("Sum of %v and %v is %v\n", a, b, Add(a, b))
	},
}

func Execute() {
	if err := addCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
