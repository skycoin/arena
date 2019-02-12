package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/skycoin/arena/cmd/add"
)

func main() {
	myCli := cli.NewCLI("arena", "1.0.0")

	myCli.Args = os.Args[1:]

	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	myCli.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) { return add.New(ui), nil },
	}

	exitStatus, err := myCli.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
