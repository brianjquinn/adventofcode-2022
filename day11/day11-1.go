package day11

import (
	"fmt"
	"sort"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func MonkeyInTheMiddlePart1() {
	fmt.Println("Day 11 Part 1: Monkey in the Middle")
	monkeyNotes := utils.ReadFileLinesToStringSlice("day11/monkey-notes.txt")
	var monkeys []*Monkey = createMonkeys(monkeyNotes)
	var inspectionCounts []int = make([]int, len(monkeys))

	for i := 0; i < 20; i++ {
		for j, monkey := range monkeys {
			inspectionCounts[j] += monkey.inspectAndThrowItems(monkeys, 3)
		}
	}

	sort.Ints(inspectionCounts)

	levelOfMonkeyBusiness := inspectionCounts[len(inspectionCounts)-1] * inspectionCounts[len(inspectionCounts)-2]

	fmt.Printf("The level of monkey business after 20 rounds of stuff-slinging simian shenanagins is %d\n\n", levelOfMonkeyBusiness)
}
