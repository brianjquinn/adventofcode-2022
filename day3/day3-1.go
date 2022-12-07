package day3

import (
	"fmt"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func RucksackReorganizationPart1() {
	fmt.Println("Day 3 Part 1: Rucksack Reorganization")

	rucksacks := utils.ReadFileLinesToStringSlice("day3/rucksacks.txt")

	var sum int = 0

	for _, rucksackContent := range rucksacks {
		firstCompartmentContents := []rune(rucksackContent[:len(rucksackContent)/2])
		secondCompartmentContents := []rune(rucksackContent[(len(rucksackContent) / 2):])
		commonContent := findFirstIntersection(firstCompartmentContents, secondCompartmentContents)

		if commonContent >= 97 {
			sum += (int(commonContent) - 96)
		} else {
			sum += (int(commonContent) - 38)
		}
	}

	fmt.Printf("The sum of the priorities for the common items between compartments in all rucksacks is %d\n\n", sum)
}

func findFirstIntersection(s1 []rune, s2 []rune) rune {
	s1map := make(map[string]bool)

	for i := 0; i < len(s1); i++ {
		s1map[string(s1[i])] = true
	}

	for j := 0; j < len(s2); j++ {
		if s1map[string(s2[j])] {
			return s2[j]
		}
	}

	return ' '
}
