package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeOperation(t *testing.T) {
	input := "25 45\n"
	reader := bufio.NewReader(strings.NewReader(input))

	gotNum1, gotNum2, gotRes, err := MakeOperation(reader)
	assert.NoError(t, err)
	assert.Equal(t, float64(70), gotRes)
	assert.Equal(t, float64(25), gotNum1)
	assert.Equal(t, float64(45), gotNum2)

	input = "ert 45\n"
	reader = bufio.NewReader(strings.NewReader(input))

	_, _, _, err = MakeOperation(reader)
	assert.Error(t, err)

}
