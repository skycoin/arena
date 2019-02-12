package add

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/skycoin/arena/utils"

	"github.com/hashicorp/consul/command/flags"
	"github.com/mitchellh/cli"
)

// New create cmd
func New(ui cli.Ui) *cmd {
	c := &cmd{UI: ui}
	c.init()
	return c
}

// cmd struct
type cmd struct {
	UI    cli.Ui
	flags *flag.FlagSet
	help  string
}

// cmd initialization
func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	c.help = flags.Usage(help, c.flags)
}

// Run add command
func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); err != nil {
		return 1
	}

	if err := process(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	return 0
}

// process function
func process(args []string) error {

	// check input parameter
	if len(args) < 2 {
		return fmt.Errorf("ERROR: Invalid input parameters! \n %v", help)
	}

	// numeric validation
	if !utils.IsNumeric(args[0]) || !utils.IsNumeric(args[1]) {
		return fmt.Errorf("ERROR: Invalid number! \n %v", help)
	}

	// assign number #1
	number1, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return err
	}

	// assign number #2
	number2, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return err
	}

	// test add function
	fmt.Printf(" %v + %v = %.2f \n", number1, number2, number1+number2)

	return nil
}

// Synopsis show command synopsis
func (c *cmd) Synopsis() string {
	return synopsis
}

// Help show help instruction
func (c *cmd) Help() string {
	return c.help
}

const synopsis = "Runs add."
const help = `
Usage: 
   arena add <number1> <number2>

Sample: 
   arena add 1 2
`
