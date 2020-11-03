package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumUpNumber(t *testing.T) {
	var data = []interface{}{
		map[string]interface{}{"test": []int{}, "result": 0},
		map[string]interface{}{"test": []int{1, 2}, "result": 3},
		map[string]interface{}{"test": []int{1}, "result": 1},
		map[string]interface{}{"test": []int{1, 2, 3}, "result": 6},
	}
	for _, val := range data {
		testData := SumUpNumber(val.(map[string]interface{})["test"].([]int))
		expectedResult := val.(map[string]interface{})["result"]
		assert.Equal(t, testData, expectedResult, "The two words should be the same.")

	}
}
