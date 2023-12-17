package main

import (
	"fmt"
	"log"
	"os"
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

func part1(lines []string) int {
	total := 0

	for _, v := range lines {

		firstIndex := strings.IndexFunc(v, unicode.IsNumber)
		secondIndex := strings.LastIndexFunc(v, unicode.IsNumber)

		if firstIndex == -1 || secondIndex == -1 {
			// If no number in string break out
			break
		} else {
			// Get both numbers and combine them

			numAsString := string(v[firstIndex]) + string(v[secondIndex])

			num, err := strconv.Atoi(numAsString)
			if err != nil {
				panic(err)
			}

			total += num
		}
	}

	return total
}

func part2(lines []string) int {
	total := 0

	for _, v := range lines {

		numAsString := getNumberFromString(v)

		if numAsString == "-1-1" {
			continue
		}

		num, err := strconv.Atoi(numAsString)
		if err != nil {
			panic(err)
		}

		total += num
	}

	return total
}

func getNumberFromString(s string) string {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	smallestIndex := strings.IndexFunc(s, unicode.IsNumber)
	largestIndex := strings.LastIndexFunc(s, unicode.IsNumber)

	smallestNumber := -1
	largestNumber := -1

	if smallestIndex != -1 {

		num1, err := strconv.Atoi(string(s[smallestIndex]))
		if err != nil {
			panic(err)
		}

		num2, err := strconv.Atoi(string(s[largestIndex]))
		if err != nil {
			panic(err)
		}

		smallestNumber = num1
		largestNumber = num2
	}

	for i, v := range words {

		firstWordIndex := strings.Index(s, v)
		lastWordIndex := strings.LastIndex(s, v)

		if firstWordIndex == -1 || lastWordIndex == -1 {
			continue
		}

		if smallestIndex == -1 || smallestIndex >= firstWordIndex {
			smallestIndex = firstWordIndex
			smallestNumber = i + 1
		}

		if largestIndex == -1 || largestIndex <= lastWordIndex {
			largestIndex = lastWordIndex
			largestNumber = i + 1
		}
	}

	return fmt.Sprintf("%d%d", smallestNumber, largestNumber)
}
