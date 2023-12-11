package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type tile struct {
	char         string
	partOfLoop   bool
	inputDir     string
	nextInputDir string
	outside      bool
}

func main() {

	fmt.Println("Advent of Code 2023 - Day 10 - Part Two")

	inputBytes, _ := os.ReadFile("input.txt")
	input := string(inputBytes)

	lines := strings.Split(input, "\n")

	total := 0

	sPositionX := 0
	sPositionY := 0

	tileMap := [][]*tile{}

	for y, line := range lines {

		if line == "" {
			continue
		}

		for x, char := range strings.TrimSpace(line) {
			if len(tileMap) <= x {
				tileMap = append(tileMap, []*tile{})
			}

			if tileMap[x] == nil {
				tileMap[x] = []*tile{}
			}

			tileMap[x] = append(tileMap[x], &tile{
				char:         string(char),
				partOfLoop:   false,
				inputDir:     "",
				nextInputDir: "",
			})

			if string(char) == "S" {
				sPositionX = x
				sPositionY = y
			}
		}

	}

	foundLoop := false
	tileMap[sPositionX][sPositionY].partOfLoop = true
	sValue := ""

	possibleStartingDirections := []string{}

	if sPositionY-1 >= 0 {
		if _, ok := getDirection(tileMap[sPositionX][sPositionY-1].char, "down"); ok {
			possibleStartingDirections = append(possibleStartingDirections, "down")
		}
	}

	if sPositionY+1 < len(lines) {
		if _, ok := getDirection(tileMap[sPositionX][sPositionY+1].char, "up"); ok {
			possibleStartingDirections = append(possibleStartingDirections, "up")
		}
	}

	if sPositionX-1 >= 0 {
		if _, ok := getDirection(tileMap[sPositionX-1][sPositionY].char, "right"); ok {
			possibleStartingDirections = append(possibleStartingDirections, "right")
		}
	}

	if sPositionX+1 < len(lines[sPositionY]) {
		if _, ok := getDirection(tileMap[sPositionX+1][sPositionY].char, "left"); ok {
			possibleStartingDirections = append(possibleStartingDirections, "left")
		}
	}

	possibleTiles := []*tile{}

B:
	for _, possibleStartDir := range possibleStartingDirections {
		currentX := sPositionX
		currentY := sPositionY
		currentDir := possibleStartDir
		thisSteps := 0

		if len(possibleTiles) > 0 {
			for _, t := range possibleTiles {
				t.partOfLoop = false
				t.inputDir = ""
				t.nextInputDir = ""
			}
		}
		possibleTiles = []*tile{}

		currentTile := tileMap[currentX][currentY]
		currentTile.partOfLoop = true
		currentTile.nextInputDir = currentDir
		possibleTiles = append(possibleTiles, currentTile)

		for !foundLoop {
			thisSteps++

			if thisSteps != 0 {
				if currentDir == "up" && currentY+1 == sPositionY && currentX == sPositionX ||
					currentDir == "down" && currentY-1 == sPositionY && currentX == sPositionX ||
					currentDir == "left" && currentX+1 == sPositionX && currentY == sPositionY ||
					currentDir == "right" && currentX-1 == sPositionX && currentY == sPositionY {
					foundLoop = true
					tileMap[sPositionX][sPositionY].inputDir = currentDir
					sValue = fmt.Sprintf("%s%s", currentDir, possibleStartDir)
					break B
				}
			}

			if currentDir == "up" {
				if currentY+1 >= len(tileMap[0]) {
					continue B
				}
				newDir, ok := getDirection(tileMap[currentX][currentY+1].char, currentDir)
				if !ok {
					continue B
				} else {
					currentY++
					currentTile.nextInputDir = currentDir
					currentTile = tileMap[currentX][currentY]
					currentTile.partOfLoop = true
					currentTile.inputDir = currentDir
					currentDir = newDir
					possibleTiles = append(possibleTiles, currentTile)
					continue
				}
			}

			if currentDir == "down" {
				if currentY-1 < 0 {
					continue B
				}
				newDir, ok := getDirection(tileMap[currentX][currentY-1].char, currentDir)
				if !ok {
					continue B
				} else {
					currentY--
					currentTile.nextInputDir = currentDir
					currentTile = tileMap[currentX][currentY]
					currentTile.partOfLoop = true
					currentTile.inputDir = currentDir
					currentDir = newDir
					possibleTiles = append(possibleTiles, currentTile)
					continue
				}
			}

			if currentDir == "left" {
				if currentX+1 >= len(tileMap) {
					continue B
				}
				newDir, ok := getDirection(tileMap[currentX+1][currentY].char, currentDir)
				if !ok {
					continue B
				} else {
					currentX++
					currentTile.nextInputDir = currentDir
					currentTile = tileMap[currentX][currentY]
					currentTile.partOfLoop = true
					currentTile.inputDir = currentDir
					currentDir = newDir
					possibleTiles = append(possibleTiles, currentTile)
					continue
				}
			}

			if currentDir == "right" {
				if currentX-1 < 0 {
					continue B
				}
				newDir, ok := getDirection(tileMap[currentX-1][currentY].char, currentDir)
				if !ok {
					continue B
				} else {
					currentX--
					currentTile.nextInputDir = currentDir
					currentTile = tileMap[currentX][currentY]
					currentTile.partOfLoop = true
					currentTile.inputDir = currentDir
					currentDir = newDir
					possibleTiles = append(possibleTiles, currentTile)
					continue
				}
			}

		}

	}

	if !tileMap[0][0].partOfLoop {
		tileMap[0][0].outside = true
	}

	if !tileMap[len(tileMap)-1][0].partOfLoop {
		tileMap[len(tileMap)-1][0].outside = true
	}

	if !tileMap[0][len(tileMap[0])-1].partOfLoop {
		tileMap[0][len(tileMap[0])-1].outside = true
	}

	if !tileMap[len(tileMap)-1][len(tileMap[0])-1].partOfLoop {
		tileMap[len(tileMap)-1][len(tileMap[0])-1].outside = true
	}

	for x := 0; x < len(tileMap); x++ {
		for y := 0; y < len(tileMap[0]); y++ {
			if tileMap[x][y].outside {
				markNeighboursOutside(x, y, tileMap)
			}
		}
	}

	for x := 0; x < len(tileMap); x++ {
		for y := 0; y < len(tileMap[0]); y++ {
			if !tileMap[x][y].outside && !tileMap[x][y].partOfLoop {
				// possible internal, see how many loop parts between it and the outside
				loopCount := 0
				horizontalValid := []string{"|", "L", "J", "7", "F"}
				if sValue != "leftleft" && sValue != "rightright" {
					horizontalValid = append(horizontalValid, "S")
				}
				ups := 0.0
				downs := 0.0
				for tx := 0; tx < x; tx++ {
					if tileMap[tx][y].partOfLoop && slices.Contains(horizontalValid, tileMap[tx][y].char) {
						if x == 9 && y == 9 {
							fmt.Println(*tileMap[tx][y])
						}
						if tileMap[tx][y].inputDir == "up" || tileMap[tx][y].nextInputDir == "up" {
							if tileMap[tx][y].char == "|" {
								ups++
							} else {
								ups += 0.5
							}

						}
						if tileMap[tx][y].inputDir == "down" || tileMap[tx][y].nextInputDir == "down" {
							if tileMap[tx][y].char == "|" {
								downs++
							} else {
								downs += 0.5
							}
						}
					}
				}
				max := max(ups, downs)
				min := min(ups, downs)
				loopCount = int(max) - int(min)
				if loopCount%2 == 0 {
					tileMap[x][y].outside = true
				}
			}

			//if !tileMap[x][y].outside && !tileMap[x][y].partOfLoop {
			//	// possible internal, see how many loop parts between it and the outside
			//	loopCount := 0
			//	verticalValid := []string{"-", "L", "J", "7", "F"}
			//	if sValue != "upup" && sValue != "downdown" {
			//		verticalValid = append(verticalValid, "S")
			//	}
			//	currentDir := ""
			//	for ty := 0; ty < y; ty++ {
			//		if slices.Contains(verticalValid, tileMap[x][ty].char) {
			//			if tileMap[x][ty].inputDir != currentDir && (tileMap[x][ty].inputDir == "left" || tileMap[x][ty].inputDir == "right") {
			//				loopCount++
			//				currentDir = tileMap[x][ty].inputDir
			//			}
			//		}
			//	}
			//	if loopCount%2 == 0 {
			//		tileMap[x][y].outside = true
			//	}
			//}
		}
	}

	for y := 0; y < len(tileMap[0]); y++ {
		for x := 0; x < len(tileMap); x++ {
			t := tileMap[x][y]
			if t.outside {
				//fmt.Printf("O")
			} else if t.partOfLoop {
				//fmt.Printf(t.char)
			} else {
				total++
				//fmt.Printf("X")
			}
		}
		//fmt.Printf("\n")
	}

	fmt.Println("Total:", total)
	//fmt.Println("S value:", sValue)
}

