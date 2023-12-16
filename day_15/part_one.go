package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 15")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	// Remove new lines
	input = strings.ReplaceAll(input, "\n", "")

	total := 0

	steps := strings.Split(input, ",")

	for _, step := range steps {
		total += HashString(step)
	}

	fmt.Println("Total:", total)

}

func HashString(s string) int {
	val := 0
	for _, c := range s {
		val += int(c)
		val *= 17
		val %= 256
	}
	return val
}
