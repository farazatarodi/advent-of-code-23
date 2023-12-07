package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Puzzle2(lines []string) {
	cards := make([]int, len(lines))

	for lineIndex, line := range lines {
		lineChunks := strings.Split(line, ":")
		numberLists := strings.Split(lineChunks[1], "|")

		winningNumbers := strings.Split(strings.TrimSpace(numberLists[0]), " ")
		myNumbers := strings.Split(strings.TrimSpace(numberLists[1]), " ")

		cards[lineIndex]++

		var matches []int
		for _, number := range myNumbers {
			if slices.Contains(winningNumbers, number) {
				numberInt, err := strconv.Atoi(number)
				if err == nil {
					matches = append(matches, numberInt)
				}
			}
		}

		for i := range matches {
			cards[lineIndex+i+1] += cards[lineIndex]
		}

	}

	sum := 0
	for _, amount := range cards {
		sum += amount
	}

	fmt.Println("Day 4, Puzzle 2:", sum)
}
