package main

import (
	"strconv"
	"strings"

	readInputFile "github.com/farazatarodi/advent-of-code-23/src/utils"
)

func main() {
	seeds := readInputFile.ReadInputFile("src/day5/input/seeds.text")
	seedToSoil := readInputFile.ReadInputFile("src/day5/input/seed-to-soil.text")
	soilToFertilizer := readInputFile.ReadInputFile("src/day5/input/soil-to-fertilizer.text")
	fertilizerToWater := readInputFile.ReadInputFile("src/day5/input/fertilizer-to-water.text")
	waterToLight := readInputFile.ReadInputFile("src/day5/input/water-to-light.text")
	lightToTemperature := readInputFile.ReadInputFile("src/day5/input/light-to-temperature.text")
	temperatureToHumidity := readInputFile.ReadInputFile("src/day5/input/temperature-to-humidity.text")
	humidityToLocation := readInputFile.ReadInputFile("src/day5/input/humidity-to-location.text")

	seedsMap := parseLines(seeds)
	seedToSoilMap := parseLines(seedToSoil)
	soilToFertilizerMap := parseLines(soilToFertilizer)
	fertilizerToWaterMap := parseLines(fertilizerToWater)
	waterToLightMap := parseLines(waterToLight)
	lightToTemperatureMap := parseLines(lightToTemperature)
	temperatureToHumidityMap := parseLines(temperatureToHumidity)
	humidityToLocationMap := parseLines(humidityToLocation)

	Puzzle1(
		seedsMap[0],
		seedToSoilMap,
		soilToFertilizerMap,
		fertilizerToWaterMap,
		waterToLightMap,
		lightToTemperatureMap,
		temperatureToHumidityMap,
		humidityToLocationMap,
	)
	Puzzle2(
		seedsMap[0],
		seedToSoilMap,
		soilToFertilizerMap,
		fertilizerToWaterMap,
		waterToLightMap,
		lightToTemperatureMap,
		temperatureToHumidityMap,
		humidityToLocationMap,
	)

}

func parseLines(lines []string) [][]int {
	var parsedMap [][]int
	for _, line := range lines {
		chunks := strings.Split(line, " ")
		var numbers []int
		for _, chunk := range chunks {
			number, _ := strconv.Atoi(chunk)
			numbers = append(numbers, number)
		}
		parsedMap = append(parsedMap, numbers)
	}

	return parsedMap
}
