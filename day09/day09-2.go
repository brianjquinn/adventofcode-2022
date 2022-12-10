package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func RopeBridgePart2() {
	fmt.Println("Day 9 Part 2: Rope Bridge")

	headMoves := utils.ReadFileLinesToStringSlice("day09/series-of-motions.txt")

	directionMap := map[string]int{"R": 1, "L": -1, "U": -1, "D": 1}
	trackMap := make(map[string]bool)

	// all knot initial positions are 0,0
	knots := [10][2]int{}
	// knots[0] = head knot
	// knots[len(knots) - 1] = tail knot
	headIdx := 0
	tailIdx := len(knots) - 1
	for _, move := range headMoves {
		moveSplit := strings.Split(move, " ")
		direction := moveSplit[0]
		num, _ := strconv.Atoi(moveSplit[1])

		horizontal := direction == "R" || direction == "L"

		for i := 0; i < num; i++ {
			if horizontal {
				knots[headIdx][1] += directionMap[direction]
			} else {
				knots[0][headIdx] += directionMap[direction]
			}
			// adjust the rest of the knots
			lastKnot := knots[headIdx]
			for j := 1; j < len(knots); j++ {
				currKnot := knots[j]
				currKnotNewRow, currKnotNewCol := calcNewKnotPosition(currKnot, lastKnot)
				if currKnotNewRow != currKnot[0] || currKnotNewCol != currKnot[1] {
					knots[j] = [...]int{currKnotNewRow, currKnotNewCol}
					if j == tailIdx {
						trackMap[fmt.Sprintf("%d,%d", knots[j][0], knots[j][1])] = true
					}
				}
				lastKnot = knots[j]
			}
		}
	}

	fmt.Printf("The tail of the 10 knot rope visited %d positions at least once\n\n", len(trackMap)+1)
}

func calcNewKnotPosition(currKnot [2]int, lastKnot [2]int) (int, int) {
	// above = -1, below = +1 row (0 index of the knot)
	// left = -1, right = +1 col (1st index of the knot)
	currKnotRow := currKnot[0]
	currKnotCol := currKnot[1]

	lastKnotRow := lastKnot[0]
	lastKnotCol := lastKnot[1]

	// direct right
	if lastKnotRow == currKnotRow && lastKnotCol-currKnotCol == 2 {
		return currKnotRow, currKnotCol + 1
	}

	// direct left
	if lastKnotRow == currKnotRow && lastKnotCol-currKnotCol == -2 {
		return currKnotRow, currKnotCol - 1
	}

	// direct above
	if lastKnotCol == currKnotCol && lastKnotRow-currKnotRow == -2 {
		return currKnotRow - 1, currKnotCol
	}

	// direct below
	if lastKnotCol == currKnotCol && lastKnotRow-currKnotRow == 2 {
		return currKnotRow + 1, currKnotCol
	}

	// left up
	if lastKnotCol-currKnotCol <= -1 && lastKnotRow-currKnotRow == -2 ||
		lastKnotCol-currKnotCol == -2 && lastKnotRow-currKnotRow <= -1 {
		return currKnotRow - 1, currKnotCol - 1
	}

	// right up
	if lastKnotCol-currKnotCol >= 1 && lastKnotRow-currKnotRow == -2 ||
		lastKnotCol-currKnotCol == 2 && lastKnotRow-currKnotRow <= -1 {
		return currKnotRow - 1, currKnotCol + 1
	}

	// right down
	if lastKnotCol-currKnotCol >= 1 && lastKnotRow-currKnotRow == 2 ||
		lastKnotCol-currKnotCol == 2 && lastKnotRow-currKnotRow >= 1 {
		return currKnotRow + 1, currKnotCol + 1
	}

	// left down
	if lastKnotCol-currKnotCol >= -1 && lastKnotRow-currKnotRow == 2 ||
		lastKnotCol-currKnotCol == -2 && lastKnotRow-currKnotRow >= 1 {
		return currKnotRow + 1, currKnotCol - 1
	}

	return currKnotRow, currKnotCol
}
