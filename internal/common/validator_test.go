package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewValidator(t *testing.T) {
	var payload = []interface{}{7,5}
	v := NewValidator(payload)
	assert.Equal(t, true, !v.HasAnyError())
}

func TestValidator_IsWithProperLength(t *testing.T) {
	payloadData := [][]interface{}{
		{7, 5, 2},
		{7},
		{5,3},
	}
	correctResults := []bool{
		false,
		false,
		true,
	}
	for i, payload := range payloadData {
		validator := NewValidator(payload)
		assert.Equal(t, correctResults[i], !validator.HasAnyError())
	}
}

func TestValidator_IsWithNumberValues(t *testing.T) {
	payloadData := [][]interface{}{
		{"7", "5"},
		{7, "o"},
		{5, 3},
	}
	correctResults := []bool{
		true,
		false,
		true,
	}
	for i, payload := range payloadData {
		validator := NewValidator(payload)
		assert.Equal(t, correctResults[i], !validator.HasAnyError())
	}
}