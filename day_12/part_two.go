package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 12 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	total := 0

	type row struct {
		records            string
		counts             []int
		possibleVariations int
	}

	rows := []*row{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		countsStr := strings.Split(parts[1], ",")
		countsInt := []int{}
		for _, countStr := range countsStr {
			num, _ := strconv.Atoi(countStr)
			countsInt = append(countsInt, num)
		}
		newCountsInt := []int{}
		for i := 0; i < 5; i++ {
			newCountsInt = append(newCountsInt, countsInt...)
		}

		records := parts[0] + "?" + parts[0] + "?" + parts[0] + "?" + parts[0] + "?" + parts[0]

		rows = append(rows, &row{records: records, counts: newCountsInt})

	}

	for _, row := range rows {
		total += countVariations(row.records, row.counts)
	}

	fmt.Println("Total:", total)

}

type cacheKey struct {
	records string
	counts  string
}

var cache = map[cacheKey]int{}

// Function I wrote by shamelessly looking for help on the internet after I had no idea where to start
// It's smart but I would have never come up with it
func countVariations(records string, counts []int) int {
	if records == "" {
		if len(counts) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(counts) == 0 {
		if strings.Contains(records, "#") {
			return 0
		} else {
			return 1
		}
	}

	countsStr := []string{}
	for _, count := range counts {
		countsStr = append(countsStr, strconv.Itoa(count))
	}
	key := cacheKey{records: records, counts: strings.Join(countsStr, ",")}

	if val, ok := cache[key]; ok {
		return val
	}

	result := 0

	if strings.Contains(".?", string(records[0])) {
		result += countVariations(records[1:], counts)
	}

	if strings.Contains("#?", string(records[0])) {
		if counts[0] <= len(records) &&
			!strings.Contains(records[:counts[0]], ".") &&
			(counts[0] == len(records) || records[counts[0]] != '#') {
			newCounts := counts
			newRecords := records
			if len(newRecords) == counts[0] {
				newRecords = ""
			} else {
				newRecords = newRecords[counts[0]+1:]
			}

			if len(newCounts) == 1 {
				newCounts = []int{}
			} else {
				newCounts = newCounts[1:]
			}

			result += countVariations(newRecords, newCounts)
		}

	}

	cache[key] = result
	return result
}
