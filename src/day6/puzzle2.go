package main

import (
	"fmt"
)

func Puzzle2() {
	race := [2]int{49787980, 298118510661181}
	time := race[0]
	distance := race[1]

	var start int
	for i := 0; i < time/2; i++ {
		if i*(time-i) > distance {
			start = i
			break
		}
	}

	fmt.Println("Day 6, Puzzle 2:", (time - 2*start + 1))
}
