package main

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name    string
		number1 string
		number2 string
	}{
		{"success", "2.3", "2.4"},
		{"errorString", "ab", "ba"},
		{"errorLess", "2.3", "2.4"},
		{"errorMore", "2.3", "2.4"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "errorLess" {
				os.Args = []string{"dummy", test.number1}
				sum, err := AddNumbers()
				if err == nil {
					t.Error("Error should have occurred")
				}
				if sum != 0 {
					t.Error("Sum should have been 0")
				}
				t.Log(sum)
			}
			if test.name == "errorMore" {
				os.Args = []string{"dummy", test.number1, test.number2, "233"}
				sum, err := AddNumbers()
				if err == nil {
					t.Error("Error should have occurred")
				}
				if sum != 0 {
					t.Error("Sum should have been 0")
				}
			}
			if test.name == "success" {
				os.Args = []string{"dummy", test.number1, test.number2}
				sum, err := AddNumbers()
				if err != nil {
					t.Error("Error with correct data")
				}
				t.Log(sum)
			}
			if test.name == "errorString" {
				os.Args = []string{"dummy", test.number1, test.number2}
				sum, err := AddNumbers()
				if err == nil {
					t.Error("Error should have occurred")
				}
				if sum != 0 {
					t.Error("Sum should have been 0")
				}
			}
		})
	}
}
