package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gunjan01/arena/cmd"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Print an $ for a shell like effect.
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = executeCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func executeCommand(cmdString string) error {
	cmdString = strings.TrimSuffix(cmdString, "\n")
	args := strings.Fields(cmdString)

	switch args[0] {
	case "exit":
		os.Exit(0)
	case "add":
		{
			result, err := cmd.Add(args)
			if err != nil {
				return err
			}

			fmt.Fprintln(os.Stdout, result)
			return nil
		}
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
