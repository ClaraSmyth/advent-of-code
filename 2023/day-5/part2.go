package main

import (
	"fmt"
)

func part2(lines []string) string {
	seeds, maps := formatLines(lines)

	seedsMinMax := getSeedsMinMax(seeds)

	finalValue := -1

	for i, currentMap := range maps {
		for _, line := range currentMap {

			val1 := line[1]
			val2 := val1 + line[2]

			for j := 0; j < i+1; j++ {
				if i-(1+j) != -1 {

					val1 = checkValue(val1, maps[i-(1+j)])
					val2 = checkValue(val2, maps[i-(1+j)])
				}
			}

			for _, seed := range seedsMinMax {
				if val1 > seed[0] && val1 < seed[1] {

					seedVal := checkSeed(val1, maps)

					if finalValue == -1 {
						finalValue = seedVal
					} else if seedVal < finalValue {
						finalValue = seedVal
					}

				}

				if val2 > seed[0] && val2 < seed[1] {

					seedVal := checkSeed(val2, maps)

					if finalValue == -1 {
						finalValue = seedVal
					} else if seedVal < finalValue {
						finalValue = seedVal
					}
				}
			}
		}
	}

	return fmt.Sprintf("%d", finalValue)
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

func getSeedsMinMax(seeds []int) [][]int {
	finalSeeds := [][]int{}

	for i := 0; i < len(seeds)/2; i++ {
		min := seeds[i*2]
		max := min + seeds[(i*2)+1]

		finalSeeds = append(finalSeeds, []int{min, max})
	}

	return finalSeeds
}

func checkValue(value int, anyMap [][]int) int {
	for _, line := range anyMap {
		min := line[0]
		max := min + line[2]

		if value >= min && value < max {

			value = value + (line[1] - line[0])
			break
		}
	}

	return value
}
