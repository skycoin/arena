package skycoinArenaTask

import (
	"errors"
	"math"
	"strconv"
)

var ArgumentsCountError = errors.New("add function accepts 2 or more numbers")
var MalformedNumberError = errors.New("error parsing numbers provided")
var NumberTooHighError = errors.New("provided number too high for addition")

// Add function takes dynamic number of arguments ( variadic ) and returns their sum.
func Add(nums ...string) (float64, error) {
	if len(nums) < 2 {
		return -1, ArgumentsCountError
	}
	var sum float64 = 0
	for _, num := range nums[1:] {
		parsedInt, parseError := strconv.ParseFloat(num, 64)
		if parseError != nil {
			return -1, MalformedNumberError
		}
		if parsedInt == math.MaxFloat64 {
			return -1, NumberTooHighError
		}
		sum += parsedInt
	}
	return sum, nil
}
