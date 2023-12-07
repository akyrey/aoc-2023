package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"
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
	Seed   Seed
}

func (s Step) proceed(mappings []Info) int {
	if s.Source == "location" {
		return s.Seed.Start
	}

	next := internal.Find(mappings, func(c Info) bool {
		return c.Source == s.Source
	})

	fmt.Println(s)
	if next == nil {
		log.Fatal("cannot proceed to next step")
	}

	overlappings := getOverlappints(s.Seed.Start, s.Seed.End, next.Remappings)
	transValue := internal.Find(next.Remappings, func(c Remapping) bool {
		return s.Value >= c.SourceRangeStart && s.Value <= c.SourceRangeStart+c.RangeLength
	})

	if transValue != nil {
		newStep := Step{Source: next.Destination, Value: transValue.DestRangeStart + s.Value - transValue.SourceRangeStart}
		return newStep.proceed(mappings)
	}

	newStep := Step{Source: next.Destination, Seed: s.Seed}
	return newStep.proceed(mappings)
}

type Seed struct {
	Start int
	End   int
}

func getOverlappints(start, end int, remappings []Remapping) []Seed {
    // Overlappings are sorted by source range start
	overlappings := internal.Filter(remappings, func(current Remapping) bool {
		endRemapping := current.SourceRangeStart + current.RangeLength
		return start < endRemapping && current.SourceRangeStart < end
	})

	if len(overlappings) == 0 {
		return []Seed{{Start: start, End: end}}
	}

	seeds := make([]Seed, 0)
	segmentStart := start
	for i := 0; i < len(overlappings); i++ {
		seedStart := segmentStart
		// The start is inside an overlapping
		if segmentStart >= overlappings[i].SourceRangeStart && segmentStart < overlappings[i].SourceRangeStart+overlappings[i].RangeLength {
			seedStart = overlappings[i].DestRangeStart + segmentStart - overlappings[i].SourceRangeStart
		}
        seedEnd := end
        if i < len(overlappings) - 1 {
            seedEnd = overlappings[i+1].SourceRangeStart
        }
        if overlappings[i].SourceRangeStart + overlappings[i].RangeLength < end {
            seedEnd = overlappings[i].SourceRangeStart + overlappings[i].RangeLength
        }




		startNextMapping := end
		if i < len(overlappings)-1 {
			startNextMapping = overlappings[i+1].SourceRangeStart
		}
		endRemapping := overlappings[i].SourceRangeStart + overlappings[i].RangeLength
		minEnd, _ := internal.MinMax([]int{end, endRemapping, startNextMapping})
		seed := Seed{Start: seedStart, End: minEnd}
		segmentStart = minEnd + 1
		seeds = append(seeds, seed)
	}

	return seeds
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

func calcLocation(seed Seed, mappings []Info) int {
	step := Step{Source: "seed", Seed: seed}

	return step.proceed(mappings)
}

func main() {
	f, err := internal.GetFileToReadFrom(5, true)
	internal.CheckError(err)
	defer f.Close()

	i := 0
	seedsRange := make([]int, 0)
	mappings := make([]Info, 0)
	var currentInfo *Info
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Scan seeds
		if i == 0 {
			seedsRange = scanSeeds(line)
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

	seeds := make([]Seed, 0)
	for i := 0; i < len(seedsRange)-1; i += 2 {
		seed := Seed{Start: seedsRange[i], End: seedsRange[i] + seedsRange[i+1]}
		seeds = append(seeds, seed)
	}

	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i].Start < seeds[j].Start
	})

	for _, mapping := range mappings {
		sort.Slice(mapping.Remappings, func(i, j int) bool {
			return mapping.Remappings[i].SourceRangeStart < mapping.Remappings[j].SourceRangeStart
		})
	}

	// locations := make([]int, 0)
	// for _, seed := range seeds {
	// 	locations = append(locations, calcLocation(seed, mappings))
	// }
	//
	// min := locations[0]
	// for _, value := range locations {
	// 	if value < min {
	// 		min = value
	// 	}
	// }
	// fmt.Println(min)
}
