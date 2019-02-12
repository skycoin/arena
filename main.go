package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	myCli := cli.NewCLI("arena", "1.0.0")

	myCli.Args = os.Args[1:]

	myCli.Commands = map[string]cli.CommandFactory{}

	exitStatus, err := myCli.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
