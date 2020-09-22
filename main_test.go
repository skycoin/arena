package main

import (
	"testing"
)

func TestSumStrings(t *testing.T) {
	if result, _ := SumStrings("1", "2"); result != float64(3) {
		t.Errorf("SumStrings(\"%v\", \"%v\") failed, expected %f, got %v", "1", "2", float64(3), result)
	}

	if result, _ := SumStrings("-1.1", "2.5"); result != float64(1.4) {
		t.Errorf("SumStrings(\"%v\", \"%v\") failed, expected %f, got %v", "-1.1", "2.5", float64(1.4), result)
	}

	if _, err := SumStrings("a", "2"); err == nil {
		t.Errorf("SumStrings(\"%v\", \"%v\") failed, expected error", "a", "2")
	}

	if _, err := SumStrings("2", "a"); err == nil {
		t.Errorf("SumStrings(\"%v\", \"%v\") failed, expected error", "2", "a")
	}

	if _, err := SumStrings("a", "b"); err == nil {
		t.Errorf("SumStrings(\"%v\", \"%v\") failed, expected error", "a", "b")
	}

	if _, err := SumStrings("1", "2.3.4"); err == nil {
		t.Errorf("SumStrings(\"%v\", \"%v\") failed, expected error", "a", "b")
	}
}
