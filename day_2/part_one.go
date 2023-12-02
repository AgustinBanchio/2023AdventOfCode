package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 2")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)
	total := 0

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	gameRegex := regexp.MustCompile("(?m)^Game (?P<id>\\d*):(?P<game>.*)$")

A:
	for _, line := range gameRegex.FindAllStringSubmatch(input, -1) {
		id, _ := strconv.Atoi(line[1])
		playedGame := line[2]

		for _, round := range strings.Split(playedGame, ";") {
			for _, colorPick := range strings.Split(round, ",") {
				colorVals := strings.Split(strings.TrimSpace(colorPick), " ")
				num, _ := strconv.Atoi(colorVals[0])
				switch colorVals[1] {
				case "red":
					if num > maxRed {
						continue A
					}
				case "green":
					if num > maxGreen {
						continue A
					}
				case "blue":
					if num > maxBlue {
						continue A
					}
				}

			}
		}
		total += id
	}

	fmt.Println("Total:", total)
}
