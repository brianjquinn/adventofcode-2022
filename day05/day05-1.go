package day05

import (
	"fmt"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func SupplyStacksPart1() {
	fmt.Println("Day 5 Part 1: Supply Stacks")

	stackAndProcedure := utils.ReadFileLinesToStringSlice("day05/rearrangement-procedure.txt")

	var gapIdx = 0
	for stackAndProcedure[gapIdx] != "" {
		gapIdx++
	}

	var initialStackState []string = stackAndProcedure[:gapIdx]
	var procedure []string = stackAndProcedure[gapIdx+1:]

	var stacks []Stack[string] = buildStacksFromInput(initialStackState)
	executeProcedurePart1(&stacks, procedure)

	var crates string
	for _, stack := range stacks {
		crates += stack.Pop()
	}

	fmt.Printf("Crates at the top of stacks are \"%s\"\n\n", crates)
}
