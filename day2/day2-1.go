package day2

import (
	"fmt"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func RockPaperScissorsPart1() {
	fmt.Println("Day 2 Part 1: Rock Paper Scissors")

	rpsRounds := utils.ReadFileLinesToStringSlice("day2/strategy-guide.txt")

	var totalScore int = 0
	for _, rpsRound := range rpsRounds {
		totalScore += scoreRoundPart1(rpsRound)
	}

	fmt.Printf("Based on the input strategy guide, the score I would receive is: %d\n\n", totalScore)
}

func scoreRoundPart1(round string) int {
	mySelectionToScore := map[string]int{"X": 1, "Y": 2, "Z": 3}
	selectionSplit := strings.Split(round, " ")
	opponentSelection := selectionSplit[0]
	mySelection := selectionSplit[1]

	return mySelectionToScore[mySelection] + evaluateWinnerAndScore(opponentSelection, mySelection)
}

func evaluateWinnerAndScore(oppSelection string, mySelection string) int {
	draw := map[string]string{"A": "X", "B": "Y", "C": "Z"}
	oppSelectionToLoser := map[string]string{"A": "Y", "B": "Z", "C": "X"}
	if draw[oppSelection] == mySelection {
		return 3
	} else if mySelection == oppSelectionToLoser[oppSelection] {
		return 6
	}
	return 0
}
