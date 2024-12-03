package main

import (
	"bufio"
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/akyrey/aoc/internal"
)

func main() {
	f, err := internal.GetFileToReadFrom(3, false)
	internal.CheckError(err)
	defer f.Close()

	total := 0
	reEnable := regexp.MustCompile(`do()`)
	reDisable := regexp.MustCompile(`don't()`)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	reNumbers := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(f)
	enabledFromPreviouseLine := true
	for scanner.Scan() {
		line := scanner.Text()
		multipliesIndexes := re.FindAllStringIndex(line, -1)
		enable := reEnable.FindAllStringIndex(line, -1)
		disable := reDisable.FindAllStringIndex(line, -1)
		for _, m := range multipliesIndexes {
			if !isEnabled(enable, disable, m[0], enabledFromPreviouseLine) {
				enabledFromPreviouseLine = false
				continue
			}
			enabledFromPreviouseLine = true
			numbers := reNumbers.FindAllString(line[m[0]:m[1]], 2)
			if len(numbers) != 2 {
				panic("Invalid input")
			}
			one, err := strconv.Atoi(numbers[0])
			internal.CheckError(err)
			two, err := strconv.Atoi(numbers[1])
			internal.CheckError(err)
			total += one * two
		}
		last := multipliesIndexes[len(multipliesIndexes)-1]
		enabledFromPreviouseLine = enableDisableNextLine(enable, disable, last[1], len(line), enabledFromPreviouseLine)
	}

	fmt.Printf("%d\n", total)
}

func isEnabled(enable [][]int, disable [][]int, index int, enabledFromPreviouseLine bool) bool {
	for i := index; i >= 0; i-- {
		if slices.ContainsFunc(enable, func(e []int) bool {
			return e[1] == i
		}) {
			return true
		}
		if slices.ContainsFunc(disable, func(e []int) bool {
			return e[1] == i
		}) {
			return false
		}
	}
	return enabledFromPreviouseLine
}

func enableDisableNextLine(enable [][]int, disable [][]int, index int, end int, enabledFromPreviouseLine bool) bool {
	for i := index; i < end; i++ {
		if slices.ContainsFunc(enable, func(e []int) bool {
			return e[0] == i
		}) {
			return true
		}
		if slices.ContainsFunc(disable, func(e []int) bool {
			return e[0] == i
		}) {
			return false
		}
	}
	return enabledFromPreviouseLine
}
