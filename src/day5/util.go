package main

func FindDestination(source int, mapInput [][]int) int {
	destination := 0
	for _, mapLine := range mapInput {
		if source >= mapLine[1] && source < mapLine[1]+mapLine[2] {
			diff := source - mapLine[1]
			destination = mapLine[0] + diff
		}
	}

	if destination == 0 {
		destination = source
	}

	return destination
}
