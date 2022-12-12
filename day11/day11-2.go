package day11

import (
	"fmt"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func MonkeyInTheMiddlePart2() {
	fmt.Println("Day 11 Part 2: Monkey in the Middle")
	monkeyNotes := utils.ReadFileLinesToStringSlice("day11/monkey-notes.txt")
	var monkeys []*Monkey = createMonkeys(monkeyNotes)
	var inspectionCounts []uint64 = make([]uint64, len(monkeys))

	for i := 0; i < 10000; i++ {
		for j, monkey := range monkeys {
			inspectionCounts[j] += monkey.inspectAndThrowItemsPart2(monkeys)
		}
	}

	max := uint64(0)
	for _, count := range inspectionCounts {
		if count > max {
			max = count
		}
	}

	secondMax := uint64(0)
	for _, count := range inspectionCounts {
		if count != max && count > secondMax {
			secondMax = count
		}
	}

	levelOfMonkeyBusiness := max * secondMax

	fmt.Printf("The level of monkey business after 20 rounds of stuff-slinging simian shenanagins is %d\n\n", levelOfMonkeyBusiness)
}
