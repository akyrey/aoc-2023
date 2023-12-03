package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/akyrey/aoc-2023/internal"
)

type SymbolKind int64

const (
	Number SymbolKind = iota
	Symbol
)

type Token struct {
	Char  string
	Line  int
	Start int
	End   int
}

type Map map[SymbolKind][]Token

func getMatches(line, expr string, index int) []Token {
	numberRegex, err := regexp.Compile(expr)
	internal.CheckError(err)

	symbols := make([]Token, 0)
	matches := numberRegex.FindAllString(line, -1)
	positions := numberRegex.FindAllStringIndex(line, -1)

	if matches != nil && positions != nil && len(matches) == len(positions) {
		for i := range matches {
			symbol := Token{Char: matches[i], Line: index, Start: positions[i][0], End: positions[i][1]}
			symbols = append(symbols, symbol)
		}
	}

	return symbols
}

func scanLine(line string, index int) Map {
	symbolsMap := make(Map, 0)
	symbolsMap[Number] = getMatches(line, `\d+`, index)
	symbolsMap[Symbol] = getMatches(line, `\*`, index)

	return symbolsMap
}

func areTokenOnAdjacentLines(a, b Token) bool {
	return a.Line == b.Line-1 || a.Line == b.Line+1 || a.Line == b.Line
}

func areTokensTouching(a, b Token) bool {
	return areTokenOnAdjacentLines(a, b) && a.Start >= b.Start-1 && a.End <= b.End+1
}

func findNumbersAdjacentsToSymbols(symbolsMap Map) []int {
	valid := make([]int, 0)

	for _, symbol := range symbolsMap[Symbol] {
		numbers := make([]int, 0)
		for _, number := range symbolsMap[Number] {
			if areTokensTouching(symbol, number) {
				value, err := strconv.Atoi(number.Char)
				internal.CheckError(err)
				numbers = append(numbers, value)
			}
		}
		if len(numbers) == 2 {
			gearRatio := 1
			for _, n := range numbers {
				gearRatio *= n
			}
			valid = append(valid, gearRatio)
		}
	}

	return valid
}

func scanSymbols(f *os.File) Map {
	list := make(Map, 0)
	i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		res := scanLine(line, i)
		list[Number] = append(list[Number], res[Number]...)
		list[Symbol] = append(list[Symbol], res[Symbol]...)
		i += 1
	}

	return list
}

func main() {
	f, err := internal.GetFileToReadFrom(3, false)
	internal.CheckError(err)
	defer f.Close()

	list := scanSymbols(f)
	valid := findNumbersAdjacentsToSymbols(list)

	sum := 0
	for _, value := range valid {
		sum += value
	}

	fmt.Print(sum)
}
