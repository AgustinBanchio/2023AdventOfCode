package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 3 -  Part Two")

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

	gearRegex := regexp.MustCompile(`\*`)

	for lineNum, line := range charsMatrix {
	A:
		for xNum, char := range line {
			if gearRegex.MatchString(char) {
				var num1, num2 int
				numCount := 0

				for _, num := range numbers {

					if num.line >= lineNum-1 && num.line <= lineNum+1 {
						if xNum >= num.minX-1 && xNum <= num.maxX+1 {
							numCount++
							if numCount == 3 {
								continue A
							}
							if numCount == 1 {
								num1 = num.value
							} else {
								num2 = num.value
							}
						}

					}

				}
				if numCount == 2 {
					total += num1 * num2
				}
			}
		}
	}

	fmt.Println("Total:", total)
}
