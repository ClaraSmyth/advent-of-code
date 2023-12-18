package main

import (
	"fmt"
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

	fmt.Println("Part 1 -", answerPart1)
	fmt.Println("Part 2 -", answerPart2)
}
