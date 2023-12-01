package main

import (
	"strings"
	"testing"
)

func TestFindFirstAndLastNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected [2]string
	}{
		{"1abc2", [2]string{"1", "2"}},
		{"pqr3stu8vwx", [2]string{"3", "8"}},
		{"a1b2c3d4e5f", [2]string{"1", "5"}},
		{"treb7uchet", [2]string{"7", "7"}},
	}

	for _, test := range tests {
		got, _ := findFirstAndLastNumber(strings.Split(test.input, ""))
		if got != test.expected {
			t.Fatalf("wrong output. got %v expected %v", got, test.expected)
		}
	}
}
