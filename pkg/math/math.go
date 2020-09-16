package math

import (
	"fmt"
	"strconv"
	"strings"
)

//NumError is returned when somebody pass wrong values like "1.2s", "foo", etc
var NumError = fmt.Errorf("Please write 2 numbers which to add (as arguments) ")

//NumError is returned when somebody pass very big numbers. Max value is math.MaxFloat64 (1.797693134862315708145274237317043567981e+308)
var RangeError = fmt.Errorf("This tool cannot add such a big numbers ")

// AddNumbers adds two numbers, which could be passed as ints or floats ("1", "1.0")
func AddNumbers(firstArg, secondArg string) (float64, error) {
	firstNum, err := strconv.ParseFloat(firstArg, 64)
	if err != nil {
		//We can't compare err == strconv.ErrRange because err returns some additional info
		if strings.Contains(err.Error(), strconv.ErrRange.Error()) {
			return 0, RangeError
		}
		return 0, NumError
	}
	secondNum, err := strconv.ParseFloat(secondArg, 64)
	if err != nil {
		//We can't compare err == strconv.ErrRange because err returns some additional info
		if strings.Contains(err.Error(), strconv.ErrRange.Error()) {
			return 0, RangeError
		}
		return 0, NumError
	}
	return firstNum + secondNum, nil
}
