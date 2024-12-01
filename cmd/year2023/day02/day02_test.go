package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		input    string
		expected Game
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Game{ID: 1, Blue: []int{3, 6}, Red: []int{4, 1}, Green: []int{2, 2}},
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Game{ID: 2, Blue: []int{1, 4, 1}, Red: []int{1}, Green: []int{2, 3, 1}},
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Game{ID: 3, Blue: []int{6, 5}, Red: []int{20, 4, 1}, Green: []int{8, 13, 5}},
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Game{ID: 4, Blue: []int{6, 15}, Red: []int{3, 6, 14}, Green: []int{1, 3, 3}},
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Game{ID: 5, Blue: []int{1, 2}, Red: []int{6, 1}, Green: []int{3, 2}},
		},
	}

	for i, tt := range tests {
		got, err := parseLine(tt.input)

		if got == nil {
			t.Fatalf("Missing response for line %d: %v", i, err)
		}

		if !reflect.DeepEqual(*got, tt.expected) {
			t.Fatalf("got %+v, expected %+v", got, tt.expected)
		}
	}
}
