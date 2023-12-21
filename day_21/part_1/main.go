package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 21")

	inputBytes, _ := os.ReadFile("day_21/part_1/input.txt")
	input := string(inputBytes)

	total := 0

	matrix := [][]string{}

	startX, startY := 0, 0

	for y, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		for x, char := range line {

			if len(matrix) <= x {
				matrix = append(matrix, []string{})
			}

			if char == 'S' {
				startX, startY = x, y
			}

			matrix[x] = append(matrix[x], string(char))
		}

	}

	type position struct {
		x, y           int
		remainingSteps int
	}

	positionChannel := make(chan position, 5000)

	positionChannel <- position{startX, startY, 64}

	resultsMatrix := map[struct{ x, y int }]bool{}

	for currentPos := range positionChannel {

		if resultsMatrix[struct{ x, y int }{currentPos.x, currentPos.y}] {
			continue
		}

		if currentPos.remainingSteps%2 == 0 {
			resultsMatrix[struct{ x, y int }{currentPos.x, currentPos.y}] = true
		}

		if currentPos.remainingSteps == 0 {

			if len(positionChannel) == 0 {
				close(positionChannel)
			}

			continue
		}

		if currentPos.x+1 < len(matrix) && matrix[currentPos.x+1][currentPos.y] != "#" {
			positionChannel <- position{currentPos.x + 1, currentPos.y, currentPos.remainingSteps - 1}
		}

		if currentPos.x-1 >= 0 && matrix[currentPos.x-1][currentPos.y] != "#" {
			positionChannel <- position{currentPos.x - 1, currentPos.y, currentPos.remainingSteps - 1}
		}

		if currentPos.y+1 < len(matrix[currentPos.x]) && matrix[currentPos.x][currentPos.y+1] != "#" {
			positionChannel <- position{currentPos.x, currentPos.y + 1, currentPos.remainingSteps - 1}
		}

		if currentPos.y-1 >= 0 && matrix[currentPos.x][currentPos.y-1] != "#" {
			positionChannel <- position{currentPos.x, currentPos.y - 1, currentPos.remainingSteps - 1}
		}

	}

	for _, entry := range resultsMatrix {
		if entry {
			total++
		}
	}

	fmt.Println("Total:", total)

}
