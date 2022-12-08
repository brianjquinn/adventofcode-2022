package day8

import (
	"fmt"
	"strconv"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func TreetopTreeHousePart1() {
	fmt.Println("Day 8 Part 1: Treetop Tree House")

	treeGridLines := utils.ReadFileLinesToStringSlice("day8/tree-height-map.txt")

	treeGrid := make([][]int, 0)
	for _, gridLine := range treeGridLines {
		rowArr := make([]int, 0)
		for _, treeHeight := range gridLine {
			intTreeHeight, _ := strconv.Atoi(string(treeHeight))
			rowArr = append(rowArr, intTreeHeight)
		}
		treeGrid = append(treeGrid, rowArr)
	}

	// - 4 to remove the corners which get duplicated
	visibleTrees := len(treeGrid)*2 + len(treeGrid[0])*2 - 4

	for row := 1; row < len(treeGrid)-1; row++ {
		for col := 1; col < len(treeGrid[row])-1; col++ {
			if visible(row, col, treeGrid) {
				visibleTrees++
			}
		}
	}

	fmt.Printf("The number of visible trees are %d\n\n", visibleTrees)
}

func visible(row int, col int, treeGrid [][]int) bool {
	var currTree int = treeGrid[row][col]

	top := row - 1
	for currTree > treeGrid[top][col] && top > 0 {
		top--
	}
	if top == 0 && currTree > treeGrid[top][col] {
		return true
	}

	bottom := row + 1
	for currTree > treeGrid[bottom][col] && bottom < len(treeGrid)-1 {
		bottom++
	}
	if bottom == len(treeGrid)-1 && currTree > treeGrid[bottom][col] {
		return true
	}

	left := col - 1
	for currTree > treeGrid[row][left] && left > 0 {
		left--
	}

	if left == 0 && currTree > treeGrid[row][left] {
		return true
	}

	right := col + 1
	for currTree > treeGrid[row][right] && right < len(treeGrid[row])-1 {
		right++
	}

	if right == len(treeGrid[row])-1 && currTree > treeGrid[row][right] {
		return true
	}

	return false
}
