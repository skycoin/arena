package input

import (
	"fmt"
	"math/big"
)

const prec = 53

func Input(info string) *big.Float {
	var s string
	var number *big.Float
	var success bool
	for {
		fmt.Println(info)
		fmt.Scanln(&s)
		number, success = new(big.Float).SetPrec(prec).SetString(s)
		if !success {
			fmt.Println("Please Enter Numbers Only: ")
			continue
		} else {
			break
		}
	}
	return number
}
