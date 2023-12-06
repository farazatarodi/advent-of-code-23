package main

import (
	"fmt"
	"slices"
	"strconv"
)

func Puzzle2(lines []string) {
	stars := make(map[string][]string)

	for lineIndex, line := range lines {
		for characterIndex := 0; characterIndex < len(line); characterIndex++ {
			character := string(line[characterIndex])

			_, err := strconv.Atoi(character)
			if err == nil {

				var currentNumber string
				var nearbyStars []string

				for i := characterIndex; i < len(line); i++ {
					_, err2 := strconv.Atoi(string(line[i]))
					if err2 != nil {
						characterIndex = i
						break
					}

					currentNumber = currentNumber + string(line[i])

					for column := i - 1; column <= i+1; column++ {
						for row := lineIndex - 1; row <= lineIndex+1; row++ {
							if column >= 0 && row >= 0 && row < len(lines) && column < len(line) {
								target := string(lines[row][column])

								if target == "*" {
									id := strconv.Itoa(row) + "," + strconv.Itoa(column)
									if !slices.Contains(nearbyStars, id) {
										nearbyStars = append(nearbyStars, id)
									}
								}
							}
						}
					}
				}

				if len(nearbyStars) > 0 {
					for _, starId := range nearbyStars {
						stars[starId] = append(stars[starId], currentNumber)
					}
				}
			}

		}
	}

	sum := 0
	for _, starNumbers := range stars {
		if len(starNumbers) == 2 {
			number1, _ := strconv.Atoi(starNumbers[0])
			number2, _ := strconv.Atoi(starNumbers[1])
			sum = sum + number1*number2
		}
	}

	fmt.Println("Day 3, Puzzle 2:", sum)
}
