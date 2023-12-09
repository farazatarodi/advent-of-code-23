package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Puzzle1(lines []string) {
	var hands [][]string
	for _, line := range lines {
		chunks := strings.Split(line, " ")
		hands = append(hands, chunks)
	}

	for mainIndex := 0; mainIndex < len(hands); mainIndex++ {
		mainHand := hands[mainIndex]
		handType := getHandType1(mainHand[0])

		mainHand = append(mainHand, handType)
		hands[mainIndex] = mainHand

		mainTypeNumber, _ := strconv.Atoi(handType)
		for i := 0; i < mainIndex; i++ {
			typeNumber, _ := strconv.Atoi(hands[i][2])

			if mainTypeNumber < typeNumber {
				temp := hands[i]
				hands[i] = hands[mainIndex]
				hands[mainIndex] = temp
			}

			if mainTypeNumber == typeNumber {
				for j := 0; j < len(hands[i][0]); j++ {
					mainChar := string(hands[mainIndex][0][j])
					targetChar := string(hands[i][0][j])

					mainCharInt := convertToInt1(mainChar)
					targetCharInt := convertToInt1(targetChar)

					if mainCharInt < targetCharInt {
						temp := hands[i]
						hands[i] = hands[mainIndex]
						hands[mainIndex] = temp
					}

					if mainCharInt == targetCharInt {
						continue
					}

					break
				}
			}
		}
	}

	sum := 0
	for index, hand := range hands {
		bid, _ := strconv.Atoi(hand[1])
		rank := index + 1
		sum += bid * rank
	}

	fmt.Println("Day 7, Puzzle 1:", sum)
}

func getHandType1(hand string) string {
	reduced := make(map[string]int)
	for _, card := range hand {
		reduced[string(card)]++
	}

	for _, amount := range reduced {
		if amount == 5 {
			return "7"
		}

		if amount == 4 {
			return "6"
		}

		if amount == 3 {
			if len(reduced) == 2 {
				return "5"
			} else {
				return "4"
			}
		}

		if amount == 2 {
			if len(reduced) == 3 {
				return "3"
			}
			if len(reduced) == 4 {
				return "2"
			}
		}
	}

	return "1"
}

func convertToInt1(card string) int {
	cardInt, err := strconv.Atoi(card)

	if err != nil {
		switch card {
		case "A":
			cardInt = 14
			break
		case "K":
			cardInt = 13
			break
		case "Q":
			cardInt = 12
			break
		case "J":
			cardInt = 11
			break
		case "T":
			cardInt = 10
			break
		default:
			cardInt = 0
		}
	}

	return cardInt
}
