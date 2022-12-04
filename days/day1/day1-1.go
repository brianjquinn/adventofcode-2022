package day1

import (
	"fmt"
	"strconv"
)

func MostCaloricElf() {
	fmt.Println("Day 1 Part 1: Calorie Counting")
	items := getItems("days/day1/calories-by-elf.txt")

	var maxCalories int = -1
	var currCalories int = 0

	for _, item := range items {
		if item != "" {
			var convertedCalories, err = strconv.Atoi(item)
			if err == nil {
				currCalories += convertedCalories
			}
		} else {
			if currCalories > maxCalories {
				maxCalories = currCalories
			}
			currCalories = 0
		}
	}

	fmt.Printf("The most caloric elf is carrying %d calories\n\n", maxCalories)
}
