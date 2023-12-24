package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type hailstone struct {
	posX, posY, posZ int
	velX, velY, velZ int
}

func main() {

	fmt.Println("Advent of Code 2023 - Day 24")

	inputBytes, _ := os.ReadFile("day_24/part_1/input.txt")
	input := string(inputBytes)

	total := 0

	hailstones := []*hailstone{}

	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		parts := strings.Split(line, "@")

		positions := strings.Split(strings.TrimSpace(parts[0]), ",")

		thisHailstone := &hailstone{}
		for i, pos := range positions {
			posNum, _ := strconv.Atoi(strings.TrimSpace(pos))
			switch i {
			case 0:
				thisHailstone.posX = posNum
			case 1:
				thisHailstone.posY = posNum
			case 2:
				thisHailstone.posZ = posNum
			}
		}

		velocities := strings.Split(strings.TrimSpace(parts[1]), ",")

		for i, vel := range velocities {
			velNum, _ := strconv.Atoi(strings.TrimSpace(vel))
			switch i {
			case 0:
				thisHailstone.velX = velNum
			case 1:
				thisHailstone.velY = velNum
			case 2:
				thisHailstone.velZ = velNum
			}
		}
		hailstones = append(hailstones, thisHailstone)
	}

	for i, hailA := range hailstones {

		for j := i + 1; j < len(hailstones); j++ {
			hailB := hailstones[j]

			if isIntersectionPointInArea(hailA, hailB, 200000000000000, 400000000000000, 200000000000000, 400000000000000) {
				total++
			}
		}

	}

	fmt.Println("Total:", total)

}

func isIntersectionPointInArea(A, B *hailstone, testMinX, testMaxX, testMinY, testMaxY int) bool {

	// Get intersection point of two lines
	slopeA, yInterceptA := calculateLine(A)
	slopeB, yInterceptB := calculateLine(B)

	if slopeA == slopeB {
		// parallel
		return false
	}

	xInterception := (yInterceptB - yInterceptA) / (slopeA - slopeB)
	yInterception := slopeA*xInterception + yInterceptA

	// check intersection point is not in the past
	if math.Signbit(xInterception-float64(A.posX)) != math.Signbit(float64(A.velX)) {
		return false
	}
	if math.Signbit(yInterception-float64(A.posY)) != math.Signbit(float64(A.velY)) {
		return false
	}
	if math.Signbit(xInterception-float64(B.posX)) != math.Signbit(float64(B.velX)) {
		return false
	}
	if math.Signbit(yInterception-float64(B.posY)) != math.Signbit(float64(B.velY)) {
		return false
	}

	// Check if intersection point is in area.
	if xInterception >= float64(testMinX) &&
		xInterception <= float64(testMaxX) &&
		yInterception >= float64(testMinY) &&
		yInterception <= float64(testMaxY) {
		return true
	}
	return false
}

// This is embarrassing but I had to get a math refresher from chat gpt.
func calculateLine(stone *hailstone) (float64, float64) {
	slope := float64(stone.velY) / float64(stone.velX)
	yIntercept := float64(stone.posY) - (slope * float64(stone.posX))
	return slope, yIntercept
}
