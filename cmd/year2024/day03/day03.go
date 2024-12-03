package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"

	"github.com/akyrey/aoc/internal"
)

func main() {
	f, err := internal.GetFileToReadFrom(3, false)
	internal.CheckError(err)
	defer f.Close()

	total := 0
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	reNumbers := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		multiplies := re.FindAllString(line, -1)
		for _, m := range multiplies {
			numbers := reNumbers.FindAllString(m, 2)
			if len(numbers) != 2 {
				panic("Invalid input")
			}
			one, err := strconv.Atoi(numbers[0])
			internal.CheckError(err)
			two, err := strconv.Atoi(numbers[1])
			internal.CheckError(err)
			total += one * two
		}
	}

	fmt.Printf("%d\n", total)
}
