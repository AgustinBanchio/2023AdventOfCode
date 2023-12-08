package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 8")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	total := 0

	instructions := []string{}

	type direction struct {
		name  string
		left  string
		right string
	}

	directions := map[string]*direction{}

	for _, line := range lines {

		if line == "" {
			continue
		}

		if len(instructions) == 0 {
			for _, char := range line {
				instructions = append(instructions, string(char))
			}
			continue
		}

		parts := strings.Split(line, "=")

		name := strings.TrimSpace(parts[0])

		dirs := strings.TrimSpace(parts[1])
		dirs = strings.Replace(dirs, "(", "", -1)
		dirs = strings.Replace(dirs, ")", "", -1)
		dirsParts := strings.Split(dirs, ",")

		left := strings.TrimSpace(dirsParts[0])
		right := strings.TrimSpace(dirsParts[1])

		directions[name] = &direction{name, left, right}
	}

	current := "AAA"

	step := 0
	for current != "ZZZ" {

		inst := instructions[step]

		currDir := directions[current]
		if inst == "L" {
			current = currDir.left
		} else {
			current = currDir.right
		}

		step++
		total++
		if step >= len(instructions) {
			step = 0
		}

	}

	fmt.Println("Total:", total)
}
