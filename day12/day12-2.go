package day12

import (
	"fmt"
	"math"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func HillClimbingAlgorithmPart2() {
	fmt.Println("Day 12 Part 2: Hill Climbing Algorithm")
	terrainMapLines := utils.ReadFileLinesToStringSlice("day12/terrain-map.txt")

	terrainMap := make([][]rune, 0)

	for _, terrainLine := range terrainMapLines {
		terrainMapRow := make([]rune, 0)
		for _, square := range terrainLine {
			terrainMapRow = append(terrainMapRow, square)
		}
		terrainMap = append(terrainMap, terrainMapRow)
	}

	var shortestPathStartingFromAnA int = math.MaxInt
	for i := 0; i < len(terrainMap); i++ {
		for j := 0; j < len(terrainMap[i]); j++ {
			if terrainMap[i][j] == 'a' || terrainMap[i][j] == 'S' {
				pathLenFromAnA := ascend(i, j, terrainMap)
				if pathLenFromAnA != -1 && pathLenFromAnA < shortestPathStartingFromAnA {
					shortestPathStartingFromAnA = pathLenFromAnA
				}
			}
		}
	}

	fmt.Printf("The shortest path starting from a square with elevation 'a' is %d\n\n", shortestPathStartingFromAnA)
}
