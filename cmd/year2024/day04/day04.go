package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	"github.com/akyrey/aoc/internal"
)

func main() {
	f, err := internal.GetFileToReadFrom(4, false)
	internal.CheckError(err)
	defer f.Close()

	total := 0
	matrix := make([][]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	for i := range matrix {
		for j := range matrix[i] {
			total += wordFoundPart2(matrix, i, j)
		}
	}

	fmt.Printf("%d\n", total)
}

func wordFoundPart2(matrix [][]string, row int, col int) int {
	wordToSearch := []string{"MAS", "SAM"}
	count := 0
	if row+2 < len(matrix) && col+2 < len(matrix[row]) {
		word1 := matrix[row][col] + matrix[row+1][col+1] + matrix[row+2][col+2]
		word2 := matrix[row][col+2] + matrix[row+1][col+1] + matrix[row+2][col]
		if slices.Contains(wordToSearch, word1) && slices.Contains(wordToSearch, word2) {
			count += 1
		}
	}

	return count
}

func wordFoundPart1(matrix [][]string, row int, col int) int {
	wordToSearch := []string{"XMAS", "SAMX"}
	count := 0
	if col+3 < len(matrix[row]) {
		word := strings.Join(matrix[row][col:col+4], "")
		if slices.Contains(wordToSearch, word) {
			count += 1
		}
	}
	if row+3 < len(matrix) {
		word := matrix[row][col] + matrix[row+1][col] + matrix[row+2][col] + matrix[row+3][col]
		if slices.Contains(wordToSearch, word) {
			count += 1
		}
	}
	if row+3 < len(matrix) && col+3 < len(matrix[row]) {
		word := matrix[row][col] + matrix[row+1][col+1] + matrix[row+2][col+2] + matrix[row+3][col+3]
		if slices.Contains(wordToSearch, word) {
			count += 1
		}
	}
	if row-3 >= 0 && col+3 < len(matrix[row]) {
		word := matrix[row][col] + matrix[row-1][col+1] + matrix[row-2][col+2] + matrix[row-3][col+3]
		if slices.Contains(wordToSearch, word) {
			count += 1
		}
	}

	return count
}
