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
	Destination string
	Result      int
}

func (s Step) proceed(mappings []Info) int {
	if s.Destination == "location" {
		return s.Result
	}
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
	start := internal.FindFunc(mappings, func(current Info) bool {
		return current.Source == "seed"
	})
	if start == nil {
		log.Fatal("Couldn't find start")
	}
	end := internal.FindFunc(mappings, func(current Info) bool {
		return current.Destination == "location"
	})
	if end == nil {
		log.Fatal("Couldn't find end")
	}
}

func main() {
	f, err := internal.GetFileToReadFrom(5, true)
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

	fmt.Println(seeds, mappings)
}
