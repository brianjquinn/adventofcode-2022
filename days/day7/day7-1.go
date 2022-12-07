package day7

import (
	"fmt"

	utils "github.com/brianjquinn/adventofcode/days"
)

func NoSpaceLeftOnDevicePart1() {
	fmt.Println("Day 7 Part 1: No Space Left On Device")

	terminalOutput := utils.ReadFileLinesToStringSlice("days/day7/terminal-output.txt")

	var root *Directory = buildFileSystem(terminalOutput)
	root.calcSize()
	var solution *int = new(int)
	sumDirsWithSizeGreaterThan100k(root, solution)

	fmt.Printf("The sum of the sizes of the dirctories whose size is less than or equal to 100,000 is %d\n\n", *solution)
}

func sumDirsWithSizeGreaterThan100k(dir *Directory, sum *int) {
	if dir.size <= 100000 {
		*sum += dir.size
	}
	for _, directory := range dir.directories {
		sumDirsWithSizeGreaterThan100k(directory, sum)
	}
}
