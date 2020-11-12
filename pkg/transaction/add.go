package transaction

import (
	"errors"
	"strconv"
)

// AddTwoNumber : add two number must be integer
func AddTwoNumber(x, y interface{}) (int, error) {

	numX, error := strconv.Atoi(x.(string))
	if error != nil {
		return 0, errors.New("parameter must be integer")
	}

	numY, error := strconv.Atoi(y.(string))
	if error != nil {
		return 0, errors.New("parameter must be integer")
	}

	return numX + numY, nil
}
