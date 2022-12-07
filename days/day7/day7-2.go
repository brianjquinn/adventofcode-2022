package day7

import (
	"fmt"
	"math"

	utils "github.com/brianjquinn/adventofcode/days"
)

const TOTAL_DISK_SPACE int = 70000000
const NEEDED_UNUSED_SPACE int = 30000000

func NoSpaceLeftOnDevicePart2() {
	fmt.Println("Day 7 Part 2: No Space Left On Device")

	terminalOutput := utils.ReadFileLinesToStringSlice("days/day7/terminal-output.txt")

	var root *Directory = buildFileSystem(terminalOutput)
	root.calcSize()

	currentUnusedSpace := TOTAL_DISK_SPACE - root.size
	minAmountNeededToFree := NEEDED_UNUSED_SPACE - currentUnusedSpace

	var solution *int = new(int)
	*solution = math.MaxInt
	findSmallestDirectoryToDelete(root, minAmountNeededToFree, solution)

	fmt.Printf("The size of the smallest directory to delete in order to free up enough space is %d\n\n", *solution)
}

func findSmallestDirectoryToDelete(dir *Directory, size int, smallest *int) {
	if dir.size >= size && dir.size < *smallest {
		*smallest = dir.size
	}

	for _, directory := range dir.directories {
		findSmallestDirectoryToDelete(directory, size, smallest)
	}
}
