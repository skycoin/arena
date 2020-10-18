package common

import (
	"fmt"
	"os"
	"strconv"
)

type UserInputHelper struct {
	rawValues []string
	a, b int
}

func NewUserInput(input []string) UserInputHelper {
	ui := UserInputHelper{
		rawValues: input,
		a:         0,
		b:         0,
	}

	ui.validate()

	ui.assignABValues()

	return ui
}

func (ui *UserInputHelper)validate() {
	validator := NewValidator(ui.ToSlicedInterface())
	if validator.HasAnyError() {
		fmt.Println(validator.GetErrors())
		os.Exit(1)
	}
}

func (ui *UserInputHelper) GetA() int  {
	return ui.a
}

func (ui *UserInputHelper) GetB() int  {
	return ui.b
}

func (ui *UserInputHelper) ToSlicedInterface() []interface{} {
	var data []interface{}
	for _, val := range ui.rawValues {
		data = append(data, val)
	}
	return data
}

func (ui *UserInputHelper) assignABValues() {
	payload := ui.ToSlicedInterface()
	var err error
	ui.a, err = strconv.Atoi(fmt.Sprintf("%v", payload[0]))
	if err != nil {
		panic(err)
	}
	ui.b, err = strconv.Atoi(fmt.Sprintf("%v", payload[1]))
	if err != nil {
		panic(err)
	}
}
