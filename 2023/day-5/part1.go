package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func part1(lines []string) string {
	maps := formatLines(lines)

	seeds := strings.Fields(maps[0][0])

	finalValues := []int{}

	for _, seed := range seeds {
		output := stringToNum(seed)

		for i := 0; i < len(maps); i++ {

			if i == 0 {
				continue
			}

			currentMap := maps[i]

			for _, line := range currentMap {

				values := strings.Fields(line)

				value1 := stringToNum(values[0])
				value2 := stringToNum(values[1])
				value3 := stringToNum(values[2])

				if output >= value2 && output < (value2+value3) {
					output = output + (value1 - value2)
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

func formatLines(lines []string) map[int][]string {
	maps := map[int][]string{}

	dataIndex := 0

	for i, line := range lines {

		if i == 0 {
			maps[dataIndex] = []string{strings.Split(line, ":")[1]}
		}

		if len(line) == 0 {
			dataIndex += 1
			continue
		}

		if strings.Contains(line, ":") {
			continue
		} else {
			maps[dataIndex] = append(maps[dataIndex], line)
		}
	}

	return maps
}
