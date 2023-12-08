package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 8 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

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

	currents := []*travels{}

	for dir := range directions {
		if strings.HasSuffix(dir, "A") {
			currents = append(currents, &travels{dir, dir, 0})
		}
	}
	allZ := false

	step := 0
A:
	for !allZ {
		inst := instructions[step]
		for i := 0; i < len(currents); i++ {
			if strings.HasSuffix(currents[i].end, "Z") {
				continue
			}
			currDir := directions[currents[i].end]
			if inst == "L" {
				currents[i].end = currDir.left
			} else {
				currents[i].end = currDir.right
			}
			currents[i].steps++
		}

		step++
		if step >= len(instructions) {
			step = 0
		}

		for _, curr := range currents {
			if !strings.HasSuffix(curr.end, "Z") {
				continue A
			}
		}
		allZ = true
	}

	fmt.Println("Total:", MCM(currents))
}

type travels struct {
	start string
	end   string
	steps int
}

func MCM(nums []*travels) int {

	dcm := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}

	result := nums[0].steps * nums[1].steps / dcm(nums[0].steps, nums[1].steps)

	for i := 2; i < len(nums); i++ {
		result = result * nums[i].steps / dcm(result, nums[i].steps)
	}

	return result
}
