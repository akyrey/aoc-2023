package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/akyrey/aoc/internal"
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

func scanNumbers(line string) int {
	stringValues := strings.Split(strings.TrimSpace(line), " ")
	values := ""
	for i := range stringValues {
		values += strings.TrimSpace(stringValues[i])
	}

	value, err := strconv.Atoi(values)
	internal.CheckError(err)

	return value
}

func main() {
	f, err := internal.GetFileToReadFrom(6, false)
	internal.CheckError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	isTimeLine := true
	race := Race{Results: make([]Result, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ":")

		if isTimeLine {
			race.Time = scanNumbers(tokens[1])
		} else {
			race.RecordDistance = scanNumbers(tokens[1])
		}

		isTimeLine = false
	}
	race.getWinResults()

	result := 1
	result *= len(race.Results)
	fmt.Println(result)
}
