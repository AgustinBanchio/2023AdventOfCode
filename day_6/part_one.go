package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 6")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	times := []int{}
	distances := []int{}
	winningWays := []int{}

	total := 1

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
			numInt, _ := strconv.Atoi(num)
			if timesLine {
				times = append(times, numInt)
			} else {
				distances = append(distances, numInt)
			}
		}
	}

	for i := 0; i < len(times); i++ {
		time := times[i]
		recordDistance := distances[i]
		numWiningWays := 0
		for t := 0; t < time; t++ {
			if (1)*t*(time-t) > recordDistance {
				numWiningWays++
			}
		}
		fmt.Println("Time:", time, "Record Distance:", recordDistance, "Winning Ways:", numWiningWays)
		winningWays = append(winningWays, numWiningWays)
	}

	for _, winningWay := range winningWays {
		total *= winningWay
	}

	fmt.Println("Total:", total)
}
