package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/akyrey/aoc/internal"
)

func main() {
	f, err := internal.GetFileToReadFrom(5, false)
	internal.CheckError(err)
	defer f.Close()

	total := 0
	rules := make(map[int][]int, 0)
	updates := make([][]int, 0)
	scanningRules := true
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanningRules = false
			continue
		}

		if scanningRules {
			scanRules(rules, line)
		} else {
			updates = append(updates, scanUpdate(line))
		}
	}

	fmt.Printf("Rules: %v\n", rules)
	fmt.Printf("Updates: %v\n", updates)

	for _, update := range updates {
		if !isValid(rules, update) {
			for !isValid(rules, update) {
				update = fixOrdering(rules, update)
			}
			mid := len(update) / 2
			fmt.Printf("Fixed update: %v, mid: %d, value: %d\n", update, mid, update[mid])
			total += update[mid]
		}
	}

	fmt.Printf("%d\n", total)
}

func scanRules(rules map[int][]int, line string) {
	rule := strings.Split(line, "|")
	if len(rule) != 2 {
		panic("Invalid rule")
	}
	x, err := strconv.Atoi(rule[0])
	internal.CheckError(err)
	y, err := strconv.Atoi(rule[1])
	internal.CheckError(err)
	_, ok := rules[y]
	if !ok {
		rules[y] = make([]int, 0)
	}
	rules[y] = append(rules[y], x)
}

func scanUpdate(line string) []int {
	pages := make([]int, 0)
	page := strings.Split(line, ",")
	for _, p := range page {
		x, err := strconv.Atoi(p)
		internal.CheckError(err)
		pages = append(pages, x)
	}

	return pages
}

func isValid(rules map[int][]int, pages []int) bool {
	for i, page := range pages {
		if values, ok := rules[page]; ok {
			valuesPresent := internal.Filter(values, func(v int) bool {
				return slices.Contains(pages, v)
			})
			if len(valuesPresent) > 0 && !containsAll(pages[:i], valuesPresent) {
				// fmt.Printf("Page %d is invalid. Rules: %v prev pages: %v\n", page, rules[page], pages[:i])
				return false
			}
		}
	}
	return true
}

func fixOrdering(rules map[int][]int, pages []int) []int {
	result := make([]int, len(pages))
	_ = copy(result, pages)
	for i := 0; i < len(result); {
		page := result[i]
		switched := false
		if values, ok := rules[page]; ok {
			for _, v := range values {
				j := slices.Index(result[i:], v)
				if j >= 0 {
					fmt.Printf("Page %d Value %d. Swapping %d and %d\n", page, v, result[i], result[i+j])
					result[i], result[i+j] = result[i+j], result[i]
				}
			}
		}
		if !switched {
			i++
		}
	}
	return result
}

func containsAll(array []int, values []int) bool {
	for _, value := range values {
		if !slices.Contains(array, value) {
			return false
		}
	}

	return true
}
