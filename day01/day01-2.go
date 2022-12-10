package day01

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func CalorieCountingPart2() {
	fmt.Println("Day 1 Part 2: Calorie Counting")
	caloriesByElf := utils.ReadFileLinesToStringSlice("day01/calories-by-elf.txt")

	var caloricTotals []int
	var currTotal int = 0
	for _, calorieCount := range caloriesByElf {
		if calorieCount != "" {
			var convertedCalories, err = strconv.Atoi(calorieCount)
			if err == nil {
				currTotal += convertedCalories
			}
		} else {
			caloricTotals = append(caloricTotals, currTotal)
			currTotal = 0
		}
	}

	sort.Ints(caloricTotals)

	var top3 []int = caloricTotals[len(caloricTotals)-3:]
	var top3Sum int = 0
	for _, x := range top3 {
		top3Sum += x
	}

	fmt.Printf("The top 3 most caloric elves are carrying %d total calories\n\n", top3Sum)
}
