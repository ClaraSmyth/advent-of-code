package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(lines []string) string {
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])

	total := 1

	for i, race := range times {

		time := stringToNum(race)
		dist := stringToNum(distances[i])

		winningMoves := 0

		for j := 0; j < time; j++ {
			if j*(time-j) > dist {
				winningMoves += 1
			}
		}

		total = total * winningMoves
	}

	return fmt.Sprintf("%d", total)
}

func stringToNum(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
