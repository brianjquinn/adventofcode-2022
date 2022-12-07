package day7

import (
	"fmt"

	utils "github.com/brianjquinn/adventofcode/days"
)

func NoSpaceLeftOnDevicePart2() {
	fmt.Println("Day 7 Part 2: No Space Left On Device")

	terminalOutput := utils.ReadFileLinesToStringSlice("days/day7/terminal-output.txt")

	var root *Directory = buildFileSystem(terminalOutput)
	root.calcSize()
}
