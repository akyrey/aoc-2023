package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/akyrey/aoc/internal"
)

var NUMBERS = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func getNumberFromChar(value string) *string {
	_, err := strconv.Atoi(value)
	if err == nil {
		return &value
	}

	return nil
}

func getNumberFromCharOrString(chars []string, index int, alsoString bool) *string {
	result := getNumberFromChar(chars[index])
	if result != nil {
		return result
	}

	if !alsoString {
		return nil
	}

	for key, value := range NUMBERS {
		if index+len(key) <= len(chars) && key == strings.Join(chars[index:index+len(key)], "") {
			return &value
		}
	}

	return nil
}

func findFirstAndLastNumber(chars []string, alsoString bool) ([2]string, error) {
	var first *string
	var last *string
	for i, j := 0, len(chars)-1; i < len(chars) && j >= 0; i, j = i+1, j-1 {
		if first == nil {
			first = getNumberFromCharOrString(chars, i, alsoString)
		}
		if last == nil {
			last = getNumberFromCharOrString(chars, j, alsoString)
		}

		if first != nil && last != nil {
			return [2]string{*first, *last}, nil
		}
	}

	return [2]string{}, errors.New("couldn't find 2 numbers in the string")
}

func main() {
	f, err := internal.GetFileToReadFrom(1, false)
	internal.CheckError(err)
	defer f.Close()

	values := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		numbers, err := findFirstAndLastNumber(chars, true)
		internal.CheckError(err)

		number, err := strconv.Atoi(fmt.Sprintf("%s%s", numbers[0], numbers[1]))
		internal.CheckError(err)
		values = append(values, number)
	}

	count := 0
	for _, value := range values {
		count += value
	}

	fmt.Printf("%d", count)
}
