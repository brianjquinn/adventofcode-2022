package day11

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

type Monkey struct {
	items       []uint64
	operation   string
	divisibleBy uint64
	targets     [2]int
	maxWorry    uint64
}

func (m Monkey) applyOperationToItem(item uint64) (uint64, error) {
	rawOp := strings.TrimPrefix(m.operation, "new = ")
	parsedOp := strings.Split(rawOp, " ")
	firstOperand, err := strconv.ParseUint(parsedOp[0], 10, 64)
	if err != nil {
		firstOperand = item
	}
	op := parsedOp[1]
	secondOperand, err := strconv.ParseUint(parsedOp[2], 10, 64)
	if err != nil {
		secondOperand = item
	}

	if op == "*" {
		return firstOperand * secondOperand, nil
	} else if op == "+" {
		return firstOperand + secondOperand, nil
	}

	return 0, fmt.Errorf("uknown operation: %s", rawOp)
}

func (m *Monkey) inspectAndThrowItems(monkeys []*Monkey, relaxFactor uint64) int {
	numItemsInspected := len(m.items)
	for len(m.items) > 0 {
		item := m.items[0]
		if len(m.items) == 1 {
			m.items = []uint64{}
		} else {
			m.items = m.items[1:]
		}
		worryLevel, err := m.applyOperationToItem(item)
		worryLevel /= relaxFactor
		if err != nil {
			log.Fatal("bad operation")
		}
		var targetMonkeyIdx int
		if worryLevel%m.divisibleBy == 0 {
			targetMonkeyIdx = m.targets[0]
		} else {
			targetMonkeyIdx = m.targets[1]
		}

		monkeys[targetMonkeyIdx].items = append(monkeys[targetMonkeyIdx].items, worryLevel)
	}
	return numItemsInspected
}

func (m *Monkey) inspectAndThrowItemsPart2(monkeys []*Monkey) uint64 {
	numItemsInspected := uint64(len(m.items))
	for len(m.items) > 0 {
		item := m.items[0]
		if len(m.items) == 1 {
			m.items = []uint64{}
		} else {
			m.items = m.items[1:]
		}
		worryLevel, err := m.applyOperationToItem(item)
		worryLevel %= m.maxWorry
		if err != nil {
			log.Fatal("bad operation")
		}
		var targetMonkeyIdx int
		if worryLevel%m.divisibleBy == 0 {
			targetMonkeyIdx = m.targets[0]
		} else {
			targetMonkeyIdx = m.targets[1]
		}

		monkeys[targetMonkeyIdx].items = append(monkeys[targetMonkeyIdx].items, worryLevel)
	}
	return numItemsInspected
}

func createMonkeys(monkeyNotes []string) []*Monkey {
	monkeys := make([]*Monkey, 0)
	for i := 0; i < len(monkeyNotes); i++ {
		noteLine := monkeyNotes[i]
		if strings.HasPrefix(noteLine, "Monkey") {
			items := utils.MapStringArrToUInt64Arr(strings.Split(strings.TrimPrefix(monkeyNotes[i+1], "  Starting items: "), ", "))
			operation := strings.TrimPrefix(monkeyNotes[i+2], "  Operation: ")
			divisibleBy, _ := strconv.ParseUint(strings.TrimPrefix(monkeyNotes[i+3], "  Test: divisible by "), 10, 64)
			trueDestination, _ := strconv.Atoi(strings.TrimPrefix(monkeyNotes[i+4], "    If true: throw to monkey "))
			falseDestination, _ := strconv.Atoi(strings.TrimPrefix(monkeyNotes[i+5], "    If false: throw to monkey "))
			i += 5
			monkeys = append(monkeys, &Monkey{
				items:       items,
				operation:   operation,
				divisibleBy: divisibleBy,
				targets:     [2]int{trueDestination, falseDestination},
			})
		}
	}

	var maxWorry uint64 = 1
	for _, monkey := range monkeys {
		maxWorry *= monkey.divisibleBy
	}

	for _, monkey := range monkeys {
		monkey.maxWorry = maxWorry
	}
	return monkeys
}
