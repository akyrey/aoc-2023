package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/akyrey/aoc/internal"
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
	// fmt.Println(node)

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

	// fmt.Println(lines)

	nodesToCheck := make([]string, 0)
	for key := range lines {
		if strings.HasSuffix(key, "A") {
			nodesToCheck = append(nodesToCheck, key)
		}
	}
	if len(nodesToCheck) == 0 {
		log.Fatalf("Could not find node **A")
	}

	// fmt.Println(nodesToCheck)
	i := 0
	paths := make(map[string][]string, 0)
	for index, node := range nodesToCheck {
		paths[node] = make([]string, 0)
		for !strings.HasSuffix(node, "Z") {
			paths[node] = append(paths[node], node)
			move := instructions[i%len(instructions)]
			switch move {
			case "L":
				nodesToCheck[index] = lines[node].Left
			case "R":
				nodesToCheck[index] = lines[node].Right
			}
		}
		paths[node] = append(paths[node], node)
		fmt.Println(paths[node])
		i += 1
	}

	fmt.Println(paths)
}
