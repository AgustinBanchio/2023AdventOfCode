package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"strconv"
	"strings"
)

type converter struct {
	sources      []int
	ranges       []int
	destinations []int
}

func (c converter) getDest(input int) int {
	for i := 0; i < len(c.sources); i++ {
		if input >= c.sources[i] && input < c.sources[i]+c.ranges[i] {
			return c.destinations[i] + (input - c.sources[i])
		}
	}
	return input
}
func main() {

	fmt.Println("Advent of Code 2023 - Day 5 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	preSeeds := []int{}

	converters := []*converter{}

	starts := []int{}
	ends := []int{}
	ranges := []int{}

	for _, line := range lines {

		if line == "" {
			continue
		}

		parts := strings.Split(strings.TrimSpace(line), ":")

		if parts[0] == "seeds" {
			seedsNums := strings.Split(strings.TrimSpace(parts[1]), " ")
			for _, seedNum := range seedsNums {
				n, _ := strconv.Atoi(seedNum)
				preSeeds = append(preSeeds, n)
			}
		} else {
			if parts[0] != "" && strings.Contains(line, ":") && len(starts) > 0 {
				converters = append(converters, &converter{
					sources:      starts,
					ranges:       ranges,
					destinations: ends,
				})
				starts = []int{}
				ends = []int{}
				ranges = []int{}
				continue
			}
			if parts[0] != "" && !strings.Contains(line, ":") {
				converterNums := strings.Split(strings.TrimSpace(parts[0]), " ")
				converterNumsInts := []int{}
				for _, converterNum := range converterNums {
					n, _ := strconv.Atoi(converterNum)
					converterNumsInts = append(converterNumsInts, n)
				}

				starts = append(starts, converterNumsInts[1])
				ends = append(ends, converterNumsInts[0])
				ranges = append(ranges, converterNumsInts[2])
			}

		}

	}

	var lowest *int

	resultChannel := make(chan int)
	group, _ := errgroup.WithContext(context.Background())

	for i := 0; i < len(preSeeds)/2; i++ {
		initial := preSeeds[i*2]
		seedRange := preSeeds[(i*2)+1]
		group.Go(func() error {
			var lowest *int

			for seed := initial; seed < initial+seedRange; seed++ {
				val := seed
				for _, converter := range converters {
					val = converter.getDest(val)
				}

				if lowest == nil || *lowest > val {
					lowest = &val
				}
			}

			resultChannel <- *lowest
			return nil
		})
	}

	for count := 0; count < len(preSeeds)/2; count++ {
		select {
		case result := <-resultChannel:
			fmt.Println("checkpoint")
			if lowest == nil || *lowest > result {
				lowest = &result
			}
		}
	}

	fmt.Println("Lowest:", *lowest)
}
