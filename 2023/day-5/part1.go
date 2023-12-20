package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func part1(lines []string) string {
	seeds, maps := formatLines(lines)

	finalValues := []int{}

	for _, seed := range seeds {
		output := seed

		for i := 0; i < len(maps)-1; i++ {

			currentMap := maps[i]

			for _, line := range currentMap {
				if output >= line[1] && output < (line[1]+line[2]) {
					output = output + (line[0] - line[1])
					break
				}
			}
		}

		finalValues = append(finalValues, output)
	}

	return fmt.Sprintf("%d", slices.Min(finalValues))
}

func stringToNum(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func formatLines(lines []string) ([]int, [][][]int) {
	seeds := []int{}

	maps := [][][]int{}

	for i, line := range lines {

		if i == 0 {

			seedStrings := strings.Fields(strings.Split(line, ":")[1])

			for _, seed := range seedStrings {
				seeds = append(seeds, stringToNum(seed))
			}
		}

		if len(line) == 0 {
			maps = append(maps, [][]int{})
			continue
		}

		if strings.Contains(line, ":") {
			continue
		}

		fields := strings.Fields(line)

		newMap := []int{}

		for _, num := range fields {
			newMap = append(newMap, stringToNum(num))
		}

		index := len(maps) - 1

		maps[index] = append(maps[index], newMap)

	}

	// Checks if there is an empty slice at the end and removes it
	if len(maps[len(maps)-1]) == 0 {
		maps = maps[:len(maps)-1]
	}

	return seeds, maps
}
