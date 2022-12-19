package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func CathodeRayTubePart1() {
	fmt.Println("Day 10 Part 1: Cathode-Ray Tube")

	instructs := utils.ReadFileLinesToStringSlice("day10/cpu-instructions.txt")
	var instructionIdx int = 0
	var registerX int = 1
	var cycleCountsOfInterest []int = []int{20, 60, 100, 140, 180, 220}
	var cycleCountIdx int = 0
	var signalStrengthSum int = 0
	var addToSignalSumNextCycle bool = false

	for cycle := 1; cycle <= 220; cycle++ {
		if cycle == cycleCountsOfInterest[cycleCountIdx] {
			addToSigSum := cycle * registerX
			signalStrengthSum += addToSigSum
			cycleCountIdx++
		}

		instruction := instructs[instructionIdx]
		instructionSplit := strings.Split(instruction, " ")
		action := instructionSplit[0]
		if action == "noop" {
			instructionIdx++
		} else if action == "addx" && !addToSignalSumNextCycle {
			addToSignalSumNextCycle = true
		} else if action == "addx" && addToSignalSumNextCycle {
			amt, _ := strconv.Atoi(instructionSplit[1])
			registerX += amt
			instructionIdx++
			addToSignalSumNextCycle = false
		}
	}

	fmt.Printf("The sum of the signal strengths during the 20th, 60th, 100th, 140th, 180th, 220th is %d\n\n", signalStrengthSum)
}
