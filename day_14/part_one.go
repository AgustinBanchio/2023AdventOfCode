package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 14")

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

	for _, pattern := range patterns {
		for x := 0; x < len(pattern); x++ {

			for y := 0; y < len(pattern[x]); y++ {

				if pattern[x][y] == "O" {
					newPos := y
					for ty := y - 1; ty >= 0; ty-- {

						if pattern[x][ty] == "." {
							newPos = ty
						} else if pattern[x][ty] == "O" || pattern[x][ty] == "#" {
							break
						}

					}
					pattern[x][y] = "."
					pattern[x][newPos] = "O"
				}

			}

		}
	}

	for _, pattern := range patterns {
		for x := 0; x < len(pattern); x++ {

			for y := 0; y < len(pattern[x]); y++ {

				if pattern[x][y] == "O" {
					total += len(pattern[x]) - y
				}

			}

		}
	}

	fmt.Println("Total:", total)

}
