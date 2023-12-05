package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/akyrey/aoc-2023/internal"
)

type Remapping struct {
	DestRangeStart   int
	SourceRangeStart int
	RangeLength      int
}

type Info struct {
	Source      string
	Destination string
	Remappings  []Remapping
}

type Step struct {
	Source string
	Value  int
}

func (s Step) proceed(mappings []Info) int {
	if s.Source == "location" {
		return s.Value
	}

	next := internal.FindFunc(mappings, func(c Info) bool {
		return c.Source == s.Source
	})

	fmt.Println(s)
	if next == nil {
		log.Fatal("cannot proceed to next step")
	}

	transValue := internal.FindFunc(next.Remappings, func(c Remapping) bool {
		return s.Value >= c.SourceRangeStart && s.Value <= c.SourceRangeStart+c.RangeLength
	})

	if transValue != nil {
		newStep := Step{Source: next.Destination, Value: transValue.DestRangeStart + s.Value - transValue.SourceRangeStart}
		return newStep.proceed(mappings)
	}

	newStep := Step{Source: next.Destination, Value: s.Value}
	return newStep.proceed(mappings)
}

func scanNumbers(line string) []int {
	stringValues := strings.Split(line, " ")
	values := make([]int, 0)
	for i := range stringValues {
		value, err := strconv.Atoi(stringValues[i])
		internal.CheckError(err)
		values = append(values, value)
	}

	return values
}

func scanSeeds(line string) []int {
	nums := strings.Split(line, ":")
	if len(nums) != 2 {
		log.Fatal("couldn't split on seeds row")
	}

	tokens := strings.TrimSpace(nums[1])
	return scanNumbers(tokens)
}

func scanMapLine(line string) *Info {
	seps := strings.Split(line, " ")
	if len(seps) != 2 {
		log.Fatalf("Wrong map format %s", seps)
	}
	tokens := strings.Split(seps[0], "-to-")
	if len(tokens) != 2 {
		log.Fatalf("Wrong map format %s", tokens)
	}

	return &Info{Source: tokens[0], Destination: tokens[1], Remappings: make([]Remapping, 0)}
}

func calcLocation(seed int, mappings []Info) int {
	step := Step{Source: "seed", Value: seed}

	return step.proceed(mappings)
}

func main() {
	f, err := internal.GetFileToReadFrom(5, false)
	internal.CheckError(err)
	defer f.Close()

	i := 0
	seeds := make([]int, 0)
	mappings := make([]Info, 0)
	var currentInfo *Info
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Scan seeds
		if i == 0 {
			seeds = scanSeeds(line)
		} else if len(line) == 0 {
			// Skip empty lines
			if currentInfo != nil {
				mappings = append(mappings, *currentInfo)
			}
			continue
		} else {
			tokens := strings.Split(line, " ")
			if len(tokens) < 1 {
				log.Fatalf("Failed scanning line %d", i)
			}
			_, err := strconv.Atoi(tokens[0])
			if err != nil {
				// Map line
				currentInfo = scanMapLine(line)
			} else {
				// Number line
				if currentInfo == nil {
					log.Fatalf("Missing info for line %v", line)
				}
				values := scanNumbers(line)
				if len(values) != 3 {
					log.Fatalf("Wrong format for remapping line %v", values)
				}
				remapping := Remapping{
					DestRangeStart:   values[0],
					SourceRangeStart: values[1],
					RangeLength:      values[2],
				}
				currentInfo.Remappings = append(currentInfo.Remappings, remapping)
			}
		}
		i += 1
	}

	locations := make([]int, 0)
	for _, seed := range seeds {
		locations = append(locations, calcLocation(seed, mappings))
	}

	min := locations[0]
	for _, value := range locations {
		if value < min {
			min = value
		}
	}
	fmt.Println(min)
}
