package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 1")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)
	total := 0

	numberRegex := regexp.MustCompile("(?mU)^[^0-9]*(?P<first>[0-9]).*(?P<last>[0-9])?[^0-9]*$")
	for _, line := range numberRegex.FindAllStringSubmatch(input, -1) {
		firstDigit := line[1]
		lastDigit := line[2]
		if lastDigit == "" {
			lastDigit = firstDigit
		}
		number, _ := strconv.Atoi(firstDigit + lastDigit)
		total += number
	}

	fmt.Println("Total:", total)
}
