package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 10")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	total := 0

	sPositionX := 0
	sPositionY := 0

A:
	for idx, line := range lines {

		if line == "" {
			continue
		}

		for jdx, char := range line {

			if string(char) == "S" {
				sPositionX = jdx
				sPositionY = idx
				break A
			}

		}

	}

	foundLoop := false
	currentX := sPositionX
	currentY := sPositionY
	sValue := ""
	loopSteps := 0

	possibleStartingDirections := []string{}

	if currentY-1 >= 0 {
		if _, ok := getDirection(string(lines[currentY-1][currentX]), "down"); ok {
			possibleStartingDirections = append(possibleStartingDirections, "down")
		}
	}

	if currentY+1 < len(lines) {
		if _, ok := getDirection(string(lines[currentY+1][currentX]), "up"); ok {
			possibleStartingDirections = append(possibleStartingDirections, "up")
		}
	}

	if currentX-1 >= 0 {
		if _, ok := getDirection(string(lines[currentY][currentX-1]), "right"); ok {
			possibleStartingDirections = append(possibleStartingDirections, "right")
		}
	}

	if currentX+1 < len(lines[currentY]) {
		if _, ok := getDirection(string(lines[currentY][currentX+1]), "left"); ok {
			possibleStartingDirections = append(possibleStartingDirections, "left")
		}
	}

B:
	for _, possibleStartDir := range possibleStartingDirections {
		runDied := false
		currentX = sPositionX
		currentY = sPositionY
		currentDir := possibleStartDir
		thisSteps := 0
		var ok bool

		for !runDied && !foundLoop {
			thisSteps++
			if thisSteps != 0 {
				if currentDir == "up" && currentY+1 == sPositionY && currentX == sPositionX ||
					currentDir == "down" && currentY-1 == sPositionY && currentX == sPositionX ||
					currentDir == "left" && currentX+1 == sPositionX && currentY == sPositionY ||
					currentDir == "right" && currentX-1 == sPositionX && currentY == sPositionY {
					foundLoop = true
					loopSteps = thisSteps
					sValue = fmt.Sprintf("%s%s", currentDir, possibleStartDir)
					break B
				}
			}

			if currentDir == "up" {
				if currentY+1 >= len(lines) {
					runDied = true
					continue B
				}
				currentDir, ok = getDirection(string(lines[currentY+1][currentX]), currentDir)
				if !ok {
					runDied = true
					continue B
				} else {
					currentY++
					continue
				}
			}

			if currentDir == "down" {
				if currentY-1 < 0 {
					runDied = true
					continue B
				}
				currentDir, ok = getDirection(string(lines[currentY-1][currentX]), currentDir)
				if !ok {
					runDied = true
					continue B
				} else {
					currentY--
					continue
				}
			}

			if currentDir == "left" {
				if currentX+1 >= len(lines[currentY]) {
					runDied = true
					continue B
				}
				currentDir, ok = getDirection(string(lines[currentY][currentX+1]), currentDir)
				if !ok {
					runDied = true
					continue B
				} else {
					currentX++
					continue
				}
			}

			if currentDir == "right" {
				if currentX-1 < 0 {
					runDied = true
					continue B
				}
				currentDir, ok = getDirection(string(lines[currentY][currentX-1]), currentDir)
				if !ok {
					runDied = true
					continue B
				} else {
					currentX--
					continue
				}
			}

		}

	}
	total = loopSteps / 2

	fmt.Println("Total:", total)
	fmt.Println("S value:", sValue)
}

func getDirection(char string, inputDir string) (string, bool) {

	if char == "." {
		return "", false
	}

	switch inputDir {
	case "up":
		{
			if char == "|" {
				return "up", true
			}
			if char == "L" {
				return "left", true
			}
			if char == "J" {
				return "right", true
			}
		}
	case "down":
		{
			if char == "|" {
				return "down", true
			}
			if char == "7" {
				return "right", true
			}
			if char == "F" {
				return "left", true
			}
		}
	case "left":
		{
			if char == "-" {
				return "left", true
			}
			if char == "J" {
				return "down", true
			}
			if char == "7" {
				return "up", true
			}
		}
	case "right":
		{
			if char == "-" {
				return "right", true
			}
			if char == "L" {
				return "down", true
			}
			if char == "F" {
				return "up", true
			}
		}
	}
	return "", false
}
