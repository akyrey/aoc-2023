package internal

import (
	"fmt"
	"os"
	"sort"
)

const (
	Test bool   = true
	Year string = "2024"
)

func Contains[T int | string](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Find[T any](s []T, check func(current T) bool) *T {
	for _, a := range s {
		if check(a) {
			return &a
		}
	}
	return nil
}

func ContainsFunc[T any](s []T, check func(current T) bool) bool {
	for _, a := range s {
		if check(a) {
			return true
		}
	}
	return false
}

func Filter[T any](s []T, check func(current T) bool) []T {
	filtered := make([]T, 0)
	for _, a := range s {
		if check(a) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

func Every[T any](s []T, condition func(current T) bool) bool {
	for _, a := range s {
		if !condition(a) {
			return false
		}
	}
	return true
}

func Some[T any](s []T, condition func(current T) bool) bool {
	for _, a := range s {
		if condition(a) {
			return true
		}
	}
	return false
}

func MinMax(array []int) (int, int) {
	max := array[0]
	min := array[0]

	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func GetFileToReadFrom(day int, test bool) (*os.File, error) {
	dayStr := getDayString(day)
	if test {
		return os.Open(fmt.Sprintf("cmd/year%s/day%s/test%s.txt", Year, dayStr, dayStr))
	}
	return os.Open(fmt.Sprintf("cmd/year%s/day%s/input%s.txt", Year, dayStr, dayStr))
}

func getDayString(day int) string {
	if day < 10 {
		return fmt.Sprintf("0%d", day)
	}

	return fmt.Sprintf("%d", day)
}

func StringPtrToString(p *string) string {
	if p != nil {
		return *p
	}

	return "(nil)"
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
