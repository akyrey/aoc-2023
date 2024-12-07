package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/akyrey/aoc/internal"
)

type Operator string

const (
	ADD           Operator = "+"
	MULTIPLY      Operator = "*"
	CONCATENATION Operator = "||"
)

type Operation struct {
	Operators []Operator
	Values    []int
	Result    int
}

func main() {
	f, err := internal.GetFileToReadFrom(7, false)
	internal.CheckError(err)
	defer f.Close()

	total := 0
	operations := make([]Operation, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		if len(split) != 2 {
			panic("Invalid input")
		}

		operation := Operation{
			Values:    make([]int, 0),
			Operators: make([]Operator, 0),
		}
		operation.Result, err = strconv.Atoi(split[0])
		internal.CheckError(err)

		values := strings.Split(split[1], " ")
		for _, v := range values {
			value, err := strconv.Atoi(v)
			internal.CheckError(err)
			operation.Values = append(operation.Values, value)
		}

		operations = append(operations, operation)
	}

	for _, o := range operations {
		for _, c := range o.findCombinations() {
			o.Operators = c
			if o.isValid() {
				total += o.calculate()
				break
			}
		}
	}

	fmt.Printf("Total: %d\n", total)
}

func (o *Operation) findCombinations() [][]Operator {
	operations := []Operator{ADD, MULTIPLY, CONCATENATION}
	count := len(o.Values) - 1
	combinations := make([][]Operator, count)
	for i := 0; i < count; i++ {
		for _, op := range operations {
			combinations[i] = append(combinations[i], op)
		}
	}

	return product(1, combinations...)
}

func (o *Operation) isValid() bool {
	res := o.calculate()
	// fmt.Printf("Calculating: %v with result: %d\n", o, res)
	return o.Result == res
}

func (o *Operation) calculate() int {
	result := 0
	first := true
	var value int
	values := make([]int, len(o.Values))
	var operator Operator
	operators := make([]Operator, len(o.Operators))
	copy(values, o.Values)
	copy(operators, o.Operators)
	var err error
	for len(values) > 0 {
		value, values = values[0], values[1:]
		if first {
			result = value
			first = false
		} else {
			operator, operators = operators[0], operators[1:]
			switch operator {
			case ADD:
				result += value
			case MULTIPLY:
				result *= value
			case CONCATENATION:
				concatenation := fmt.Sprintf("%d%d", result, value)
				result, err = strconv.Atoi(concatenation)
				internal.CheckError(err)
			}
		}
	}

	return result
}

// product([1,2,3],[4,5,6],2) = > [[1,2,3],[4,5,6],[1,2,3],[4,5,6]]
func product(n int, input ...[]Operator) [][]Operator {
	// append all input array into pools
	// account repeat value (n) so it repeats n times
	pools := make([][]Operator, 0)
	for i := 1; i <= n; i++ {
		for _, x := range input {
			pools = append(pools, [][]Operator{x}...)
		}
	}
	prod := make([][]Operator, 1)
	// go over each and every pool
	for _, pool := range pools {
		next := make([][]Operator, 0)
		for _, x := range prod {
			for _, y := range pool {
				// x = [1]
				// y = 2
				// t = [1,2]
				t := make([]Operator, 0)
				t = append(t, x...)
				t = append(t, y)
				next = append(next, [][]Operator{t}...)
			}
		}
		// make prod = intermediate array next
		prod = next
	}
	return prod
}
