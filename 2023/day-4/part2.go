package main

import (
	"fmt"
	"slices"
	"strings"
)

func part2(lines []string) string {
	total := 0

	allCards := map[int]int{}

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		allCards[i+1] += 1

		origialCard := checkCard(line)

		for l := 0; l < allCards[i+1]; l++ {
			for j := 1; j < origialCard+1; j++ {

				if i+j > len(lines)-1 {
					continue
				}

				allCards[(i+1)+j] += 1
			}
		}

		total += allCards[i+1]
	}

	return fmt.Sprintf("%d", total)
}

func checkCard(card string) int {
	numsOnly := strings.Split(strings.Split(card, ":")[1], "|")

	winningNums := strings.Fields(numsOnly[0])
	elfNums := strings.Fields(numsOnly[1])

	count := 0

	for _, num := range elfNums {

		isWinning := slices.Contains(winningNums, num)

		if isWinning {
			count += 1
		}
	}

	return count
}
