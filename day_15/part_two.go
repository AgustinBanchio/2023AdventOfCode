package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 15 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	// Remove new lines
	input = strings.ReplaceAll(input, "\n", "")

	total := 0

	steps := strings.Split(input, ",")

	for _, step := range steps {
		ProcessStep(step)
	}

	for hash, lenses := range lensesMap {

		for idx, lens := range lenses {
			total += (hash + 1) * (idx + 1) * lens.focalLength
		}

	}

	fmt.Println("Total:", total)

}

type lens struct {
	label       string
	focalLength int
}

var lensesMap = map[int][]*lens{}

func ProcessStep(s string) {

	parts := []string{}
	removing := false
	if strings.Contains(s, "-") {
		parts = strings.Split(s, "-")
		removing = true
	} else {
		parts = strings.Split(s, "=")
	}
	label := parts[0]
	hash := 0
	for _, c := range label {
		hash += int(c)
		hash *= 17
		hash %= 256
	}

	if removing {
		if _, ok := lensesMap[hash]; ok {
			for idx, lens := range lensesMap[hash] {
				if lens.label == label {
					lensesMap[hash] = append(lensesMap[hash][:idx], lensesMap[hash][idx+1:]...)
					break
				}
			}
		}
	} else {
		focalLength := 0
		focalLength, _ = strconv.Atoi(parts[1])

		if _, ok := lensesMap[hash]; !ok {
			lensesMap[hash] = []*lens{}
		}
		for _, l := range lensesMap[hash] {
			if l.label == label {
				l.focalLength = focalLength
				return
			}
		}

		lensesMap[hash] = append(lensesMap[hash], &lens{
			label:       label,
			focalLength: focalLength,
		})

	}

}
