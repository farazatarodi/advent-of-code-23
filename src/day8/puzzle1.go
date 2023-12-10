package main

import (
	"fmt"
	"strings"
)

func Puzzle1(lines []string) {
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

	steps := follow(nodes, instructions, "AAA", 0)

	fmt.Println("Day 8, Puzzle 1:", steps)
}

func follow(nodes map[string][]string, instructions []string, source string, counter int) int {
	if source == "ZZZ" {
		return counter
	}

	instructionIndex := counter % len(instructions)
	instruction := 0
	if instructions[instructionIndex] == "R" {
		instruction = 1
	}

	nextNode := nodes[source][instruction]
	result := follow(nodes, instructions, nextNode, counter+1)

	return result
}
