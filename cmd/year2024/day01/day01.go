package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/akyrey/aoc/internal"
)

func main() {
	f, err := internal.GetFileToReadFrom(1, false)
	internal.CheckError(err)
	defer f.Close()

	similarity := make(map[int]int, 0)
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lists := strings.Split(line, "   ")
		if len(lists) != 2 {
			panic("invalid input")
		}
		internal.CheckError(err)

		number, err := strconv.Atoi(lists[0])
		internal.CheckError(err)
		leftList = append(leftList, number)

		number, err = strconv.Atoi(lists[1])
		internal.CheckError(err)
		rightList = append(rightList, number)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	distance := 0
	for i := 0; i < len(rightList); i++ {
		_, found := similarity[rightList[i]]
		if !found {
			similarity[rightList[i]] = 0
		}
		similarity[rightList[i]]++
	}
	for i := 0; i < len(leftList); i++ {
		similarity, found := similarity[leftList[i]]
		if found {
			distance = distance + leftList[i]*similarity
		}
	}

	fmt.Printf("%d\n", distance)
}
