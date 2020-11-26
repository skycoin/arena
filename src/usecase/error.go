package usecase

import "errors"

var ErrorNilParam error = errors.New("invalid parameter, should have not nil value")

func NewNilParamError() error {
	return ErrorNilParam
}
