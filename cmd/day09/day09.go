package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/akyrey/aoc-2023/internal"
)

func scanLine(line string) []int {
	numbers := make([]int, 0)
	for _, number := range strings.Split(line, " ") {
		value, err := strconv.Atoi(number)
		internal.CheckError(err)
		numbers = append(numbers, value)
	}
	return numbers
}

func findPrevValue(sequence []int) int {
	if internal.Every(sequence, func(current int) bool {
		return current == 0
	}) {
		return 0
	}

	newSequence := make([]int, len(sequence)-1)
	for i := len(sequence) - 2; i >= 0; i-- {
		newSequence[i] = sequence[i+1] - sequence[i]
	}

	return newSequence[0] - findPrevValue(newSequence)
}

func main() {
	f, err := internal.GetFileToReadFrom(9, false)
	internal.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, scanLine(line))
	}

	sum := 0
	for _, sequence := range lines {
		sum += sequence[0] - findPrevValue(sequence)
	}

	fmt.Println(sum)
}
