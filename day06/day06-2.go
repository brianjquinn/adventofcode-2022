package day06

import (
	"fmt"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func TuningTroublePart2() {
	fmt.Println("Day 6 Part 2: Tuning Trouble")

	datastreamBuffer := utils.ReadFileLinesToStringSlice("day06/datastream-buffer.txt")[0]

	charactersProcessed, _ := FindEndIdxOfMarkerWithUniqLength(14, datastreamBuffer)

	fmt.Printf("%d characters need to be processed before a start-of-packet marker is detected\n\n", charactersProcessed)
}
