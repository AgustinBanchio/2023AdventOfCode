package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 1 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)
	total := 0

	firstDigitRegex := regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine")
	lastDigitRegex := regexp.MustCompile(".*(\\d|one|two|three|four|five|six|seven|eight|nine|zero)")

	for lineCount, line := range strings.Split(input, "\n") {

		matches := firstDigitRegex.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}
		firstDigit := strToDigit(matches[0])

		lastMatch := lastDigitRegex.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			continue
		}
		lastDigit := strToDigit(lastMatch[0][1])
		number, _ := strconv.Atoi(firstDigit + lastDigit)
		fmt.Println("Line", lineCount+1, ":", matches, "=>", number)
		total += number
	}

	fmt.Println("Total:", total)
}

func strToDigit(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	case "zero":
		return "0"
	}
	return s
}
