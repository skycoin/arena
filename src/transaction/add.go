package transaction

import "errors"

// AddTwoNumber : add two number must be integer
func AddTwoNumber(x, y interface{}) (int, error) {
	numX, ok := x.(int)
	if !ok {
		return 0, errors.New("parameter must be integer")
	}

	numY, ok := y.(int)
	if !ok {
		return 0, errors.New("parameter must be integer")
	}

	return numX + numY, nil
}
