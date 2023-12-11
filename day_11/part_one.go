package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 11")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	total := 0

	expandedUniverse := [][]string{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		thisLine := []string{}
		for _, char := range line {
			thisLine = append(thisLine, string(char))
		}
		// repeat to expand empty rows
		expandedUniverse = append(expandedUniverse, thisLine)
		if !strings.Contains(line, "#") {
			expandedUniverse = append(expandedUniverse, thisLine)
		}
	}
	rows := len(expandedUniverse)
	columns := len(expandedUniverse[0])

	columnsToExpand := []int{}
A:
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			if expandedUniverse[j][i] == "#" {
				continue A
			}
		}
		columnsToExpand = append(columnsToExpand, i)
	}
	for row, _ := range expandedUniverse {
		line := expandedUniverse[row]
		for colExpNum, column := range columnsToExpand {
			line = append(line[:column+colExpNum+1], line[column+colExpNum:]...)
			line[column+colExpNum] = "."
		}
		expandedUniverse[row] = line
	}

	columns = len(expandedUniverse[0])

	type galaxy struct {
		x int
		y int
	}

	galaxies := []*galaxy{}
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			if expandedUniverse[x][y] == "#" {
				galaxies = append(galaxies, &galaxy{x, y})
			}
		}
	}

	for i, galaxy := range galaxies {

		for j, otherGalaxy := range galaxies {
			if i == j {
				continue
			}
			maxX := max(galaxy.x, otherGalaxy.x)
			minX := min(galaxy.x, otherGalaxy.x)
			maxY := max(galaxy.y, otherGalaxy.y)
			minY := min(galaxy.y, otherGalaxy.y)
			distance := maxX - minX + maxY - minY
			total += distance
		}
	}

	fmt.Println("Total:", total/2)

}
