package main

import (
	"fmt"
	"strconv"
)

func Puzzle1(lines []string) {
	var calibrationValues []int
	for _, line := range lines {
		var integers []int
		for _, characterByte := range line {
			character := string(characterByte)

			integer, integerError := strconv.Atoi(character)

			if integerError == nil {
				integers = append(integers, integer)
			}

		}

		calibrationValue := integers[0]*10 + integers[len(integers)-1]
		calibrationValues = append(calibrationValues, calibrationValue)
	}
	sum := 0

	for _, value := range calibrationValues {
		sum += value
	}
	fmt.Println("Day 1, Puzzle 1:", sum)
}
