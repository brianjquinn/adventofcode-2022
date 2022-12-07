package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func CampCleanupPart1() {
	fmt.Println("Day 4 Part 1: Camp Cleanup")

	sectionAssignments := utils.ReadFileLinesToStringSlice("day4/section-assignments.txt")

	var pairsThatHaveARangeFullyContainingAnother int = 0

	for _, pair := range sectionAssignments {
		assignments := strings.Split(pair, ",")

		assig1StartEnd := strings.Split(assignments[0], "-")
		assig1Start, _ := strconv.Atoi(assig1StartEnd[0])
		assig1End, _ := strconv.Atoi(assig1StartEnd[1])
		assig2StartEnd := strings.Split(assignments[1], "-")
		assig2Start, _ := strconv.Atoi(assig2StartEnd[0])
		assig2End, _ := strconv.Atoi(assig2StartEnd[1])

		if assig1Start <= assig2Start && assig1End >= assig2End || assig2Start <= assig1Start && assig2End >= assig1End {
			pairsThatHaveARangeFullyContainingAnother++
		}

	}

	fmt.Printf("%d assignment pairs have one range that fully contains another\n\n", pairsThatHaveARangeFullyContainingAnother)
}
