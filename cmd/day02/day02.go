package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/akyrey/aoc-2023/internal"
)

const (
	RED   = 12
	BLUE  = 14
	GREEN = 13
)

type Game struct {
	Blue  []int
	Red   []int
	Green []int
	ID    int
}

func parseGameID(s string) (int, error) {
	idArr := strings.Split(s, " ")
	if len(idArr) != 2 {
		return -1, errors.New("wrongly formatted string")
	}

	id, err := strconv.Atoi(idArr[1])
	if err != nil {
		return -1, err
	}

	return id, nil
}

func parseGameManches(id int, s string) (*Game, error) {
	tokens := strings.Split(s, ";")

	red := make([]int, 0)
	blue := make([]int, 0)
	green := make([]int, 0)

	for _, token := range tokens {
		set := strings.Split(token, ",")
		for _, extraction := range set {
			valueAndColor := strings.Split(strings.Trim(extraction, " "), " ")
			if len(valueAndColor) != 2 {
				return nil, errors.New("invalid input")
			}

			valueStr := strings.Trim(valueAndColor[0], " ")
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				return nil, err
			}
			color := strings.Trim(valueAndColor[1], " ")

			switch color {
			case "green":
				{
					green = append(green, value)
				}
			case "red":
				{
					red = append(red, value)
				}
			case "blue":
				{
					blue = append(blue, value)
				}
			}
		}
	}

	return &Game{ID: id, Red: red, Blue: blue, Green: green}, nil
}

func parseLine(line string) (*Game, error) {
	tokens := strings.Split(line, ":")
	if len(tokens) != 2 {
		return nil, errors.New("wrongly formatted string. cannot split on :")
	}

	id, err := parseGameID(tokens[0])
	if err != nil {
		return nil, err
	}

	game, err := parseGameManches(id, tokens[1])
	if err != nil {
		return nil, err
	}

	return game, nil
}

func findPossibleGamesIDs(games []Game) []int {
	possibleIDs := make([]int, 0)

	for _, game := range games {
		if !internal.ContainsFunc(game.Red, func(current int) bool {
			return current > RED
		}) && !internal.ContainsFunc(game.Blue, func(current int) bool {
			return current > BLUE
		}) && !internal.ContainsFunc(game.Green, func(current int) bool {
			return current > GREEN
		}) {
			possibleIDs = append(possibleIDs, game.ID)
		}
	}

	return possibleIDs
}

func findMinCubesPerGame(games []Game) [][]int {
	res := make([][]int, 0)

	for _, game := range games {
		currentMin := make([]int, 0)

		_, maxRed := internal.MinMax(game.Red)
		_, maxBlue := internal.MinMax(game.Blue)
		_, maxGreen := internal.MinMax(game.Green)

		currentMin = append(currentMin, maxRed)
		currentMin = append(currentMin, maxBlue)
		currentMin = append(currentMin, maxGreen)

		res = append(res, currentMin)
	}

	return res
}

func main() {
	f, err := internal.GetFileToReadFrom(2, false)
	internal.CheckError(err)
	defer f.Close()

	games := make([]Game, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		game, err := parseLine(line)
		internal.CheckError(err)

		games = append(games, *game)
	}

    _ = findPossibleGamesIDs(games)
	minCubes := findMinCubesPerGame(games)
	sum := 0
	for _, game := range minCubes {
		power := 1
		for _, value := range game {
			power *= value
		}
		sum += power
	}

	fmt.Print(sum)
}
