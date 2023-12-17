package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 17")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	total := 0

	matrix := [][]int{}

	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		for x, char := range line {

			if len(matrix) <= x {
				matrix = append(matrix, []int{})
			}

			charInt, _ := strconv.Atoi(string(char))
			matrix[x] = append(matrix[x], charInt)
		}
	}

	total = calculateHeatLoss(matrix)

	fmt.Println("Total:", total)

}

// Calculate with a priority queue and a cache
// As I was stuck I had to look for clues and the priority queue was one.
// Go has the heap pkg: https://pkg.go.dev/container/heap#Interface
// But is old and untyped because it's from before generics.
// I decided to do the implementation in GO with generics and thread safety. Could be a library.
func calculateHeatLoss(matrix [][]int) int {

	pq := &PriorityQueue[tile]{
		items: []*tile{{
			accumulatedHeat: 0,
			x:               0,
			y:               0,
			dir:             "",
			timesInDir:      0,
		}},
		valueFunc: func(t *tile) int {
			return t.accumulatedHeat
		},
	}

	type cacheEntry struct {
		x, y       int
		dir        string
		timesInDir int
	}

	cacheMap := map[cacheEntry]bool{}
	for currentTile, noMore := pq.Pop(); !noMore; currentTile, noMore = pq.Pop() {

		if currentTile.x == len(matrix)-1 && currentTile.y == len(matrix[0])-1 {
			return currentTile.accumulatedHeat
		}

		if currentTile.x < 0 || currentTile.x >= len(matrix) || currentTile.y < 0 || currentTile.y >= len(matrix[0]) {
			continue
		}
		if cacheMap[cacheEntry{x: currentTile.x, y: currentTile.y, dir: currentTile.dir, timesInDir: currentTile.timesInDir}] {
			continue
		}
		cacheMap[cacheEntry{x: currentTile.x, y: currentTile.y, dir: currentTile.dir, timesInDir: currentTile.timesInDir}] = true

		// Go in same direction
		if currentTile.timesInDir < 3 && currentTile.dir != "" {
			newX, newY := currentTile.x, currentTile.y
			switch currentTile.dir {
			case "up":
				newY--
			case "down":
				newY++
			case "left":
				newX--
			case "right":
				newX++
			}

			if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) {
				pq.Push(&tile{
					accumulatedHeat: currentTile.accumulatedHeat + matrix[newX][newY],
					x:               newX,
					y:               newY,
					dir:             currentTile.dir,
					timesInDir:      currentTile.timesInDir + 1,
				})
			}
		}

		// Go in other directions
		for _, dir := range []string{"up", "down", "left", "right"} {

			if dir == currentTile.dir {
				continue
			}

			newX, newY := currentTile.x, currentTile.y
			switch dir {
			case "up":
				if currentTile.dir == "down" {
					continue
				}
				newY--
			case "down":
				if currentTile.dir == "up" {
					continue
				}
				newY++
			case "left":
				if currentTile.dir == "right" {
					continue
				}
				newX--
			case "right":
				if currentTile.dir == "left" {
					continue
				}
				newX++
			}

			if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) {
				pq.Push(&tile{
					accumulatedHeat: currentTile.accumulatedHeat + matrix[newX][newY],
					x:               newX,
					y:               newY,
					dir:             dir,
					timesInDir:      1,
				})
			}

		}

	}
	return 0
}

type tile struct {
	accumulatedHeat int
	x               int
	y               int
	dir             string
	timesInDir      int
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
