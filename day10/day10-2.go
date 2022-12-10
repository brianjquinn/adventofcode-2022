package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func CathodeRayTubePart2() {
	fmt.Println("Day 10 Part 2: Cathode-Ray Tube")

	instructs := utils.ReadFileLinesToStringSlice("day10/cpu-instructions.txt")
	var CRT [6][40]string = [6][40]string{}
	var instructionIdx int = 0
	var executeAddNextCycle bool = false
	var spritePositionCenter int = 1

	for cycle := 0; cycle < 240; cycle++ {
		currentlyDrawingRow := cycle / 40
		currentlyDrawingPosition := cycle % 40

		if spritePositionCenter >= currentlyDrawingPosition-1 && spritePositionCenter <= currentlyDrawingPosition+1 {
			CRT[currentlyDrawingRow][currentlyDrawingPosition] = "#"
		} else {
			CRT[currentlyDrawingRow][currentlyDrawingPosition] = " "
		}

		instruction := instructs[instructionIdx]
		instructionSplit := strings.Split(instruction, " ")
		action := instructionSplit[0]
		if action == "noop" {
			instructionIdx++
		} else if action == "addx" && !executeAddNextCycle {
			executeAddNextCycle = true
		} else if action == "addx" && executeAddNextCycle {
			amt, _ := strconv.Atoi(instructionSplit[1])
			spritePositionCenter += amt
			instructionIdx++
			executeAddNextCycle = false
		}
	}

	for i := 0; i < len(CRT); i++ {
		fmt.Println(CRT[i])
	}
}
