package main

import (
	"github.com/akyrey/aoc-2023/internal"
)

func main() {
	f, err := internal.GetFileToReadFrom(1, internal.Test)
	internal.CheckError(err)

	f.Close()
}
