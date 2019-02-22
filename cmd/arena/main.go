package main

import (
	"log"
	"os"

	"github.com/skycoin/arena"
)

func main() {
	log.SetFlags(0)

	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatalf("Required at least 2 args")
	}
	result, err := arena.ExecInt(arena.Add2(arena.StrInt(args[0]), arena.StrInt(args[1])))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Result: %v", result)
}