func markNeighboursOutside(x int, y int, tileMap [][]*tile) {
	for tx := x - 1; tx <= x+1; tx++ {
		for ty := y - 1; ty <= y+1; ty++ {
			if tx < 0 || tx >= len(tileMap) || ty < 0 || ty >= len(tileMap[0]) {
				continue
			}
			if !tileMap[tx][ty].partOfLoop && !tileMap[tx][ty].outside {
				tileMap[tx][ty].outside = true
				markNeighboursOutside(tx, ty, tileMap)
			}
		}
	}
}

func getDirection(char string, inputDir string) (string, bool) {

	if char == "." {
		return "", false
	}
	switch inputDir {
	case "up":
		{
			if char == "|" {
				return "up", true
			}
			if char == "L" {
				return "left", true
			}
			if char == "J" {
				return "right", true
			}
		}
	case "down":
		{
			if char == "|" {
				return "down", true
			}
			if char == "7" {
				return "right", true
			}
			if char == "F" {
				return "left", true
			}
		}
	case "left":
		{
			if char == "-" {
				return "left", true
			}
			if char == "J" {
				return "down", true
			}
			if char == "7" {
				return "up", true
			}
		}
	case "right":
		{
			if char == "-" {
				return "right", true
			}
			if char == "L" {
				return "down", true
			}
			if char == "F" {
				return "up", true
			}
		}
	}
	return "", false
}
