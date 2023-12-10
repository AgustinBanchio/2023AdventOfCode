package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 9 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	total := 0

	type sequence struct {
		numbers            []int
		extrapolatedNumber int
	}

	sequences := []*sequence{}

	for _, line := range lines {

		if line == "" {
			continue
		}

		numbers := strings.Split(strings.TrimSpace(line), " ")
		numbInts := []int{}

		for _, number := range numbers {
			integer, _ := strconv.Atoi(strings.TrimSpace(number))
			numbInts = append(numbInts, integer)
		}

		allZeroes := false
		allLines := [][]int{}
		allLines = append(allLines, numbInts)
		currentLine := 0
	A:
		for !allZeroes {
			newLine := []int{}
			for i := 0; i < len(allLines[currentLine])-1; i++ {
				newLine = append(newLine, allLines[currentLine][i+1]-allLines[currentLine][i])
			}
			allLines = append(allLines, newLine)
			currentLine++
			for _, number := range newLine {
				if number != 0 {
					continue A
				}
			}
			allZeroes = true
		}

		predictedNum := 0

		for i := len(allLines) - 1; i >= 0; i-- {
			currLine := allLines[i]

			thisPartNum := 0
			if i != len(allLines)-1 {
				thisPartNum = currLine[0] - allLines[i+1][0]
			}
			currLine = append([]int{thisPartNum}, currLine...)
			allLines[i] = currLine
			if i == 0 {
				predictedNum += thisPartNum
			}
		}

		thisSequence := &sequence{
			numbers:            numbInts,
			extrapolatedNumber: predictedNum,
		}

		sequences = append(sequences, thisSequence)
	}

	for _, seq := range sequences {
		total += seq.extrapolatedNumber
	}

	fmt.Println("Total:", total)

}
