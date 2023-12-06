package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 6 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	times := []string{}
	distances := []string{}

	total := 0

	for _, line := range lines {

		if line == "" {
			continue
		}

		parts := strings.Split(strings.TrimSpace(line), ":")
		numbers := strings.Split(strings.TrimSpace(parts[1]), " ")

		timesLine := true
		if len(times) != 0 {
			timesLine = false
		}

		for _, num := range numbers {
			if num == "" {
				continue
			}
			if timesLine {
				times = append(times, num)
			} else {
				distances = append(distances, num)
			}
		}
	}

	timeStr := strings.Join(times, "")
	distanceStr := strings.Join(distances, "")

	timeInt, _ := strconv.Atoi(timeStr)
	distanceInt, _ := strconv.Atoi(distanceStr)

	for t := 0; t < timeInt; t++ {
		if (1)*t*(timeInt-t) > distanceInt {
			total++
		}
	}

	fmt.Println("Total:", total)
}
