package helpers

func IntPointer(num int) *int {
	return &num
}

func IntToInt64(num *int) *int64 {
	if num == nil {
		return nil
	}
	int64Val := int64(*num)
	return &int64Val
}
