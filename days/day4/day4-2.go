package day4

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/brianjquinn/adventofcode/days"
)

func CampCleanupPart2() {
	fmt.Println("Day 4 Part 1: Camp Cleanup")

	sectionAssignments := utils.ReadFileLinesToStringSlice("days/day4/section-assignments.txt")

	var pairsWithOverlappingRanges int = 0

	for _, pair := range sectionAssignments {
		assignments := strings.Split(pair, ",")

		assig1StartEnd := strings.Split(assignments[0], "-")
		assig1Start, _ := strconv.Atoi(assig1StartEnd[0])
		assig1End, _ := strconv.Atoi(assig1StartEnd[1])
		assig2StartEnd := strings.Split(assignments[1], "-")
		assig2Start, _ := strconv.Atoi(assig2StartEnd[0])
		assig2End, _ := strconv.Atoi(assig2StartEnd[1])

		if assig1Start >= assig2Start && assig1Start <= assig2End || assig1Start <= assig2Start && assig1End >= assig2Start {
			pairsWithOverlappingRanges++
		}

	}

	fmt.Printf("%d assignment pairs have one range that overlaps another\n\n", pairsWithOverlappingRanges)
}
