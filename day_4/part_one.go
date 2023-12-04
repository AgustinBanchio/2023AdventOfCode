package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 4")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)
	total := 0

	cards := strings.Split(input, "\n")

	for _, card := range cards {
		if card == "" {
			continue
		}
		parts := strings.Split(card, ":")
		numSets := strings.Split(strings.TrimSpace(parts[1]), "|")

		winningNums := strings.Split(strings.TrimSpace(numSets[0]), " ")
		availableNums := strings.Split(strings.TrimSpace(numSets[1]), " ")

		winningNumsInt := []int{}
		availableNumsInt := []int{}

		matches := []int{}

		for _, num := range winningNums {
			if strings.TrimSpace(num) == "" {
				continue
			}
			winNumInt, _ := strconv.Atoi(num)
			winningNumsInt = append(winningNumsInt, winNumInt)
		}

		for _, num := range availableNums {
			if strings.TrimSpace(num) == "" {
				continue
			}
			availNumInt, _ := strconv.Atoi(num)
			availableNumsInt = append(availableNumsInt, availNumInt)
		}

		for _, n := range availableNumsInt {
			if slices.Contains(winningNumsInt, n) {
				matches = append(matches, n)
			}
		}

		if len(matches) > 0 {
			cardPoints := 1
			if len(matches) > 1 {
				cardPoints = int(math.Pow(2, float64(len(matches)-1)))
			}
			total += cardPoints
		}
	}

	fmt.Println("Total:", total)
}
