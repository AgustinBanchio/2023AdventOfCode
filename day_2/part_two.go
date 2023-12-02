package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 2 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)
	total := 0

	gameRegex := regexp.MustCompile("(?m)^Game (?P<id>\\d*):(?P<game>.*)$")

	for _, line := range gameRegex.FindAllStringSubmatch(input, -1) {
		playedGame := line[2]

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, round := range strings.Split(playedGame, ";") {
			for _, colorPick := range strings.Split(round, ",") {
				colorVals := strings.Split(strings.TrimSpace(colorPick), " ")
				num, _ := strconv.Atoi(colorVals[0])
				switch colorVals[1] {
				case "red":
					if num > maxRed {
						maxRed = num
					}
				case "green":
					if num > maxGreen {
						maxGreen = num
					}
				case "blue":
					if num > maxBlue {
						maxBlue = num
					}
				}

			}
		}
		total += maxRed * maxGreen * maxBlue
	}

	fmt.Println("Total:", total)
}
