package main

import (
	"github.com/artem-karpenko/arena"
	"os"
)

func main() {
	const numbersToRead = 2
	arena.ReadSumAndPrint(os.Stdin, os.Stdout, numbersToRead)
}
