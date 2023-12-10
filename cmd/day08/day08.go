package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/akyrey/aoc-2023/internal"
)

type Node struct {
	Value string
	Left  string
	Right string
}

func scanNodeLine(line string) Node {
	node := Node{}
	split := strings.Split(line, " = ")
	node.Value = strings.TrimSpace(split[0])
	tokens := strings.Split(split[1], ", ")
	node.Left = strings.Trim(tokens[0], "(")
	node.Right = strings.Trim(tokens[1], ")")
	fmt.Println(node)

	return node
}

func main() {
	f, err := internal.GetFileToReadFrom(8, false)
	// dayStr := "08"
	// f, err := os.Open(fmt.Sprintf("cmd/day%s/test%s-2.txt", dayStr, dayStr))

	internal.CheckError(err)
	defer f.Close()

	lines := make(map[string]Node, 0)
	scanner := bufio.NewScanner(f)
	instructions := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(instructions) == 0 {
			instructions = strings.Split(line, "")
		} else if len(line) > 0 {
			node := scanNodeLine(line)
			lines[node.Value] = node
		}
	}

	fmt.Println(lines)

	root, ok := lines["AAA"]
	if !ok {
		log.Fatalf("Could not find node AAA")
	}

	node := root
	i := 0
	for node.Value != "ZZZ" {
		move := instructions[i%len(instructions)]
		switch move {
		case "L":
			node = lines[node.Left]
		case "R":
			node = lines[node.Right]
		}
		i += 1
	}

	fmt.Println(i)
}
