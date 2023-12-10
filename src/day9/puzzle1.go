package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Puzzle1(lines []string) {
	var histories [][]int
	for _, line := range lines {
		var numbers []int
		for _, numString := range strings.Split(line, " ") {
			number, _ := strconv.Atoi(numString)
			numbers = append(numbers, number)
		}
		histories = append(histories, numbers)
	}

	sum := 0
	for _, history := range histories {
		last := diff(history)
		sum += last
	}

	fmt.Println("Day 9, Puzzle 1:", sum)
}

func diff(numbers []int) int {
	finished := true
	for _, number := range numbers {
		if number != 0 {
			finished = false
		}
	}

	if finished {
		return 0
	}

	var diffs []int
	for i, number := range numbers {
		if i < len(numbers)-1 {
			diffs = append(diffs, numbers[i+1]-number)
		}
	}

	diffNumber := diff(diffs)
	lastNumber := diffNumber + numbers[len(numbers)-1]

	return lastNumber

}
