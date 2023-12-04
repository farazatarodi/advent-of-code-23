package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Puzzle2(lines []string) {
	var calibrationValues []int
	for _, line := range lines {
		integers := make(map[int]int)
		firstIntegerIndex := 999
		lastIntegerIndex := 0
		for i, characterByte := range line {
			character := string(characterByte)

			integer, integerError := strconv.Atoi(character)

			if integerError == nil {
				integers[i] = integer

				if i < firstIntegerIndex {
					firstIntegerIndex = i
				}

				if i >= lastIntegerIndex {

					lastIntegerIndex = i
				}
			}

			for k, v := range findNumberStrings(line) {
				integers[k] = v

				if k < firstIntegerIndex {
					firstIntegerIndex = k
				}

				if k >= lastIntegerIndex {
					lastIntegerIndex = k
				}
			}

		}

		calibrationValue := integers[firstIntegerIndex]*10 + integers[lastIntegerIndex]
		calibrationValues = append(calibrationValues, calibrationValue)
	}
	sum := 0

	for _, value := range calibrationValues {
		sum += value
	}
	fmt.Println("Day 1, Puzzle 2:", sum)
}

func findNumberStrings(line string) map[int]int {
	numberStrings := [...]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine", 0: "zero"}

	results := make(map[int]int)
	for numberInteger, numberString := range numberStrings {
		firstIndex := strings.Index(line, numberString)
		lastIndex := strings.LastIndex(line, numberString)

		if firstIndex != -1 {
			results[firstIndex] = numberInteger
		}
		if lastIndex != -1 {
			results[lastIndex] = numberInteger
		}
	}

	return results
}
