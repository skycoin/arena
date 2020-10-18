package common

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

var ShouldBePairOfDigits =errors.New("we await a two digits")
var ShouldBeANumberValue = errors.New("value should be an number")
var PayloadValuesCantBeProcessed = errors.New("payload values can't be processed")

type Validable interface {
	isWithProperLength()
	isWithNumberValues()
	HasAnyError() bool
}

type Validator struct {
	payload []interface{}
	errors  []error
}

func NewValidator(payload []interface{}) Validator {
	 v := Validator{
		payload: payload,
		errors:  []error{},
	}
	v.isWithNumberValues()
	v.isWithProperLength()

	return v
}

func (v *Validator) HasAnyError() bool {
	if len(v.errors) > 0 {
		return true
	}

	return false
}

func (v *Validator) GetErrors() []error {
	return v.errors
}

func (v *Validator) isWithProperLength() {
	if len(v.payload) != 2 {
		v.errors = append(v.errors, ShouldBePairOfDigits)
	}
}

func (v *Validator) isWithNumberValues()  {
	for _, value := range v.payload {
		number, err := strconv.Atoi(fmt.Sprintf("%v", value))
		if err != nil {
			v.errors = append(v.errors, PayloadValuesCantBeProcessed)
		}
		if isNumber := unicode.IsNumber(rune(number)); isNumber != false {
			v.errors = append(v.errors, ShouldBeANumberValue)
		}
	}
}
