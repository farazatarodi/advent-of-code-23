package main

import (
	readInputFile "github.com/farazatarodi/advent-of-code-23/src/utils"
)

func main() {
	lines := readInputFile.ReadInputFile("src/day4/input.text")
	Puzzle1(lines)
	Puzzle2(lines)
}
