package main

import (
	"fmt"
	"slices"
)

func part2(lines []string) string {
	seeds, maps := formatLines(lines)

	finalValues := []int{}

	for i := 0; i < len(seeds)/2; i++ {
		value1 := seeds[i*2]
		value2 := seeds[(i*2)+1]

		lowestOutput := checkSeed(value1, maps)

		for j := 0; j < value2; j++ {
			output := checkSeed(value1+j, maps)

			if output < lowestOutput {
				lowestOutput = output
			}
		}

		finalValues = append(finalValues, lowestOutput)
	}

	return fmt.Sprintf("%d", slices.Min(finalValues))
}

func checkSeed(seed int, maps [][][]int) int {
	output := seed

	for _, currentMap := range maps {
		for _, line := range currentMap {
			if output >= line[1] && output < (line[1]+line[2]) {
				output = output + (line[0] - line[1])
				break
			}
		}
	}

	return output
}
