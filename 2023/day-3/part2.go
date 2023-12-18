package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func part2(lines []string) string {
	finalSum := 0

	for i, line := range lines {
		if !strings.Contains(line, "*") {
			continue
		}

		lineBelow := ""
		lineAbove := ""

		if i < len(lines)-1 {
			lineBelow = lines[i+1]
		}

		if i > 0 {
			lineAbove = lines[i-1]
		}

		for j, char := range line {
			if char == 42 {

				left, right := checkLeftAndRight(line, j)
				allAbove := checkAboveOrBelow(lineAbove, j)
				allBelow := checkAboveOrBelow(lineBelow, j)

				allNumbers := append(allAbove, allBelow...)

				if left != "-1" {
					allNumbers = append(allNumbers, left)
				}

				if right != "-1" {
					allNumbers = append(allNumbers, right)
				}

				if len(allNumbers) == 2 {
					finalSum += convertStringToInt(allNumbers[0]) * convertStringToInt(allNumbers[1])
				}

			}
		}
	}

	return fmt.Sprintf("%d", finalSum)
}

func checkLeftAndRight(line string, index int) (string, string) {
	left := "-1"
	right := "-1"

	if index != 0 && unicode.IsNumber(rune(line[index-1])) {

		allNumsBefore := strings.FieldsFunc(line[:index], func(r rune) bool {
			return !unicode.IsNumber(r)
		})

		if len(allNumsBefore) > 0 {
			left = allNumsBefore[len(allNumsBefore)-1]
		}
	}

	if index < len(line)-1 && unicode.IsNumber(rune(line[index+1])) {
		allNumsAfter := strings.FieldsFunc(line[index+1:], func(r rune) bool {
			return !unicode.IsNumber(r)
		})

		if len(allNumsAfter) > 0 {
			right = allNumsAfter[0]
		}
	}

	return left, right
}

func checkAboveOrBelow(line string, index int) []string {
	allNumbers := []string{}

	if len(line) <= 0 {
		return allNumbers
	}

	center := string(line[index])

	left, right := checkLeftAndRight(line, index)

	if convertStringToInt(center) != -1 {
		if left != "-1" {
			center = fmt.Sprintf("%v%v", left, center)
		}

		if right != "-1" {
			center = fmt.Sprintf("%v%v", center, right)
		}

		allNumbers = append(allNumbers, center)
	} else {
		if left != "-1" {
			allNumbers = append(allNumbers, left)
		}

		if right != "-1" {
			allNumbers = append(allNumbers, right)
		}
	}

	return allNumbers
}

func convertStringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return num
}
