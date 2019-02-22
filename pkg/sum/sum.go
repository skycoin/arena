package sum

// Int64 takes variable number of int64 and returns their sum.
func Int64(args ...int64) int64 {
	var sum int64

	for _, i := range args {
		sum += i
	}

	return sum
}
