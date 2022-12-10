package day08

import (
	"fmt"
	"strconv"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func TreetopTreeHousePart2() {
	fmt.Println("Day 8 Part 2: Treetop Tree House")

	treeGridLines := utils.ReadFileLinesToStringSlice("day08/tree-height-map.txt")

	treeGrid := make([][]int, 0)
	for _, gridLine := range treeGridLines {
		rowArr := make([]int, 0)
		for _, treeHeight := range gridLine {
			intTreeHeight, _ := strconv.Atoi(string(treeHeight))
			rowArr = append(rowArr, intTreeHeight)
		}
		treeGrid = append(treeGrid, rowArr)
	}

	var maxScenicScore = -1
	for row := 1; row < len(treeGrid)-1; row++ {
		for col := 1; col < len(treeGrid[row])-1; col++ {
			var scenicScore int = calcScenicScore(row, col, treeGrid)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Printf("The maximum scenic score is %d\n\n", maxScenicScore)
}

func calcScenicScore(row int, col int, treeGrid [][]int) int {
	var currTree int = treeGrid[row][col]

	top := row - 1
	for currTree > treeGrid[top][col] && top > 0 {
		top--
	}

	numTreesTop := row - top

	bottom := row + 1
	for currTree > treeGrid[bottom][col] && bottom < len(treeGrid)-1 {
		bottom++
	}

	numTreesBottom := bottom - row

	left := col - 1
	for currTree > treeGrid[row][left] && left > 0 {
		left--
	}

	numTreesLeft := col - left

	right := col + 1
	for currTree > treeGrid[row][right] && right < len(treeGrid[row])-1 {
		right++
	}

	numTreesRight := right - col

	return numTreesTop * numTreesBottom * numTreesLeft * numTreesRight
}
