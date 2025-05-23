package main

import (
	"bufio"

	"github.com/akyrey/aoc/internal"
)

func main() {
	f, err := internal.GetFileToReadFrom(11, true)
	internal.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
	}
}
