package day07

import (
	"fmt"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func NoSpaceLeftOnDevicePart1() {
	fmt.Println("Day 7 Part 1: No Space Left On Device")

	terminalOutput := utils.ReadFileLinesToStringSlice("day07/terminal-output.txt")

	var filesystem *Filesystem = buildFileSystem(terminalOutput)
	filesystem.calculateAllDirectorySizes()
	sum := filesystem.sumOfSizeOfDirsWithSizeLessThan(100000)
	fmt.Printf("The sum of the sizes of the directories whose size is less than or equal to 100,000 is %d\n\n", sum)
}
