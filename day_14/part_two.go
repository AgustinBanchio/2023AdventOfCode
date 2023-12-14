package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 14 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	total := 0

	pattern := [][]string{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		for x, char := range line {

			if len(pattern) <= x {
				pattern = append(pattern, []string{string(char)})
			} else {
				pattern[x] = append(pattern[x], string(char))
			}

		}
	}

	loopHistory := []string{}
	found := false

	for i := 0; i < 1000000000; i++ {
		if i%1000 == 0 {
			fmt.Println("Cycle", i, "Cache size:", len(cache))
		}
		cycle(pattern)

		if i > 1000 && !found {

			patternString := ""
			for x := 0; x < len(pattern); x++ {
				patternString += strings.Join(pattern[x], "")
			}

			for t := len(loopHistory) - 1; t >= 0; t-- {

				if loopHistory[t] == patternString {
					// loop found
					loopLength := len(loopHistory) - t
					remainingCycles := 1000000000 - i - 1

					i += (remainingCycles / loopLength) * loopLength

				}

			}
			if !found {
				loopHistory = append(loopHistory, patternString)
			}
		}

	}

	for x := 0; x < len(pattern); x++ {

		for y := 0; y < len(pattern[x]); y++ {

			if pattern[x][y] == "O" {
				total += len(pattern[x]) - y
			}

		}

	}

	fmt.Println("Total:", total)

}

func cycle(pattern [][]string) {

	tiltNorth(pattern)
	tiltWest(pattern)
	tiltSouth(pattern)
	tiltEast(pattern)
}

type key struct {
	direction string
	input     string
}

var cache = map[key][]string{}

func tiltNorth(pattern [][]string) {

	for x := 0; x < len(pattern); x++ {

		k := key{
			direction: "north",
			input:     strings.Join(pattern[x], ""),
		}

		if _, ok := cache[k]; ok {
			copy(pattern[x], cache[k])
			continue
		}

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

		cache[k] = make([]string, len(pattern[x]))
		copy(cache[k], pattern[x])

	}
}

func tiltSouth(pattern [][]string) {
	for x := 0; x < len(pattern); x++ {

		k := key{
			direction: "south",
			input:     strings.Join(pattern[x], ""),
		}

		if _, ok := cache[k]; ok {
			copy(pattern[x], cache[k])
			continue
		}

		for y := len(pattern[x]) - 1; y >= 0; y-- {

			if pattern[x][y] == "O" {
				newPos := y
				for ty := y + 1; ty < len(pattern[x]); ty++ {

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

		cache[k] = make([]string, len(pattern[x]))
		copy(cache[k], pattern[x])

	}
}

func tiltWest(pattern [][]string) {
	for y := 0; y < len(pattern[0]); y++ {

		inputSlice := []string{}
		for x := 0; x < len(pattern); x++ {
			inputSlice = append(inputSlice, pattern[x][y])
		}

		k := key{
			direction: "west",
			input:     strings.Join(inputSlice, ""),
		}

		if _, ok := cache[k]; ok {
			for x := 0; x < len(pattern); x++ {
				pattern[x][y] = cache[k][x]
			}
			continue
		}

		for x := 0; x < len(pattern); x++ {

			if pattern[x][y] == "O" {
				newPos := x
				for tx := x - 1; tx >= 0; tx-- {

					if pattern[tx][y] == "." {
						newPos = tx
					} else if pattern[tx][y] == "O" || pattern[tx][y] == "#" {
						break
					}

				}
				pattern[x][y] = "."
				pattern[newPos][y] = "O"
			}

		}

		cache[k] = make([]string, len(pattern))
		for x := 0; x < len(pattern); x++ {
			cache[k][x] = pattern[x][y]
		}
	}
}

func tiltEast(pattern [][]string) {
	for y := 0; y < len(pattern[0]); y++ {

		inputSlice := []string{}
		for x := 0; x < len(pattern); x++ {
			inputSlice = append(inputSlice, pattern[x][y])
		}

		k := key{
			direction: "east",
			input:     strings.Join(inputSlice, ""),
		}

		if _, ok := cache[k]; ok {
			for x := 0; x < len(pattern); x++ {
				pattern[x][y] = cache[k][x]
			}
			continue
		}

		for x := len(pattern) - 1; x >= 0; x-- {

			if pattern[x][y] == "O" {
				newPos := x
				for tx := x + 1; tx < len(pattern); tx++ {

					if pattern[tx][y] == "." {
						newPos = tx
					} else if pattern[tx][y] == "O" || pattern[tx][y] == "#" {
						break
					}

				}
				pattern[x][y] = "."
				pattern[newPos][y] = "O"
			}

		}

		cache[k] = make([]string, len(pattern))
		for x := 0; x < len(pattern); x++ {
			cache[k][x] = pattern[x][y]
		}

	}
}
