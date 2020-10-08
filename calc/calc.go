package calc

import (
	"math/big"
)

type AddStruct struct {
	X *big.Float
	Y *big.Float
}

func Add(obj *AddStruct) string {
	return new(big.Float).Add(obj.X, obj.Y).String()
}
