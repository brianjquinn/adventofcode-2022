package day6

import (
	"fmt"

	utils "github.com/brianjquinn/adventofcode/days"
)

func TuningTroublePart2() {
	fmt.Println("Day 6 Part 2: Tuning Trouble")

	datastreamBuffer := utils.ReadFileLinesToStringSlice("days/day6/datastream-buffer.txt")[0]

	charactersProcessed, _ := FindEndIdxOfMarkerWithUniqLength(14, datastreamBuffer)

	fmt.Printf("%d characters need to be processed before a start-of-packet marker is detected\n\n", charactersProcessed)
}
