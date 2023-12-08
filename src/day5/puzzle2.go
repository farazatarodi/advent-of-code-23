package main

import (
	"fmt"
)

func Puzzle2(
	seeds []int,
	seedToSoil [][]int,
	soilToFertilizer [][]int,
	fertilizerToWater [][]int,
	waterToLight [][]int,
	lightToTemperature [][]int,
	temperatureToHumidity [][]int,
	humidityToLocation [][]int,
) {
	minLocation := 99999999999999999

	for index, seed := range seeds {
		if index%2 == 0 {
			for j := 0; j < seeds[index+1]; j++ {
				soil := FindDestination(seed+j, seedToSoil)
				fertilizer := FindDestination(soil, soilToFertilizer)
				water := FindDestination(fertilizer, fertilizerToWater)
				light := FindDestination(water, waterToLight)
				temperature := FindDestination(light, lightToTemperature)
				humidity := FindDestination(temperature, temperatureToHumidity)
				location := FindDestination(humidity, humidityToLocation)

				if location < minLocation {
					minLocation = location
				}
			}
		}
	}

	fmt.Println("Day 5, Puzzle 2:", minLocation)
}
