package cmd

import (
	"errors"
	"strconv"
)

// Add adds two numbers and returns the result.
// It returns an error when it encounters an invalid argument
// while parsing the command string.
func Add(arrCommandStr []string) (int64, error) {
	if len(arrCommandStr) < 3 {
		return -1, errors.New("At least two numbers are required for performing addition")
	}

	var sum int64
	for i, arg := range arrCommandStr {
		// Skip the first argument as it's the
		// command to be excuted.
		if i == 0 {
			continue
		}

		// Convert string to integer and add it
		// to the final result.
		num, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return -1, err
		}
		sum += num
	}
	return sum, nil
}
