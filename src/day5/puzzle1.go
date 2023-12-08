package main

import (
	"fmt"
	"slices"
)

func Puzzle1(
	seeds []int,
	seedToSoil [][]int,
	soilToFertilizer [][]int,
	fertilizerToWater [][]int,
	waterToLight [][]int,
	lightToTemperature [][]int,
	temperatureToHumidity [][]int,
	humidityToLocation [][]int,
) {
	var locations []int

	for _, seed := range seeds {
		soil := FindDestination(seed, seedToSoil)
		fertilizer := FindDestination(soil, soilToFertilizer)
		water := FindDestination(fertilizer, fertilizerToWater)
		light := FindDestination(water, waterToLight)
		temperature := FindDestination(light, lightToTemperature)
		humidity := FindDestination(temperature, temperatureToHumidity)
		location := FindDestination(humidity, humidityToLocation)

		locations = append(locations, location)
	}

	fmt.Println("Day 5, Puzzle 1:", slices.Min(locations))
}
