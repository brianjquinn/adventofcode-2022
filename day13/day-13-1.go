package day13

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func DistressSignalPart1() {
	fmt.Println("Day 13 Part 1: Distress Signal")
	packetLines := utils.ReadFileLinesToStringSlice("day13/packets-ex.txt")
	packets := parsePackets(packetLines)

	sum := 0

	for i := 0; i < len(packets)-2; i += 2 {
		order := compare(packets[i], packets[i+1])
		if order <= 0 {
			sum += ((i / 2) + 1)
		}
	}

	fmt.Printf("The sum of the indices of the pairs who are in the correct order is %d\n\n", sum)
}

func parsePackets(packetLines []string) []any {
	packets := make([]any, 0)
	for i := 0; i < len(packetLines); i++ {
		packetLine := packetLines[i]
		if strings.HasPrefix(packetLine, "[") {
			var pkt any
			json.Unmarshal([]byte(packetLine), &pkt)
			fmt.Println(pkt)
			packets = append(packets, pkt)
		}
	}
	return packets
}

func compare(left any, right any) int {
	lefts, lok := left.([]any)
	rights, rok := right.([]any)

	switch {
	case !lok && !rok:
		return int(left.(float64) - right.(float64))
	case !lok:
		lefts = []any{left}
	case !rok:
		rights = []any{right}
	}

	for i := 0; i < len(lefts) && i < len(rights); i++ {
		c := compare(lefts[i], rights[i])
		if c != 0 {
			return c
		}
	}
	return len(lefts) - len(rights)
}
