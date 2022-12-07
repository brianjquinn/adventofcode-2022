package day7

import (
	"fmt"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

const TOTAL_DISK_SPACE int = 70000000
const NEEDED_UNUSED_SPACE int = 30000000

func NoSpaceLeftOnDevicePart2() {
	fmt.Println("Day 7 Part 2: No Space Left On Device")

	terminalOutput := utils.ReadFileLinesToStringSlice("day7/terminal-output.txt")

	var filesystem *Filesystem = buildFileSystem(terminalOutput)
	filesystem.calculateAllDirectorySizes()

	minAmountNeededToFree := NEEDED_UNUSED_SPACE - filesystem.unusedSpace()

	size := filesystem.sizeOfSmallestDirectoryToDelete(minAmountNeededToFree)
	fmt.Printf("The size of the smallest directory to delete in order to free up enough space is %d\n\n", size)
}
