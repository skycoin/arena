package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/haposan06/arena/pkg/services"

	"github.com/haposan06/arena/pkg/util"
)

const (
	ADD        = "add"
	ERR_FORMAT = "Incorrect format"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Printf(" %s -num1 NUM1 -num2 NUM2 \n", ADD)
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) add(num1, num2 float64) {
	result := services.Add(num1, num2)
	fmt.Printf("Total addition of %.2f and %.2f is %.2f", num1, num2, result)
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	addCmd := flag.NewFlagSet(ADD, flag.ExitOnError)

	num1 := addCmd.String("num1", "", "First number")
	num2 := addCmd.String("num2", "", "Second number")

	switch os.Args[1] {
	case ADD:
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addCmd.Parsed() {
		if !util.IsNumeric(*num1) && !util.IsNumeric(*num2) {
			addCmd.Usage()
			runtime.Goexit()
		}
		number1, err := strconv.ParseFloat(*num1, 64)
		if err != nil {
			fmt.Printf(" %s in arg %v", ERR_FORMAT, num1)
			addCmd.Usage()
			runtime.Goexit()
		}

		number2, err := strconv.ParseFloat(*num2, 64)
		if err != nil {
			fmt.Printf(" %s in arg %v", ERR_FORMAT, num2)
			addCmd.Usage()
			runtime.Goexit()
		}
		cli.add(number1, number2)
	}
}
