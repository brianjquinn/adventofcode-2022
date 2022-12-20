package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func RopeBridgePart1() {
	fmt.Println("Day 9 Part 1: Rope Bridge")

	headMoves := utils.ReadFileLinesToStringSlice("day09/series-of-motions.txt")

	headCurrRow := 0
	headCurrCol := 0
	tailCurrRow := 0
	tailCurrCol := 0

	trackMap := make(map[string]bool)
	directionMap := map[string]int{"R": 1, "L": -1, "U": -1, "D": 1}

	for _, move := range headMoves {
		moveSplit := strings.Split(move, " ")
		direction := moveSplit[0]
		num, _ := strconv.Atoi(moveSplit[1])

		horizontal := direction == "R" || direction == "L"

		for i := 0; i < num; i++ {
			if horizontal {
				headCurrCol += directionMap[direction]
				if utils.Abs(int64(headCurrCol)-int64(tailCurrCol)) == 2 {
					tailCurrCol += directionMap[direction]
					if headCurrRow != tailCurrRow {
						tailCurrRow = headCurrRow
					}
					trackMap[fmt.Sprintf("%d,%d", tailCurrRow, tailCurrCol)] = true
				}

			} else {
				headCurrRow += directionMap[direction]
				if utils.Abs(int64(headCurrRow)-int64(tailCurrRow)) == 2 {
					tailCurrRow += directionMap[direction]
					if headCurrCol != tailCurrCol {
						tailCurrCol = headCurrCol
					}
					trackMap[fmt.Sprintf("%d,%d", tailCurrRow, tailCurrCol)] = true
				}
			}
		}
	}

	fmt.Printf("The tail visited %d positions at least once\n\n", len(trackMap)+1)
}
