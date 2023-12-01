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

func TestFindFirstAndLastNumberAsLettersToo(t *testing.T) {
	tests := []struct {
		input    string
		expected [2]string
	}{
		{"two1nine", [2]string{"2", "9"}},
		{"eightwothree", [2]string{"8", "3"}},
		{"abcone2threexyz", [2]string{"1", "3"}},
		{"xtwone3four", [2]string{"2", "4"}},
		{"4nineeightseven2", [2]string{"4", "2"}},
		{"zoneight234", [2]string{"1", "4"}},
		{"7pqrstsixteen", [2]string{"7", "6"}},
	}

	for _, test := range tests {
		got, _ := findFirstAndLastNumberAsLettersToo(strings.Split(test.input, ""))
		if got != test.expected {
			t.Fatalf("wrong output. got %v expected %v", got, test.expected)
		}
	}
}
