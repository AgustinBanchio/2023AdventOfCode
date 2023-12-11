package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 11 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	total := 0

	universe := [][]string{}
	rowsToExpand := []int{}

	for i, line := range lines {
		if line == "" {
			continue
		}
		thisLine := []string{}
		for _, char := range line {
			thisLine = append(thisLine, string(char))
		}
		if !strings.Contains(line, "#") {
			rowsToExpand = append(rowsToExpand, i)
		}

		universe = append(universe, thisLine)
	}
	rows := len(universe)
	columns := len(universe[0])

	columnsToExpand := []int{}
A:
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			if universe[j][i] == "#" {
				continue A
			}
		}
		columnsToExpand = append(columnsToExpand, i)
	}

	type galaxy struct {
		x int
		y int
	}

	galaxies := []*galaxy{}
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			if universe[x][y] == "#" {
				galaxies = append(galaxies, &galaxy{y, x})
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

			expandedRowsPassed := 0
			expandedColumnsPassed := 0

			for x := minX; x <= maxX; x++ {
				if slices.Contains(columnsToExpand, x) {
					expandedColumnsPassed++
				}
			}

			for y := minY; y <= maxY; y++ {
				if slices.Contains(rowsToExpand, y) {
					expandedRowsPassed++
				}
			}

			distance := maxX - minX + maxY - minY + 999999*(expandedRowsPassed+expandedColumnsPassed)
			total += distance
		}
	}

	fmt.Println("Total:", total/2)

}
