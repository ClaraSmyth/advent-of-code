package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
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

func part1(lines []string) string {
	finalSum := 0

	for i, line := range lines {

		lineBelow := ""
		lineAbove := ""

		if i < len(lines)-1 {
			lineBelow = lines[i+1]
		}

		if i > 0 {
			lineAbove = lines[i-1]
		}

		numIndexes, symIndexes := getLineIndexes(line)
		_, symIndexesBelow := getLineIndexes(lineBelow)
		_, symIndexesAbove := getLineIndexes(lineAbove)

		totalNums := strings.FieldsFunc(line, func(r rune) bool {
			return r == 46 || !unicode.IsNumber(r)
		})

		checkedIndexes := 0
		for _, num := range totalNums {

			hasSymbol := false

			for j := 0; j < len(num); j++ {

				v := numIndexes[checkedIndexes+j]

				checkCurrentLine := slices.ContainsFunc(symIndexes, func(n int) bool {
					return n == v+1 || n == v-1
				})

				checkLineBelow := slices.ContainsFunc(symIndexesBelow, func(n int) bool {
					return n == v || n == v+1 || n == v-1
				})

				checkLineAbove := slices.ContainsFunc(symIndexesAbove, func(n int) bool {
					return n == v || n == v+1 || n == v-1
				})

				if checkLineBelow || checkCurrentLine || checkLineAbove {
					hasSymbol = true
					break
				}
			}

			if hasSymbol {
				numAsInt, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}

				finalSum += numAsInt
			}

			checkedIndexes += len(num)
		}

	}

	return fmt.Sprintf("%d", finalSum)
}

func part2(lines []string) string {
	return "Not finished"
}

func getLineIndexes(line string) ([]int, []int) {
	numIndexes := []int{}
	symIndexes := []int{}

	for j, char := range line {
		if unicode.IsNumber(char) {
			numIndexes = append(numIndexes, j)
		} else if char != 46 {
			symIndexes = append(symIndexes, j)
		}
	}

	return numIndexes, symIndexes
}
