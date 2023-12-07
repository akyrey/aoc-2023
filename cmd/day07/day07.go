package main

import (
	"bufio"

	"github.com/akyrey/aoc-2023/internal"
)

func main() {
	f, err := internal.GetFileToReadFrom(7, true)
	internal.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
	}
}

