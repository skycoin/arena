package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateInput(t *testing.T) {
	validator1 := Validator{FloatValue: 0,
		ErrMsg: ErrorNotNumber}
	testValidator1 := ValidateInput("testMsg")
	require.Equal(t, validator1, testValidator1, "These two value should be same.")

	validator2 := Validator{FloatValue: 0,
		ErrMsg: ErrorInvalid}
	testValidator2 := ValidateInput("1324.")
	require.Equal(t, validator2, testValidator2, "These two value should be same.")

	validator3 := Validator{FloatValue: 345.234,
		ErrMsg: ""}
	testValidator3 := ValidateInput("345.234")
	require.Equal(t, validator3, testValidator3, "These two value should be same.")
}
