package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/akyrey/aoc/internal"
)

type Direction int

const (
	NorthSouth Direction = iota
	EastWest
	NorthEast
	NorthWest
	SouthWest
	SouthEast
	Ground
	StartingPosition
)

// * | is a vertical pipe connecting north and south.
// * - is a horizontal pipe connecting east and west.
// * L is a 90-degree bend connecting north and east.
// * J is a 90-degree bend connecting north and west.
// * 7 is a 90-degree bend connecting south and west.
// * F is a 90-degree bend connecting south and east.
// * . is ground; there is no pipe in this tile.
// * S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
var PipeType = map[string]Direction{
	"|": NorthSouth,
	"-": EastWest,
	"L": NorthEast,
	"J": NorthWest,
	"7": SouthWest,
	"F": SouthEast,
	".": Ground,
	"S": StartingPosition,
}

type Point struct {
	X int
	Y int
}

type Path struct {
	Points []Point
	Length int
}

func scanLine(line string) ([]Direction, *Point) {
	tokens := strings.Split(line, "")
	columns := make([]Direction, len(tokens))
	var start *Point
	for i, token := range tokens {
		columns[i] = PipeType[token]
		if token == "S" {
			start = &Point{X: i, Y: 0}
		}
	}

	return columns, start
}

func walk(matrix [][]Direction, point Point, path *Path, seen [][]bool) {
	if point.Y >= len(matrix) || point.X >= len(matrix[0]) {
		return
	}
	if point.Y < 0 || point.X < 0 {
		return
	}
	if seen[point.Y][point.X] {
		return
	}
	if matrix[point.Y][point.X] == Ground {
		return
	}

	current := matrix[point.Y][point.X]
	nextPoint := Point{X: point.X, Y: point.Y}
	path.Points = append(path.Points, point)
	path.Length++
	seen[point.Y][point.X] = true
	if current == StartingPosition {
		if point.Y-1 >= 0 && internal.ContainsFunc([]Direction{NorthSouth, SouthEast, SouthWest}, func(c Direction) bool {
			return c == matrix[point.Y-1][point.X]
		}) {
			nextPoint.Y = point.Y - 1
		} else if point.Y+1 < len(matrix) && internal.ContainsFunc([]Direction{NorthSouth, NorthWest, NorthEast}, func(c Direction) bool {
			return c == matrix[point.Y+1][point.X]
		}) {
			nextPoint.Y = point.Y + 1
		} else if point.X-1 >= 0 && internal.ContainsFunc([]Direction{EastWest, NorthEast, SouthEast}, func(c Direction) bool {
			return c == matrix[point.Y][point.X-1]
		}) {
			nextPoint.X = point.X - 1
		} else if point.X+1 < len(matrix[0]) && internal.ContainsFunc([]Direction{EastWest, NorthWest, SouthWest}, func(c Direction) bool {
			return c == matrix[point.Y][point.X+1]
		}) {
			nextPoint.X = point.X + 1
		}
	} else {
		prevPoint := path.Points[len(path.Points)-2]
		switch current {
		case NorthSouth:
			if prevPoint.Y == point.Y-1 {
				nextPoint.Y = point.Y + 1
			} else {
				nextPoint.Y = point.Y - 1
			}
		case EastWest:
			if prevPoint.X == point.X-1 {
				nextPoint.X = point.X + 1
			} else {
				nextPoint.X = point.X - 1
			}
		case NorthEast:
			if prevPoint.X == point.X+1 {
				nextPoint.Y = point.Y - 1
			} else {
				nextPoint.X = point.X + 1
			}
		case NorthWest:
			if prevPoint.Y == point.Y-1 {
				nextPoint.X = point.X - 1
			} else {
				nextPoint.Y = point.Y - 1
			}
		case SouthWest:
			if prevPoint.X == point.X-1 {
				nextPoint.Y = point.Y + 1
			} else {
				nextPoint.X = point.X - 1
			}
		case SouthEast:
			if prevPoint.Y == point.Y+1 {
				nextPoint.X = point.X + 1
			} else {
				nextPoint.Y = point.Y + 1
			}
		}
	}

	walk(matrix, nextPoint, path, seen)
}

func main() {
	f, err := internal.GetFileToReadFrom(10, false)
	internal.CheckError(err)
	defer f.Close()

	matrix := make([][]Direction, 0)
	scanner := bufio.NewScanner(f)
	var start *Point
	for scanner.Scan() {
		line := scanner.Text()
		columns, s := scanLine(line)
		if s != nil {
			s.Y = len(matrix)
			start = s
		}
		matrix = append(matrix, columns)
	}

	path := &Path{Points: make([]Point, 0), Length: 0}
	seen := make([][]bool, len(matrix))
	for i := range seen {
		seen[i] = make([]bool, len(matrix[0]))
	}
	walk(matrix, *start, path, seen)
	fmt.Println(path.Length / 2)
}
