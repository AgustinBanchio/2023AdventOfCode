package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 4 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)
	total := 0

	type scratchCard struct {
		cardNum       int
		winningNums   []int
		availableNums []int
		count         int
	}

	cardsText := strings.Split(input, "\n")
	cards := []*scratchCard{}

	for _, card := range cardsText {
		if card == "" {
			continue
		}
		parts := strings.Split(card, ":")
		cardNum := strings.TrimLeft(strings.TrimSpace(parts[0]), "Card ")
		cardNumInt, _ := strconv.Atoi(cardNum)
		numSets := strings.Split(strings.TrimSpace(parts[1]), "|")

		winningNums := strings.Split(strings.TrimSpace(numSets[0]), " ")
		availableNums := strings.Split(strings.TrimSpace(numSets[1]), " ")

		winningNumsInt := []int{}
		availableNumsInt := []int{}

		for _, num := range winningNums {
			if strings.TrimSpace(num) == "" {
				continue
			}
			winNumInt, _ := strconv.Atoi(num)
			winningNumsInt = append(winningNumsInt, winNumInt)
		}

		for _, num := range availableNums {
			if strings.TrimSpace(num) == "" {
				continue
			}
			availNumInt, _ := strconv.Atoi(num)
			availableNumsInt = append(availableNumsInt, availNumInt)
		}

		cards = append(cards, &scratchCard{
			cardNum:       cardNumInt,
			winningNums:   winningNumsInt,
			availableNums: availableNumsInt,
			count:         1,
		})
	}

	for _, card := range cards {
		matches := []int{}

		for _, n := range card.availableNums {
			if slices.Contains(card.winningNums, n) {
				matches = append(matches, n)
			}
		}

		if len(matches) > 0 {
			for i := card.cardNum; i < card.cardNum+len(matches); i++ {
				for _, c := range cards {
					if c.cardNum == i+1 {
						c.count += card.count
					}
				}
			}
		}
		total += card.count
	}

	fmt.Println("Total:", total)
}
