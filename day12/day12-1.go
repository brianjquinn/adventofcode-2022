package day12

import (
	"fmt"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func HillClimbingAlgorithmPart1() {
	fmt.Println("Day 12 Part 1: Hill Climbing Algorithm")

	terrainMapLines := utils.ReadFileLinesToStringSlice("day12/terrain-map.txt")

	terrainMap := make([][]rune, 0)

	for _, terrainLine := range terrainMapLines {
		terrainMapRow := make([]rune, 0)
		for _, square := range terrainLine {
			terrainMapRow = append(terrainMapRow, square)
		}
		terrainMap = append(terrainMap, terrainMapRow)
	}

	var shortestPath int = -1
	for i := 0; i < len(terrainMap); i++ {
		for j := 0; j < len(terrainMap[i]); j++ {
			if terrainMap[i][j] == 'S' {
				// we found the start - begin the ascent to E
				shortestPath = ascend(i, j, terrainMap)
			}
		}
		if shortestPath > -1 {
			break
		}
	}

	fmt.Printf("The shortest path from the start to the location with the best signal is %d", shortestPath)
}
