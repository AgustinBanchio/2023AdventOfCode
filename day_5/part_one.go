package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 5")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	seeds := []int{}

	type converter struct {
		name          string
		converterStep int
		steps         int
		source        int
		destination   int
	}

	converters := []*converter{}

	currentConverterName := ""
	currentConverterStep := 0

	for _, line := range lines {

		if line == "" {
			continue
		}

		parts := strings.Split(strings.TrimSpace(line), ":")

		if parts[0] == "seeds" {
			seedsNums := strings.Split(strings.TrimSpace(parts[1]), " ")
			for _, seedNum := range seedsNums {
				n, _ := strconv.Atoi(seedNum)
				seeds = append(seeds, n)
			}
		} else {
			if parts[0] != "" && strings.Contains(line, ":") {
				currentConverterName = parts[0]
				currentConverterStep++
				continue
			}
			if parts[0] != "" {
				converterNums := strings.Split(strings.TrimSpace(parts[0]), " ")
				converterNumsInts := []int{}
				for _, converterNum := range converterNums {
					n, _ := strconv.Atoi(converterNum)
					converterNumsInts = append(converterNumsInts, n)
				}
				converters = append(converters, &converter{
					name:          currentConverterName,
					converterStep: currentConverterStep,
					steps:         converterNumsInts[2],
					source:        converterNumsInts[1],
					destination:   converterNumsInts[0],
				})
			}

		}

	}

	type seedData struct {
		seed            int
		converterValues []int
		finalValue      int
	}

	compSeeds := []*seedData{}

	for _, seed := range seeds {

		compSeed := &seedData{
			seed:            seed,
			converterValues: []int{seed},
			finalValue:      seed,
		}

		compSeeds = append(compSeeds, compSeed)

		step := 0

		for _, converter := range converters {
			if converter.converterStep > step+1 {
				val := compSeed.finalValue
				compSeed.converterValues = append(compSeed.converterValues, val)
				compSeed.finalValue = val
				step++
			}

			if converter.converterStep == step+1 {
				if converter.source <= compSeed.finalValue && converter.source+converter.steps >= compSeed.finalValue {
					val := converter.destination + converter.steps - (converter.source + converter.steps - compSeed.finalValue)
					compSeed.converterValues = append(compSeed.converterValues, val)
					compSeed.finalValue = val
					step++
				}
			}

		}

	}

	var lowest *int
	for _, cs := range compSeeds {
		if lowest == nil || *lowest > cs.finalValue {
			lowest = &cs.finalValue
		}
	}

	fmt.Println("Lowest:", *lowest)
}
