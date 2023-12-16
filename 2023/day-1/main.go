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

	var total int64 = 0

	for _, v := range lines {

		firstNum := strings.IndexFunc(v, unicode.IsNumber)
		secondNum := strings.LastIndexFunc(v, unicode.IsNumber)

		if firstNum == -1 || secondNum == -1 {
			// If no number in string break out
			break
		} else {
			// Get both numbers and combine them

			numAsString := string(v[firstNum]) + string(v[secondNum])

			num, err := strconv.ParseInt(numAsString, 10, 64)
			if err != nil {
				panic(err)
			}

			total += num
		}
	}

	fmt.Println(total)
}
