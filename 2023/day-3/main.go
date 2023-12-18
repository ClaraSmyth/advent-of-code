package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	answerPart1 := part1(lines)
	answerPart2 := part2(lines)

	println("Part 1 -", answerPart1)
	println("Part 2 -", answerPart2)
}
