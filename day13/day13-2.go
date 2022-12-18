package day13

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func DistressSignalPart2() {
	fmt.Println("Day 13 Part 2: Distress Signal")
	packetLines := utils.ReadFileLinesToStringSlice("day13/packets.txt")
	packets := parsePacketsPart2(packetLines)
	sort.SliceStable(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})

	var decoderKey int = 0
	for i, pkt := range packets {
		jsonPktByteArr, _ := json.Marshal(pkt)
		pktStr := string(jsonPktByteArr)
		if pktStr == "[[2]]" || pktStr == "[[6]]" {
			if decoderKey == 0 {
				decoderKey = (i + 1)
			} else {
				decoderKey *= (i + 1)
			}
		}
	}
	fmt.Printf("The decoder key is %d\n\n", decoderKey)
}

func parsePacketsPart2(packetLines []string) []any {
	packets := make([]any, 0)
	for i := 0; i < len(packetLines); i++ {
		packetLine := packetLines[i]
		if strings.HasPrefix(packetLine, "[") {
			var pkt any
			json.Unmarshal([]byte(packetLine), &pkt)
			packets = append(packets, pkt)
		}
	}
	var divPkt1, divPkt2 any
	json.Unmarshal([]byte("[[2]]"), &divPkt1)
	packets = append(packets, divPkt1)
	json.Unmarshal([]byte("[[6]]"), &divPkt2)
	packets = append(packets, divPkt2)
	return packets
}
