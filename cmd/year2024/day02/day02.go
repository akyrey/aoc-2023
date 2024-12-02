package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/akyrey/aoc/internal"
)

const (
	MIN int = 1
	MAX int = 3
)

func main() {
	f, err := internal.GetFileToReadFrom(2, false)
	internal.CheckError(err)
	defer f.Close()

	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		records := strings.Split(line, " ")
		safe := true
		for i, j, decremental, dampned := 0, 1, false, -1; safe && j < len(records); i, j = i+1, j+1 {
			prevIndex := i
			if dampned == i {
				prevIndex = i - 1
			}
			prev, err := strconv.Atoi(records[prevIndex])
			internal.CheckError(err)
			current, err := strconv.Atoi(records[j])
			internal.CheckError(err)
			diff := current - prev
			if prevIndex == 0 {
				decremental = diff < 0
			}
			absDiff := int(math.Abs(float64(diff)))
			if (diff < 0 && !decremental) || (diff > 0 && decremental) || absDiff < MIN || absDiff > MAX {
				if dampned == -1 {
					dampned = i
				} else {
					safe = false
					continue
				}
			}
		}
		if safe {
			count++
		}
	}

	fmt.Printf("%d\n", count)
}
