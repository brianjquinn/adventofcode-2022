package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

const (
	Down              int = 0
	DiagonalDownLeft  int = 1
	DiagonalDownRight int = 2
)

var directionMods [][]int = [][]int{{1, 0}, {1, -1}, {1, 1}}

func RegolithReservoirPart1() {
	fmt.Println("Day 14 Part 1: Regolith Reservoir")
	// initialize the sand simulation grid with every cell = "."
	rows := 1000
	cols := 1000
	sandSimGrid := make([][]string, 0)
	for i := 0; i < rows; i++ {
		gridRow := make([]string, 0)
		for j := 0; j < cols; j++ {
			gridRow = append(gridRow, ".")
		}
		sandSimGrid = append(sandSimGrid, gridRow)
	}
	// draw the starting point
	sandSimGrid[0][500] = "+"

	rockStructures := utils.ReadFileLinesToStringSlice("day14/cave-scan.txt")
	// iterate through the rock structures and draw them into the grid
	maxRow := 0
	for _, rockStructure := range rockStructures {
		rockLocations := strings.Split(rockStructure, " -> ")
		prev := rockLocations[0]
		for i := 1; i < len(rockLocations); i++ {
			prevSplit := strings.Split(prev, ",")
			prevCol, _ := strconv.Atoi(prevSplit[0])
			prevRow, _ := strconv.Atoi(prevSplit[1])

			currSplit := strings.Split(rockLocations[i], ",")
			currCol, _ := strconv.Atoi(currSplit[0])
			currRow, _ := strconv.Atoi(currSplit[1])
			// moving up or down
			if prevCol == currCol && prevRow != currRow {
				rowDirection := 1
				if prevRow > currRow {
					rowDirection = -1
				}
				for j := prevRow; j != currRow+rowDirection; j += rowDirection {
					if j > maxRow {
						maxRow = j
					}
					sandSimGrid[j][currCol] = "#"
				}
			} else if prevRow == currRow && prevCol != currCol {
				colDirection := 1
				if prevCol > currCol {
					colDirection = -1
				}
				for j := prevCol; j != currCol+colDirection; j += colDirection {
					sandSimGrid[currRow][j] = "#"
				}
			}
			prev = rockLocations[i]
		}
	}

	// start drawing the sand on the simulation area
	grainsOfSandAtRest := drawSand(&sandSimGrid, maxRow)
	fmt.Printf("The number of grains of sand that have come to rest is %d\n\n", grainsOfSandAtRest)
}

// func printSandSimGrid(ssg [][]string) {
// 	for i := 0; i < 12; i++ {
// 		gridRow := ssg[i]
// 		for j := 475; j <= 525; j++ {
// 			fmt.Print(gridRow[j])
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }

func drawSand(simGrid *[][]string, maxRow int) int {
	currLocation := []int{0, 500}
	allAtRest := false
	sandCount := 0
	for !allAtRest {
		currAtRest := false
		for !currAtRest {
			i := 0
			var nextCellRow int
			var nextCellCol int
			for ; i < 3; i++ {
				nextCellRow = currLocation[0] + directionMods[i][0]
				nextCellCol = currLocation[1] + directionMods[i][1]
				if nextCellRow < len(*simGrid) && nextCellCol > 0 && nextCellCol < len((*simGrid)[nextCellRow]) {
					nextCell := (*simGrid)[nextCellRow][nextCellCol]
					if nextCell == "." {
						break
					}
				} else if nextCellRow > maxRow {
					currAtRest = true
					allAtRest = true
				}
			}

			if i != 3 {
				(*simGrid)[nextCellRow][nextCellCol] = "o"
				if (*simGrid)[currLocation[0]][currLocation[1]] != "+" {
					(*simGrid)[currLocation[0]][currLocation[1]] = "."
				}
				currLocation = []int{nextCellRow, nextCellCol}
			} else {
				currAtRest = true
			}
		}
		if !allAtRest {
			sandCount++
		}
		currLocation = []int{0, 500}
		currAtRest = false
	}
	return sandCount
}
