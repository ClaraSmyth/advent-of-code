package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func part1(lines []string) string {
	total := 0

	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		numsOnly := strings.Split(strings.Split(line, ":")[1], "|")

		winningNums := strings.Fields(numsOnly[0])
		elfNums := strings.Fields(numsOnly[1])

		count := 0

		for _, num := range elfNums {

			isWinning := slices.Contains(winningNums, num)

			if isWinning && count == 0 {
				count += 1
			} else if isWinning {
				count = count * 2
			}
		}

		total += count
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
