package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	"github.com/akyrey/aoc/internal"
)

type Action int

const (
	OBSTABLE            = "#"
	FACING_UP           = "^"
	FACING_DOWN         = "v"
	FACING_LEFT         = "<"
	FACING_RIGHT        = ">"
	END          Action = 0
	MOVE         Action = 1
	TURN         Action = 2
)

type Guard struct {
	Facing string
	Row    int
	Col    int
}

func main() {
	f, err := internal.GetFileToReadFrom(6, false)
	internal.CheckError(err)
	defer f.Close()

	visited := make(map[string]bool, 0)
	matrix := make([][]string, 0)
	guard := Guard{Row: -1, Col: -1, Facing: ""}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)
		if idx := facingPosition(chars); idx != -1 {
			guard.Row = len(matrix) - 1
			guard.Col = idx
			guard.Facing = chars[idx]
		}
	}

	rows := len(matrix)
	cols := len(matrix[0])

	for !guard.isOutOfTheBoard(rows, cols) {
		switch guard.peek(matrix) {
		case MOVE:
			pos := fmt.Sprintf("%d-%d", guard.Row, guard.Col)
			if _, ok := visited[pos]; !ok {
				visited[pos] = true
			}
			fmt.Printf("Guard: %v\n", guard)
			guard.move()
		case TURN:
			guard.turn()
			fmt.Printf("Turning guard: %v\n", guard)
		case END:
			pos := fmt.Sprintf("%d-%d", guard.Row, guard.Col)
			if _, ok := visited[pos]; !ok {
				visited[pos] = true
			}
			guard.Row = -1
			fmt.Printf("Finished\n")
		}
	}

	for row, cols := range matrix {
		for col := range cols {
			if matrix[row][col] == OBSTABLE {
				fmt.Printf("#")
			} else if _, ok := visited[fmt.Sprintf("%d-%d", row, col)]; ok {
				fmt.Printf("X")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("%d\n", len(visited))
}

func facingPosition(chars []string) int {
	return slices.IndexFunc(chars, func(s string) bool {
		return s == FACING_UP || s == FACING_DOWN || s == FACING_LEFT || s == FACING_RIGHT
	})
}

func (g *Guard) move() {
	g.Row, g.Col = g.nextPos()
}

func (g *Guard) isOutOfTheBoard(rows int, cols int) bool {
	return g.Row < 0 || g.Row >= rows || g.Col < 0 || g.Col >= cols
}

func (g *Guard) peek(matrix [][]string) Action {
	rows := len(matrix)
	cols := len(matrix[0])
	row, col := g.nextPos()

	if row < 0 || row >= rows || col < 0 || col >= cols {
		return END
	}

	if matrix[row][col] == OBSTABLE {
		return TURN
	}

	return MOVE
}

func (g *Guard) turn() {
	directions := []string{FACING_UP, FACING_RIGHT, FACING_DOWN, FACING_LEFT}
	idx := (slices.Index(directions, g.Facing) + 1) % len(directions)
	g.Facing = directions[idx]
}

func (g *Guard) nextPos() (int, int) {
	switch g.Facing {
	case FACING_UP:
		return g.Row - 1, g.Col
	case FACING_DOWN:
		return g.Row + 1, g.Col
	case FACING_LEFT:
		return g.Row, g.Col - 1
	case FACING_RIGHT:
		return g.Row, g.Col + 1
	default:
		panic("Invalid guard text")
	}
}
