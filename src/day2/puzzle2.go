package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Puzzle2(lines []string) {
	games := make(map[int]map[string]int)
	sum := 0

	for _, line := range lines {
		parsedLine := strings.Split(line, ":")

		gameId := strings.Split(parsedLine[0], " ")[1]
		gameIdInt, _ := strconv.Atoi((gameId))

		games[gameIdInt] = map[string]int{"red": 0, "green": 0, "blue": 0}

		sets := strings.Split(parsedLine[1], ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				parsedCube := strings.Split(strings.TrimSpace(cube), " ")

				cubeAmount, _ := strconv.Atoi(parsedCube[0])
				cubeColor := parsedCube[1]

				if games[gameIdInt][cubeColor] < cubeAmount {
					games[gameIdInt][cubeColor] = cubeAmount
				}
			}
		}

		power := games[gameIdInt]["red"] * games[gameIdInt]["green"] * games[gameIdInt]["blue"]
		sum += power

	}

	fmt.Println("Day 2, Puzzle 2:", sum)
}
