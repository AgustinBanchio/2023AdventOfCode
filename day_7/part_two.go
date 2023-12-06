package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 7 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {

		if line == "" {
			continue
		}

	}

	fmt.Println("Total:", total)
}
