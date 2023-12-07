package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 7 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	type hand struct {
		handType string
		cards    []string
		bet      int
	}

	hands := []hand{}

	total := 0

	for _, line := range lines {

		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")

		betNum, _ := strconv.Atoi(parts[1])
		cards := []string{}
		cardMap := map[string]int{}

		for _, card := range parts[0] {
			cards = append(cards, string(card))
			cardMap[string(card)]++
		}

		if cardMap["J"] != 0 {
			highestNonJ := ""
			highest := 0
			for card, count := range cardMap {
				if card != "J" && count > highest {
					highestNonJ = card
					highest = count
				}
			}
			cardMap[highestNonJ] += cardMap["J"]
			cardMap["J"] = 0
		}

		handType := "high"

		threeFound := false
		twoFound := false
		tCount := 0
		for _, count := range cardMap {
			tCount += count

			if count == 5 {
				handType = "five"
				break
			}

			if count == 4 {
				handType = "four"
				break
			}

			if count == 3 {
				threeFound = true
				if twoFound {
					handType = "full"
					break
				}
				if tCount > 3 {
					handType = "three"
					break
				}
				continue
			}

			if count == 2 {
				if twoFound {
					handType = "two"
					break
				}
				twoFound = true
				if threeFound {
					handType = "full"
					break
				}
				continue
			}

		}
		if handType == "high" {
			if threeFound {
				handType = "three"
			}
			if twoFound {
				handType = "one"
			}
		}

		hands = append(hands, hand{
			handType: handType,
			cards:    cards,
			bet:      betNum,
		})

	}

	rankMap := map[string]int{
		"high":  1,
		"one":   2,
		"two":   3,
		"three": 4,
		"full":  5,
		"four":  6,
		"five":  7,
	}

	cardRankMap := map[string]int{
		"J": 0,
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": 5,
		"7": 6,
		"8": 7,
		"9": 8,
		"T": 9,
		"Q": 11,
		"K": 12,
		"A": 13,
	}

	//neg when i < j
	//pos when i > j
	slices.SortFunc(hands, func(i, j hand) int {

		if rankMap[i.handType] < rankMap[j.handType] {
			return -1
		}
		if rankMap[i.handType] > rankMap[j.handType] {
			return 1
		}
		for k := 0; k < len(i.cards); k++ {
			if cardRankMap[i.cards[k]] < cardRankMap[j.cards[k]] {
				return -1
			}
			if cardRankMap[i.cards[k]] > cardRankMap[j.cards[k]] {
				return 1
			}
		}
		return 0
	})

	for idx, hand := range hands {
		total += hand.bet * (idx + 1)
	}

	fmt.Println("Total:", total)
}
