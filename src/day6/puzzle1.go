package main

import (
	"fmt"
)

func Puzzle1() {
	data := [...][2]int{{49, 298}, {78, 1185}, {79, 1066}, {80, 1181}}

	result := 1
	for _, race := range data {
		time := race[0]
		distance := race[1]

		var start int
		for i := 0; i < time/2; i++ {
			if i*(time-i) > distance {
				start = i
				break
			}
		}

		result = result * (time - 2*start + 1)
	}

	fmt.Println("Day 6, Puzzle 1:", result)
}
