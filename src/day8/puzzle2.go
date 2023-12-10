package main

import (
	"fmt"
	"strings"
)

func Puzzle2(lines []string) {
	instructions := strings.Split("LLRRRLRLLRLRRLRLRLRRRLLRRLRRRLRRRLRRRLRRRLRRRLRRLRLLRRRLRRLLRLRLLLRRLRRLRLRLRLRRRLRLRRRLRRLLLRRRLLRRLLRRLLRRRLLLLRLRLRRRLRLRRRLRLLLRLRRLRRRLRRRLRRRLRRRLLRRLLLLRRLLRRLLRRLRLRRRLRRRLRRRLRRLRRRLRRLRRLRRLRLRRRLRRLRRRLRRRLRRLRLRRRLRRLLRLRRLRRRLRLRRLRRRLRRLRRLRRRLLRRRR", "")

	nodes := make(map[string][]string)
	for _, line := range lines {
		chunks := strings.Split(line, "=")

		sourceNode := strings.TrimSpace(chunks[0])

		chunks = strings.Split(chunks[1], ", ")

		leftNode := chunks[0][2:]
		rightNode := chunks[1][:len(chunks[1])-1]

		nodes[sourceNode] = []string{leftNode, rightNode}
	}

	var sources []string
	for source := range nodes {
		if strings.HasSuffix(source, "A") {
			sources = append(sources, source)
		}
	}

	counter := 0
	for counter > -1 {
		var nextNodes []string
		for _, source := range sources {
			nextNodes = append(nextNodes, follow2(nodes, instructions, source, counter))
		}

		finished := true
		for _, dest := range nextNodes {
			if !strings.HasSuffix(dest, "Z") {
				finished = false
			}
		}

		if counter%1000000 == 0 {
			fmt.Println(counter)
		}

		if finished {
			break
		}

		counter++
		sources = nextNodes
	}

	fmt.Println("Day 8, Puzzle 2:", counter)
}

func follow2(nodes map[string][]string, instructions []string, source string, counter int) string {
	instructionIndex := counter % len(instructions)
	instruction := 0
	if instructions[instructionIndex] == "R" {
		instruction = 1
	}

	nextNode := nodes[source][instruction]

	return nextNode
}
