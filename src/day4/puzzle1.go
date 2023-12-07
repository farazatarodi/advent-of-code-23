package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Puzzle1(lines []string) {
	sum := 0.0

	for _, line := range lines {
		lineChunks := strings.Split(line, ":")
		numberLists := strings.Split(lineChunks[1], "|")
		winningNumbers := strings.Split(strings.TrimSpace(numberLists[0]), " ")
		myNumbers := strings.Split(strings.TrimSpace(numberLists[1]), " ")

		var matches []int
		for _, number := range myNumbers {
			if slices.Contains(winningNumbers, number) {
				numberInt, err := strconv.Atoi(number)
				if err == nil {
					matches = append(matches, numberInt)
				}
			}
		}

		if len(matches) > 0 {
			sum = sum + math.Pow(2, float64(len(matches)-1))
		}

	}

	fmt.Println("Day 4, Puzzle 1:", sum)
}
