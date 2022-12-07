package day2

import (
	"fmt"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func RockPaperScissorsPart2() {
	fmt.Println("Day 2 Part 2: Rock Paper Scissors")

	rpsRounds := utils.ReadFileLinesToStringSlice("day2/strategy-guide.txt")

	var totalScore int = 0
	for _, rpsRound := range rpsRounds {
		score := scoreRoundPart2(rpsRound)
		totalScore += score
	}

	fmt.Printf("Based on the input strategy guide, the score I would receive is: %d\n\n", totalScore)
}

func scoreRoundPart2(round string) int {
	desiredOutcomeToScore := map[string]int{"X": 0, "Y": 3, "Z": 6}
	selectionSplit := strings.Split(round, " ")
	opponentSelection := selectionSplit[0]
	desiredRoundOutcome := selectionSplit[1]

	return desiredOutcomeToScore[desiredRoundOutcome] + scoreBasedOnWhatMySelectionShouldBe(opponentSelection, desiredRoundOutcome)
}

func scoreBasedOnWhatMySelectionShouldBe(oppSelection string, desiredOutcome string) int {

	mySelectionToScore := map[string]int{"X": 1, "Y": 2, "Z": 3}
	drawSelections := map[string]string{"A": "X", "B": "Y", "C": "Z"}
	winningSelections := map[string]string{"A": "Y", "B": "Z", "C": "X"}
	losingSelections := map[string]string{"A": "Z", "B": "X", "C": "Y"}
	if desiredOutcome == "X" {
		return mySelectionToScore[losingSelections[oppSelection]]
	} else if desiredOutcome == "Y" {
		return mySelectionToScore[drawSelections[oppSelection]]
	}
	return mySelectionToScore[winningSelections[oppSelection]]
}
