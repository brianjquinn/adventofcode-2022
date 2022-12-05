package day3

import (
	"errors"
	"fmt"
	"log"
	"strings"

	utils "github.com/brianjquinn/adventofcode/days"
)

const GroupSize int = 3

func RucksackReorganizationPart2() {
	fmt.Println("Day 3 Part 2: Rucksack Reorganization")

	rucksacks := utils.ReadFileLinesToStringSlice("days/day3/rucksacks.txt")
	var sum int = 0
	for i := 0; i < len(rucksacks)-2; i += GroupSize {
		group := rucksacks[i : i+GroupSize]
		badge, err := determineBadgeForGroup(group)

		if err != nil {
			log.Fatal(err)
		}

		if badge >= 97 {
			sum += (int(badge) - 96)
		} else {
			sum += (int(badge) - 38)
		}
	}

	fmt.Printf("The sum of the priorities of the badges for each %d-elf group is %d\n\n", GroupSize, sum)
}

func determineBadgeForGroup(groupRucksacks []string) (rune, error) {
	firstRucksack := []rune(groupRucksacks[0])
	secondRucksack := []rune(groupRucksacks[1])

	intersectionMap := make(map[rune]bool)

	for i := 0; i < len(firstRucksack); i++ {
		intersectionMap[firstRucksack[i]] = true
	}

	commonItems := make(map[rune]bool)
	for i := 0; i < len(secondRucksack); i++ {
		if intersectionMap[secondRucksack[i]] {
			commonItems[secondRucksack[i]] = true
		}
	}

	// search for the common items in each remaining rucksack
	// and return the first one that is found in all remaining rucksacks
	for commonItem := range commonItems {
		remainderCount := len(groupRucksacks) - 2
		foundInRemainderCount := 0
		for i := 2; i < len(groupRucksacks); i++ {
			if strings.ContainsRune(groupRucksacks[i], commonItem) {
				foundInRemainderCount++
			}
		}

		if foundInRemainderCount == remainderCount {
			return commonItem, nil
		}
	}

	return ' ', errors.New("no badge found")
}
