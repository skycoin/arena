package main

import (
	"os"

	"github.com/artem-karpenko/arena"
)

func main() {
	const numbersToRead = 2
	arena.ReadSumAndPrint(os.Stdin, os.Stdout, numbersToRead)
}
