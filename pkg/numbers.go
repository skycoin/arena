package pkg

import (
	"errors"
	"math"
	"strconv"
)

func Parse(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func getAddError(addResult float64) error {
	if addResult == math.Inf(1) {
		return errors.New("Got positive infinity, numbers are too large.")
	}
	if addResult == math.Inf(-1) {
		return errors.New("Got negative infinity, numbers are too small.")
	}
	return nil
}

func Add(a, b float64) (float64, error) {
	result := a + b
	return result, getAddError(result)
}
