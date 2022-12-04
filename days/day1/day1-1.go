package day1

import (
	"fmt"
	"strconv"

	utils "github.com/brianjquinn/adventofcode/days"
)

func CalorieCountingPart1() {
	fmt.Println("Day 1 Part 1: Calorie Counting")
	caloriesByElf := utils.ReadFileLinesToStringSlice("days/day1/calories-by-elf.txt")

	var maxCalories int = -1
	var currCalorieCount int = 0

	for _, calorieCount := range caloriesByElf {
		if calorieCount != "" {
			var convertedCalories, err = strconv.Atoi(calorieCount)
			if err == nil {
				currCalorieCount += convertedCalories
			}
		} else {
			if currCalorieCount > maxCalories {
				maxCalories = currCalorieCount
			}
			currCalorieCount = 0
		}
	}

	fmt.Printf("The most caloric elf is carrying %d calories\n\n", maxCalories)
}
