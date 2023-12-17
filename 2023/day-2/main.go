package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

func part1(lines []string) string {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	sumOfGames := 0

	for i, v := range lines {

		if v == "" {
			continue
		}

		pass := true

		currentId := i + 1

		stringWithoutId := strings.Split(v, ":")[1]

		allGames := strings.Split(stringWithoutId, ";")

		for _, game := range allGames {

			red := 0
			green := 0
			blue := 0

			fieldsSlice := strings.Fields(game)

			for j := 0; j < len(fieldsSlice)/2; j++ {
				numString := fieldsSlice[j*2]
				colorString := fieldsSlice[(j*2)+1]

				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(err)
				}

				if strings.Contains(colorString, "red") {
					red += num
				}

				if strings.Contains(colorString, "green") {
					green += num
				}

				if strings.Contains(colorString, "blue") {
					blue += num
				}
			}

			if red > maxRed || green > maxGreen || blue > maxBlue {
				pass = false
			}

		}

		if pass {
			sumOfGames += currentId
		}
	}

	return fmt.Sprintf("%d", sumOfGames)
}

func part2(lines []string) string {
	total := 0

	for _, v := range lines {

		if v == "" {
			continue
		}

		stringWithoutId := strings.Split(v, ":")[1]

		allGames := strings.Split(stringWithoutId, ";")

		minRed := 0
		minGreen := 0
		minBlue := 0

		for _, game := range allGames {

			red := 0
			green := 0
			blue := 0

			fieldsSlice := strings.Fields(game)

			for j := 0; j < len(fieldsSlice)/2; j++ {
				numString := fieldsSlice[j*2]
				colorString := fieldsSlice[(j*2)+1]

				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(err)
				}

				if strings.Contains(colorString, "red") {
					red += num
				}

				if strings.Contains(colorString, "green") {
					green += num
				}

				if strings.Contains(colorString, "blue") {
					blue += num
				}
			}

			if red > minRed {
				minRed = red
			}

			if green > minGreen {
				minGreen = green
			}

			if blue > minBlue {
				minBlue = blue
			}
		}

		total += minRed * minGreen * minBlue
	}

	return fmt.Sprintf("%d", total)
}
