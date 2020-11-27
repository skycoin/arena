package main

import (
	"arena/cmd"
	"fmt"
	"os"
)

func main() {
	fmt.Println(cmd.Add(os.Args))
}
