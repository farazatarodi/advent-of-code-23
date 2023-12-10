package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Puzzle2(lines []string) {
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
		last := diff2(history)
		sum += last
	}

	fmt.Println("Day 9, Puzzle 2:", sum)
}

func diff2(numbers []int) int {
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

	diffNumber := diff2(diffs)
	firstNumber := numbers[0] - diffNumber

	return firstNumber

}
