package main

import (
	"fmt"
	"strconv"
)

func Puzzle1(lines []string) {
	sum := 0
	var numbers []string

	for lineIndex, line := range lines {
		for characterIndex := 0; characterIndex < len(line); characterIndex++ {
			character := string(line[characterIndex])

			_, err := strconv.Atoi(character)
			if err == nil {
				isValid := false

				var currentNumber string

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
								temp := string(lines[row][column])

								if temp != "." {
									_, err3 := strconv.Atoi(temp)

									if err3 != nil {
										isValid = true
									}

								}
							}
						}
					}

				}

				if isValid {
					numbers = append(numbers, currentNumber)
					num, _ := strconv.Atoi(currentNumber)
					sum = sum + num
				}
			}

		}
	}

	fmt.Println("Day 3, Puzzle 1:", sum)
}
