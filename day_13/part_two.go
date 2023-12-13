package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 13 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	total := 0

	patterns := [][][]string{}

	newPattern := [][]string{}
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, newPattern)
			newPattern = [][]string{}
			continue
		}

		for x, char := range line {

			if len(newPattern) <= x {
				newPattern = append(newPattern, []string{string(char)})
			} else {
				newPattern[x] = append(newPattern[x], string(char))
			}

		}
	}

	verticals := []int{}
	horizontals := []int{}

	for _, pattern := range patterns {

		horizontalReflection := 0
		verticalReflection := 0

	Vert:
		for x := 0; x < len(pattern)-1; x++ {

			if slices.Equal(pattern[x], pattern[x+1]) {
				// possible vertical reflection
				for tx := 0; tx < min(x, len(pattern)-x-2); tx++ {

					if !slices.Equal(pattern[x-1-tx], pattern[x+2+tx]) {
						continue Vert
					}

				}
				verticalReflection = x + 1
				verticals = append(verticals, verticalReflection)
				break Vert
			}
		}

	Horiz:
		for y := 0; y < len(pattern[0])-1; y++ {

			if columnEquals(pattern, y, y+1) {
				// possible horizontal reflection

				for ty := 0; ty < min(y, len(pattern[0])-y-2); ty++ {

					if !columnEquals(pattern, y-1-ty, y+2+ty) {
						continue Horiz
					}

				}
				horizontalReflection = y + 1
				horizontals = append(horizontals, horizontalReflection)
				break Horiz
			}

		}

	}

	fmt.Println(verticals, horizontals)

	for _, vertical := range verticals {
		total += vertical
	}

	for _, horizontal := range horizontals {
		total += 100 * horizontal
	}

	fmt.Println("Total:", total)

}

func columnEquals(matrix [][]string, y1, y2 int) bool {
	for x := 0; x < len(matrix); x++ {
		if matrix[x][y1] != matrix[x][y2] {
			return false
		}
	}
	return true
}
