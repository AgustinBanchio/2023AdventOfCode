package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 3")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)
	total := 0

	charsMatrix := [][]string{}

	type number struct {
		str   string
		value int
		line  int
		minX  int
		maxX  int
	}

	numbers := []number{}

	numberRegex := regexp.MustCompile(`\d`)

	for lineNum, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		lineChars := []string{}
		var currentNumber *number

		for xNum, char := range line {
			lineChars = append(lineChars, string(char))
			if numberRegex.MatchString(string(char)) {
				if currentNumber == nil {
					currentNumber = &number{
						str:  string(char),
						line: lineNum,
						minX: xNum,
					}
				} else {
					currentNumber.str += string(char)
				}
			} else {
				if currentNumber != nil {
					currentNumber.maxX = xNum - 1
					value, _ := strconv.Atoi(currentNumber.str)
					currentNumber.value = value
					numbers = append(numbers, *currentNumber)
					currentNumber = nil
				}
			}
		}
		if currentNumber != nil {
			currentNumber.maxX = len(lineChars) - 1
			value, _ := strconv.Atoi(currentNumber.str)
			currentNumber.value = value
			numbers = append(numbers, *currentNumber)
			currentNumber = nil
		}

		charsMatrix = append(charsMatrix, lineChars)
	}

	symbolRegex := regexp.MustCompile(`[^a-zA-Z0-9.]`)

A:
	for _, num := range numbers {

		minLine := max(0, num.line-1)
		maxLine := min(len(charsMatrix)-1, num.line+1)

		minX := max(0, num.minX-1)
		maxX := min(len(charsMatrix[0])-1, num.maxX+1)

		for lineNum := minLine; lineNum <= maxLine; lineNum++ {
			for xNum := minX; xNum <= maxX; xNum++ {
				if symbolRegex.MatchString(charsMatrix[lineNum][xNum]) {
					total += num.value
					continue A
				}
			}
		}

	}

	fmt.Println("Total:", total)
}
