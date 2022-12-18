package day13

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

type IntOrSliceInt struct {
	arr *[]IntOrSliceInt
	v   *int
}

type Packet struct {
	left  *IntOrSliceInt
	right *IntOrSliceInt
}

func DistressSignalPart1() {
	fmt.Println("Day 13 Part 1: Distress Signal")
	packetLines := utils.ReadFileLinesToStringSlice("day13/packets-ex.txt")

	var packets []Packet = make([]Packet, 0)
	for i := 0; i < len(packetLines); i++ {
		packetLine := packetLines[i]
		if strings.HasPrefix(packetLine, "[") {
			// construct a packet from this line and the next line
			nextPacketLine := packetLines[i+1]
			left, _ := convertPacketStringToPacketHalf(packetLine)
			right, _ := convertPacketStringToPacketHalf(nextPacketLine)
			packets = append(packets, Packet{left: left, right: right})
			i++
		}
	}

	var sum int = 0
	for i, packet := range packets {
		fmt.Println("-----------------------")
		valid := validatePacket(packet.left, packet.right)

		if valid == 1 {
			fmt.Printf("index: %d is valid", i+1)
			sum += (i + 1)
		}
	}

	fmt.Printf("The sum of the indices of the pairs who are in the correct order is %d\n\n", sum)
}

func convertPacketStringToPacketHalf(packetLine string) (*IntOrSliceInt, error) {
	arrStack := make([]*IntOrSliceInt, 0)
	for i := 0; i < len(packetLine); i++ {
		pRune := packetLine[i]
		if pRune != ',' {
			if pRune == '[' {
				sliceIntPtr := new([]IntOrSliceInt)
				*sliceIntPtr = make([]IntOrSliceInt, 0)
				arrStack = append(arrStack, &IntOrSliceInt{arr: sliceIntPtr})
			} else if pRune == ']' {
				// reached the end of the current slice, pop it off the top of the stack
				// and add it to the arr property of the new top of the stack
				newArr := arrStack[len(arrStack)-1]
				arrStack = arrStack[:len(arrStack)-1]
				if len(arrStack) == 0 {
					return newArr, nil
				} else {
					arrPtr := arrStack[len(arrStack)-1].arr
					*arrPtr = append(*arrPtr, *newArr)
				}
			} else {
				convInt, _ := strconv.Atoi(string(pRune))
				newInt := IntOrSliceInt{v: &convInt}
				arrPtr := arrStack[len(arrStack)-1].arr
				*arrPtr = append(*arrPtr, newInt)
			}
		}
	}
	return nil, errors.New("something went wrong processing input")
}

func validatePacket(left *IntOrSliceInt, right *IntOrSliceInt) int {
	isLeftArray := isArray(left)
	isRightArray := isArray(right)
	if !isLeftArray && !isRightArray {
		fmt.Printf("comparing %d vs %d\n", *left.v, *right.v)
		if *left.v > *right.v {
			fmt.Println("packets are out of order because left is greater than right")
			return -1
		} else if *left.v == *right.v {
			return 0
		}
		fmt.Printf("%d < %d so packets are in order", *left.v, *right.v)
		return 1
	} else if isLeftArray && isRightArray {
		leftArr := left.arr
		rightArr := right.arr

		i := 0
		for ; i < len(*leftArr) && i < len(*rightArr); i++ {
			lefti := (*leftArr)[i]
			righti := (*rightArr)[i]

			validity := validatePacket(&lefti, &righti)
			if validity != 0 {
				return validity
			}
		}

		// right side ran out
		if len(*leftArr) > len(*rightArr) && i == len(*rightArr) {
			fmt.Println("right side ran out")
			return -1
		} else if len(*leftArr) < len(*rightArr) && i == len(*leftArr) {
			fmt.Println("left side ran out")
			return 1
		}
	} else {
		var arrToConvert *IntOrSliceInt
		if isLeftArray && !isRightArray {
			arrToConvert = right
		} else if !isLeftArray && isRightArray {
			arrToConvert = left
		}

		convertToArr(arrToConvert)
		return validatePacket(left, right)
	}
	return 0
}

func isArray(intOrSlice *IntOrSliceInt) bool {
	is := intOrSlice.arr != nil
	return is
}

func convertToArr(intOrSlice *IntOrSliceInt) {
	var tmpIntPtr *int = intOrSlice.v

	sliceIntPtr := new([]IntOrSliceInt)
	*sliceIntPtr = make([]IntOrSliceInt, 0)
	*sliceIntPtr = append(*sliceIntPtr, IntOrSliceInt{v: tmpIntPtr})
	intOrSlice.arr = sliceIntPtr
	intOrSlice.v = nil
}
