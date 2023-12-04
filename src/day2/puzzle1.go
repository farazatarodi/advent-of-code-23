package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Puzzle1(lines []string) {
	allowedAmounts := map[string]int{"red": 12, "green": 13, "blue": 14}
	var allowedGames []int

	for _, line := range lines {
		parsedLine := strings.Split(line, ":")

		gameId := strings.Split(parsedLine[0], " ")[1]
		gameIdInt, _ := strconv.Atoi((gameId))

		isAllowed := true

		sets := strings.Split(parsedLine[1], ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				parsedCube := strings.Split(strings.TrimSpace(cube), " ")

				cubeAmount, _ := strconv.Atoi(parsedCube[0])
				cubeColor := parsedCube[1]

				if allowedAmounts[cubeColor] < cubeAmount {
					isAllowed = false
				}
			}
		}

		if isAllowed {
			allowedGames = append(allowedGames, gameIdInt)
		}

	}

	sum := 0
	for _, id := range allowedGames {
		sum += id
	}

	fmt.Println("Day 2, Puzzle 1:", sum)
}
