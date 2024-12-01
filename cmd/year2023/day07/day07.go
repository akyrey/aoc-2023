package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/akyrey/aoc/internal"
)

type Card map[string]int

var CARDS_BY_STRENGTH = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards []int
	Type  HandType
	Bid   int
}

type BestCard struct {
	Card  int
	Count int
}

func (h *Hand) calcType() HandType {
	result := make(map[int]int, 0)
	keys := make([]int, 0)
	jollies := 0
	best := BestCard{0, 0}
	for _, card := range h.Cards {
		if card == 1 {
			jollies += 1
		} else {
			if _, ok := result[card]; !ok {
				result[card] = 0
				keys = append(keys, card)
			}
			result[card] += 1
		}
	}

	for _, card := range keys {
		if best.Count < result[card] {
			best.Card = card
			best.Count = result[card]
		} else if best.Card < card && best.Count == result[card] {
			best.Card = card
			best.Count = result[card]
		}
	}

	if _, ok := result[best.Card]; ok {
		result[best.Card] += jollies
	}

	if jollies == 5 {
		return FiveOfAKind
	}

	for i := 0; i < len(keys); i++ {
		if result[keys[i]] == 5 {
			return FiveOfAKind
		}
		if result[keys[i]] == 4 {
			return FourOfAKind
		}
		if result[keys[i]] == 3 && internal.ContainsFunc(keys, func(key int) bool {
			return result[key] == 2
		}) {
			return FullHouse
		}
		if result[keys[i]] == 3 && !internal.ContainsFunc(keys, func(key int) bool {
			return result[key] == 2
		}) {
			return ThreeOfAKind
		}
		if result[keys[i]] == 2 && internal.ContainsFunc(keys, func(key int) bool {
			return key != keys[i] && result[key] == 2
		}) {
			return TwoPairs
		}
		if result[keys[i]] == 2 && !internal.ContainsFunc(keys, func(key int) bool {
			return key != keys[i] && result[key] > 1
		}) {
			return OnePair
		}
	}

	return HighCard
}

func (h *Hand) compare(other Hand) int {
	for i := 0; i < len(h.Cards); i++ {
		if h.Cards[i] != other.Cards[i] {
			return other.Cards[i] - h.Cards[i]
		}
	}

	return 0
}

func scanLine(line string) Hand {
	split := strings.Split(line, " ")
	if len(split) != 2 {
		log.Fatalf("Invalid line: %s", line)
	}

	cardsAsString := strings.Split(split[0], "")
	cards := make([]int, 0)
	for _, cardAsString := range cardsAsString {
		if card, ok := CARDS_BY_STRENGTH[cardAsString]; ok {
			cards = append(cards, card)
		}
	}
	bid, err := strconv.Atoi(split[1])
	internal.CheckError(err)

	return Hand{Cards: cards, Bid: bid}
}

func main() {
	f, err := internal.GetFileToReadFrom(7, false)
	internal.CheckError(err)
	defer f.Close()

	handsByType := make(map[HandType][]Hand, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		hand := scanLine(line)
		hand.Type = hand.calcType()
		if _, ok := handsByType[hand.Type]; !ok {
			handsByType[hand.Type] = make([]Hand, 0)
		}
		handsByType[hand.Type] = append(handsByType[hand.Type], hand)
	}

	rank := 1
	totalWinning := 0
	for handType := HighCard; handType <= FiveOfAKind; handType++ {
		if hands, ok := handsByType[handType]; ok {
			sort.Slice(hands, func(i, j int) bool {
				a := hands[i]
				b := hands[j]
				return a.compare(b) > 0
			})
			for _, hand := range hands {
				totalWinning += hand.Bid * rank
				rank += 1
			}
		}
	}
	fmt.Println(totalWinning)
}
