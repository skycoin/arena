package main

import (
	"fmt"
	"os"

	"github.com/skycoin/arena/internal/calculator"
	"github.com/skycoin/arena/internal/common"
)

func main() {
	ui := common.NewUserInput(os.Args[1:])

	fmt.Println(calculator.AddTwoNumbers(ui.GetA(), ui.GetB()))
}
