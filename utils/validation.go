package utils

import (
	"strconv"
)

// IsNumeric check string is number
func IsNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}
