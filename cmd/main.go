package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "arena",
	Short: "Skycoin screening test",
}

func init() {
	registerCommands()
}

func main() {
	Execute()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err.Error())
		os.Exit(1)
	}
}

func registerCommands() {
	rootCmd.AddCommand(SumTwoNumber())
}
