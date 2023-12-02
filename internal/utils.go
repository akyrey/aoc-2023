package internal

import (
	"fmt"
	"os"
	"sort"
)

const Test bool = true

func Contains[T int | string](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsFunc[T int | string](s []T, check func(current T) bool) bool {
	for _, a := range s {
		if check(a) {
			return true
		}
	}
	return false
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func GetFileToReadFrom(day int, test bool) (*os.File, error) {
	dayStr := getDayString(day)
	if test {
		return os.Open(fmt.Sprintf("cmd/day%s/test%s.txt", dayStr, dayStr))
	}
	return os.Open(fmt.Sprintf("cmd/day%s/input%s.txt", dayStr, dayStr))
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
