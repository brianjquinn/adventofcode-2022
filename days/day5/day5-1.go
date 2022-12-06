package day5

import (
	"fmt"
	"regexp"
	"strconv"

	utils "github.com/brianjquinn/adventofcode/days"
)

func SupplyStacksPart1() {
	fmt.Println("Day 5 Part 1: Supply Stacks")

	stackAndProcedure := utils.ReadFileLinesToStringSlice("days/day5/rearrangement-procedure.txt")

	var gapIdx = 0
	for stackAndProcedure[gapIdx] != "" {
		gapIdx++
	}

	var initialStackState []string = stackAndProcedure[:gapIdx]
	var procedure []string = stackAndProcedure[gapIdx+1:]

	var stacks []Stack[string] = buildStacks(initialStackState)
	executeProcedure(&stacks, procedure)

	var crates string
	for _, stack := range stacks {
		crates += stack.Pop()
	}

	fmt.Printf("Crates at the top of stacks are \"%s\"\n\n", crates)
}

func buildStacks(initialStackState []string) []Stack[string] {
	stackIds := initialStackState[len(initialStackState)-1]
	numStacks, _ := strconv.Atoi(string(stackIds[len(stackIds)-2]))
	var stacks []Stack[string] = make([]Stack[string], numStacks)

	// start at the line in the input that represents the "bottom"
	// of all of the stacks
	for i := len(initialStackState) - 2; i >= 0; i-- {
		stackStateLine := initialStackState[i]
		// iterate over every 4th index of the string (each stack takes 3 chars, plus a space in between each stack)
		for j := 1; j < len(stackStateLine); j += 4 {
			stackEntry := string(stackStateLine[j])
			if stackEntry != " " {
				stacks[j/4].Push(stackEntry)
			}
		}
	}

	return stacks
}

func executeProcedure(stacks *[]Stack[string], procedure []string) {
	for _, action := range procedure {
		re := regexp.MustCompile(`(\d+)`)
		matches := re.FindAll([]byte(action), -1)
		quantity, _ := strconv.Atoi(string(matches[0]))
		from, _ := strconv.Atoi(string(matches[1]))
		to, _ := strconv.Atoi(string(matches[2]))

		for i := 0; i < quantity; i++ {
			(*stacks)[to-1].Push((*stacks)[from-1].Pop())
		}
	}
}
