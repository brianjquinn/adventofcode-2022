package day5

import (
	"regexp"
	"strconv"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(t T) {
	s.items = append(s.items, t)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items)
	t := s.items[n-1]
	var zero T
	s.items[n-1] = zero
	s.items = s.items[:n-1]
	return t
}

func buildStacksFromInput(initialStackState []string) []Stack[string] {
	stackIds := initialStackState[len(initialStackState)-1]
	numStacks, _ := strconv.Atoi(string(stackIds[len(stackIds)-2]))
	var stacks []Stack[string] = make([]Stack[string], numStacks)

	// start at the line in the input that represents the "bottom"
	// of all of the stacks
	for i := len(initialStackState) - 2; i >= 0; i-- {
		stackStateLine := initialStackState[i]
		// grab every 4th non-empty index of the string (each stack is 3 chars, plus a space in between each stack)
		for j := 1; j < len(stackStateLine); j += 4 {
			stackEntry := string(stackStateLine[j])
			if stackEntry != " " {
				stacks[j/4].Push(stackEntry)
			}
		}
	}

	return stacks
}

func getActionVals(action string) (int, int, int) {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAll([]byte(action), -1)
	quantity, _ := strconv.Atoi(string(matches[0]))
	from, _ := strconv.Atoi(string(matches[1]))
	to, _ := strconv.Atoi(string(matches[2]))

	return quantity, from, to
}

func executeProcedurePart1(stacks *[]Stack[string], procedure []string) {
	for _, action := range procedure {
		quantity, from, to := getActionVals(action)
		for i := 0; i < quantity; i++ {
			(*stacks)[to-1].Push((*stacks)[from-1].Pop())
		}
	}
}

func executeProcedurePart2(stacks *[]Stack[string], procedure []string) {
	for _, action := range procedure {
		quantity, from, to := getActionVals(action)

		var staging []string

		// pop crates into a staging array
		for i := 0; i < quantity; i++ {
			staging = append(staging, (*stacks)[from-1].Pop())
		}

		// iterate through the staging array in reverse - pushing each crate into
		// the destination stack
		for i := len(staging) - 1; i >= 0; i-- {
			(*stacks)[to-1].Push(staging[i])
		}
	}
}
