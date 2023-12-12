package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 12")

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

		//records := strings.Split(parts[0], ".")
		//
		//for i := 0; i < len(records); i++ {
		//	if records[i] == "" {
		//		records = append(records[:i], records[i+1:]...)
		//		i--
		//	}
		//}

		rows = append(rows, &row{records: parts[0], counts: countsInt})

	}

	for _, row := range rows {

		maxVariations := math.Pow(2, float64(strings.Count(row.records, "?")))

		for i := 0; i < int(maxVariations); i++ {

			numQuestion := strings.Count(row.records, "?")
			binaryNum := strconv.FormatInt(int64(i), 2)

			if len(binaryNum) < numQuestion {
				binaryNum = strings.Repeat("0", numQuestion-len(binaryNum)) + binaryNum
			}

			// do the variation
			variation := row.records

			for _, n := range binaryNum {
				if n == '0' {
					variation = strings.Replace(variation, "?", ".", 1)
				} else {
					variation = strings.Replace(variation, "?", "#", 1)
				}

			}

			if slices.Equal(getCounts(variation), row.counts) {
				//fmt.Println(row.records, variation)
				row.possibleVariations++
				total++
			}

		}
	}

	fmt.Println("Total:", total)

}

func getCounts(data string) []int {

	parts := strings.Split(data, ".")

	for i := 0; i < len(parts); i++ {
		if parts[i] == "" {
			parts = append(parts[:i], parts[i+1:]...)
			i--
		}
	}

	counts := []int{}

	for _, part := range parts {
		counts = append(counts, len(part))
	}
	return counts
}
