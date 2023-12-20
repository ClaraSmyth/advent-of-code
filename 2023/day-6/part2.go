package main

import (
	"fmt"
	"strings"
)

func part2(lines []string) string {
	time := stringToNum(strings.Join(strings.Fields(strings.Split(lines[0], ":")[1]), ""))
	dist := stringToNum(strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), ""))

	min := 0

	for j := 0; j < time; j++ {
		if j*(time-j) > dist {
			break
		}
		min += 1
	}

	total := time - min*2 + 1

	return fmt.Sprintf("%d", total)
}
