package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type tile struct {
	char       string
	energized  bool
	directions []string
}

func main() {

	fmt.Println("Advent of Code 2023 - Day 16")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	total := 0

	matrix := [][]*tile{}

	for _, line := range strings.Split(input, "\n") {

		for x, char := range line {

			if len(matrix) <= x {
				matrix = append(matrix, []*tile{})
			}
			matrix[x] = append(matrix[x], &tile{char: string(char)})
		}

	}

	processBeam(matrix, 0, 0, "right")

	for y := 0; y < len(matrix[0]); y++ {
		for x := 0; x < len(matrix); x++ {

			if matrix[x][y].energized {
				fmt.Printf("#")
				total++
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Println()
	}

	fmt.Println("Total:", total)

}

func processBeam(matrix [][]*tile, startX, startY int, direction string) {

	if startX < 0 || startY < 0 || startX >= len(matrix) || startY >= len(matrix[0]) {
		return
	}

	currentTile := matrix[startX][startY]
	currentTile.energized = true

	if slices.Contains(currentTile.directions, direction) {
		// already processed this tile in same direction
		return
	}
	currentTile.directions = append(currentTile.directions, direction)

	// empty space or pointy splitter
	if currentTile.char == "." ||
		(currentTile.char == "|" && (direction == "up" || direction == "down")) ||
		(currentTile.char == "-" && (direction == "left" || direction == "right")) {

		switch direction {
		case "right":
			processBeam(matrix, startX+1, startY, "right")
			return
		case "left":
			processBeam(matrix, startX-1, startY, "left")
			return
		case "up":
			processBeam(matrix, startX, startY-1, "up")
			return
		case "down":
			processBeam(matrix, startX, startY+1, "down")
			return
		}
	}

	// mirror
	if currentTile.char == "/" {
		switch direction {
		case "right":
			processBeam(matrix, startX, startY-1, "up")
			return
		case "left":
			processBeam(matrix, startX, startY+1, "down")
			return
		case "up":
			processBeam(matrix, startX+1, startY, "right")
			return
		case "down":
			processBeam(matrix, startX-1, startY, "left")
			return
		}
	}

	if currentTile.char == "\\" {
		switch direction {
		case "right":
			processBeam(matrix, startX, startY+1, "down")
			return
		case "left":
			processBeam(matrix, startX, startY-1, "up")
			return
		case "up":
			processBeam(matrix, startX-1, startY, "left")
			return
		case "down":
			processBeam(matrix, startX+1, startY, "right")
			return

		}
	}

	// splitter flat side
	if currentTile.char == "|" {
		processBeam(matrix, startX, startY-1, "up")
		processBeam(matrix, startX, startY+1, "down")
		return
	}
	if currentTile.char == "-" {
		processBeam(matrix, startX-1, startY, "left")
		processBeam(matrix, startX+1, startY, "right")
		return
	}

}
