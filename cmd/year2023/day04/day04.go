package main

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/akyrey/aoc/internal"
)

type Card struct {
	WinningNumbers []int
	FoundNumbers   []int
	ID             int
}

func parseCardID(s string) (int, error) {
	numberRegex, err := regexp.Compile(`\d+`)
	if err != nil {
		return -1, err
	}
	matches := numberRegex.FindAllString(s, 1)
	if len(matches) != 1 {
		return -1, errors.New("wrongly formatted string")
	}

	id, err := strconv.Atoi(matches[0])
	if err != nil {
		return -1, err
	}

	return id, nil
}

func parseNumbers(s string) ([]int, error) {
	trimmed := strings.Trim(s, " ")
	tokens := strings.Split(trimmed, " ")

	numbers := make([]int, 0)
	for _, valueString := range tokens {
		value, err := strconv.Atoi(valueString)
		if err == nil {
			numbers = append(numbers, value)
		}
	}

	return numbers, nil
}

func parseCard(id int, s string) (*Card, error) {
	numbers := strings.Split(s, " | ")
	if len(numbers) != 2 {
		return nil, errors.New("wrongly formatted string, couldn't differentiate winning from extracted numbers")
	}

	winning, err := parseNumbers(numbers[0])
	if err != nil {
		return nil, err
	}
	extracted, err := parseNumbers(numbers[1])
	if err != nil {
		return nil, err
	}

	return &Card{ID: id, WinningNumbers: winning, FoundNumbers: extracted}, nil
}

func parseLine(line string) (*Card, error) {
	tokens := strings.Split(line, ":")
	if len(tokens) != 2 {
		return nil, errors.New("wrongly formatted string, couldn't find card and numbers")
	}

	id, err := parseCardID(tokens[0])
	if err != nil {
		return nil, err
	}

	card, err := parseCard(id, tokens[1])
	if err != nil {
		return nil, err
	}

	return card, nil
}

func calcCardPoints(cards []Card) []int {
	res := make([]int, 0)

	for _, card := range cards {
		wins := make([]int, 0)
		for _, number := range card.FoundNumbers {
			if internal.Contains(card.WinningNumbers, number) {
				wins = append(wins, number)
			}
		}

		points := 0
		for i := range wins {
			if i == 0 {
				points = 1
			} else {
				points *= 2
			}
		}

		res = append(res, points)
	}

	return res
}

func calcTotalCards(cards []Card) int {
	count := 0
	mapOfReplays := make(map[int]int, 0)

	for i := 0; i < len(cards); {
		count += 1
		card := cards[i]
		wins := make([]int, 0)
		for _, number := range card.FoundNumbers {
			if internal.Contains(card.WinningNumbers, number) {
				wins = append(wins, number)
			}
		}

		for i := range wins {
			id := card.ID + i
			if _, ok := mapOfReplays[id]; !ok {
				mapOfReplays[id] = 0
			}
			mapOfReplays[id] += 1
		}

		if value, ok := mapOfReplays[i]; ok && value > 0 {
			mapOfReplays[i] -= 1
			// fmt.Printf("Processing card %d, decrementing current counter to %d, %v\n", i+1, mapOfReplays[i], mapOfReplays)
		} else {
			i += 1
			// fmt.Printf("Finished processing card %d, current %v\n", i, mapOfReplays)
		}
	}

	return count
}

func main() {
	f, err := internal.GetFileToReadFrom(4, false)
	internal.CheckError(err)
	defer f.Close()

	cards := make([]Card, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		card, err := parseLine(scanner.Text())
		internal.CheckError(err)
		cards = append(cards, *card)
	}

	cardPoints := calcCardPoints(cards)

	sum := 0
	for _, point := range cardPoints {
		sum += point
	}

	// fmt.Println(sum)

	totalWinnindCards := calcTotalCards(cards)

	fmt.Println(totalWinnindCards)
}
