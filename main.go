package main

import (
	"os"

	"github.com/haposan06/arena/cmd"
)

func main() {
	defer os.Exit(0)
	cl := cmd.CommandLine{}
	cl.Run()
}
