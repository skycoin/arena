package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skycoin/arena/pkg/adder"
	"github.com/skycoin/arena/pkg/version"
)

func main() {
	flVersion := flag.Bool("version", false, "show version and exit")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Ussage of %s [options] number1 number2 :\n", os.Args[0])
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *flVersion {
		version.Print()
		os.Exit(0)
	}

	a, err := adder.New(flag.Args())
	if err != nil {
		flag.Usage()
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Result: %s\n", a.Add())
}
