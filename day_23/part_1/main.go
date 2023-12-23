package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 23")

	inputBytes, _ := os.ReadFile("day_23/part_1/input.txt")
	input := string(inputBytes)

	total := 0

	matrix := [][]string{}

	startX := 0
	startY := 0

	endX := 0
	endY := 0

	for y, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		for x, char := range line {

			if char == '.' {
				if y == 0 {
					startX = x
					startY = y
				}
				if y == len(strings.Split(input, "\n"))-2 {
					endX = x
					endY = y
				}
			}

			if len(matrix) <= x {
				matrix = append(matrix, []string{})
			}

			matrix[x] = append(matrix[x], string(char))
		}

	}

	total = hike(matrix, startX, startY, endX, endY)

	fmt.Println("Total:", total)

}

func hike(matrix [][]string, startX, startY, endX, endY int) int {

	pq := &PriorityQueue[tile]{
		items: []*tile{{
			x:         startX,
			y:         startY,
			prevTiles: []*tile{},
		}},
		valueFunc: func(t *tile) int {
			return len(t.prevTiles)
		},
	}

	for currentTile, noMore := pq.Pop(); !noMore; currentTile, noMore = pq.Pop() {

		if currentTile.x == endX && currentTile.y == endY && pq.Len() == 0 {
			return len(currentTile.prevTiles)
		}

		possibleDirections := []string{}

		switch matrix[currentTile.x][currentTile.y] {
		case ".":
			possibleDirections = []string{"u", "d", "l", "r"}
		case "^":
			possibleDirections = []string{"u"}
		case "v":
			possibleDirections = []string{"d"}
		case "<":
			possibleDirections = []string{"l"}
		case ">":
			possibleDirections = []string{"r"}
		}

	DIRS:
		for _, dir := range possibleDirections {
			newX := currentTile.x
			newY := currentTile.y

			if dir == "u" {
				newY--
			}
			if dir == "d" {
				newY++
			}
			if dir == "l" {
				newX--
			}
			if dir == "r" {
				newX++
			}

			for _, prevTile := range currentTile.prevTiles {
				if prevTile.x == newX && prevTile.y == newY {
					continue DIRS
				}
			}

			if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) && matrix[newX][newY] != "#" {
				pq.Push(&tile{
					x:         newX,
					y:         newY,
					prevTiles: copyAppend(currentTile.prevTiles, currentTile),
				})
			}

		}
	}
	return 0
}

func copyAppend[T any](input []T, entry T) []T {
	output := make([]T, len(input)+1)
	copy(output, input)
	output[len(input)] = entry
	return output
}

type tile struct {
	x         int
	y         int
	prevTiles []*tile
}

type PriorityQueue[T any] struct {
	items     []*T
	valueFunc func(*T) int
	mu        sync.Mutex
}

func (pq *PriorityQueue[T]) Pop() (*T, bool) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if len(pq.items) == 0 {
		return nil, true
	}
	firstItem := pq.items[0]
	pq.items = pq.items[1:]
	return firstItem, false
}

func (pq *PriorityQueue[T]) Push(item *T) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.items = append(pq.items, item)
	slices.SortFunc(pq.items, func(i, j *T) int {
		return pq.valueFunc(i) - pq.valueFunc(j)
	})
}

func (pq *PriorityQueue[T]) Len() int {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return len(pq.items)
}
