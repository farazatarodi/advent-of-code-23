package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputByte, inputErr := os.ReadFile("src/day1/input.text")

	if inputErr != nil {
		fmt.Print(inputErr)
	}

	inputString := string(inputByte)
	lines := strings.Split(inputString, "\n")

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
	fmt.Println(sum)
}
