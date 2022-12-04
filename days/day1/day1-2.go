package day1

import (
	"fmt"
	"sort"
	"strconv"
)

func Top3MostCaloricElves() {
	fmt.Println("Day 1 Part 2: Top 3 Most Caloric Elves")
	items := getItems("days/day1/calories-by-elf.txt")

	var caloricTotals []int
	var currTotal int = 0
	for _, item := range items {
		if item != "" {
			var convertedCalories, err = strconv.Atoi(item)
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

	fmt.Printf("The top 3 most caloric elves are carrying %d total calories\n", top3Sum)
}
