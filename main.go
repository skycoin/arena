package main

import (
	"fmt"

	app "github.com/offgridsito/skycoin-arena-cli/cmd/app"
	f "github.com/offgridsito/skycoin-arena-cli/pkg/sum"
)

func main() {
	app.ParseFlags()
	fmt.Println("Sum Result ==", f.SumNumbers(app.FirstAddend, app.SecondAddend))
}
