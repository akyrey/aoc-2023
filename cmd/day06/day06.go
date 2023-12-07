package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/akyrey/aoc-2023/internal"
)

const ACCELLERATION = 1

type Result struct {
	HoldTime          int
	DistanceTravelled int
}

type Race struct {
	Results        []Result
	Time           int
	RecordDistance int
}

func (r *Race) getWinResults() {
	for tick := 1; tick < r.Time-1; tick++ {
		speed := tick * ACCELLERATION
		distance := (r.Time - tick) * speed
		if distance > r.RecordDistance {
			r.Results = append(r.Results, Result{HoldTime: tick, DistanceTravelled: (r.Time - tick) * speed})
		}
	}
}

func scanNumbers(line string) []int {
	stringValues := strings.Split(strings.TrimSpace(line), " ")
	values := make([]int, 0)
	for i := range stringValues {
		trimmed := strings.TrimSpace(stringValues[i])
		value, err := strconv.Atoi(trimmed)
		if err == nil {
			values = append(values, value)
		}
	}

	return values
}

func main() {
	f, err := internal.GetFileToReadFrom(6, false)
	internal.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	isTimeLine := true
	times := make([]int, 0)
	distances := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ":")

		if isTimeLine {
			times = scanNumbers(tokens[1])
		} else {
			distances = scanNumbers(tokens[1])
		}

		isTimeLine = false
	}

	races := make([]Race, 0)
	for i := range times {
		race := Race{Time: times[i], RecordDistance: distances[i], Results: make([]Result, 0)}
		race.getWinResults()
		races = append(races, race)
	}

	result := 1
	for _, race := range races {
		result *= len(race.Results)
	}
	fmt.Println(result)
}
