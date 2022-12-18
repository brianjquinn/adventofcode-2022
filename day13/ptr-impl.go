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

type PacketPtr struct {
	left  *IntOrSliceInt
	right *IntOrSliceInt
}

func DistressSignalPart1Pointer() {
	fmt.Println("Day 13 Part 1: Distress Signal (Pointer Implementation that worked on example data but not on the actual data)")
	packetLines := utils.ReadFileLinesToStringSlice("day13/packets.txt")

	packetsPtr := parsePacketsPtr(packetLines)

	v1Sum, v2Sum := 0, 0
	for i, packetPtr := range packetsPtr {
		fmt.Println("-----------------------")
		validV1 := validatePtrPacketV1(packetPtr.left, packetPtr.right)
		if validV1 == 1 {
			fmt.Printf("Ptr V1: index: %d is valid", i+1)
			v1Sum += (i + 1)
		}

		validV2 := validatePtrPacketV2(packetPtr.left, packetPtr.right)
		if validV2 < 0 {
			fmt.Printf("Ptr V2:index: %d is valid", i+1)
			v2Sum += (i + 1)
		}
	}
	fmt.Printf("The sum of the indices of the pairs who are in the correct order is Ptr V1: %d, Ptr V2: %d \n\n", v1Sum, v2Sum)
}

func parsePacketsPtr(packetLines []string) []PacketPtr {
	var packets []PacketPtr = make([]PacketPtr, 0)
	for i := 0; i < len(packetLines); i++ {
		packetLine := packetLines[i]
		if strings.HasPrefix(packetLine, "[") {
			nextPacketLine := packetLines[i+1]
			left, _ := convertPacketStringToPacketHalf(packetLine)
			right, _ := convertPacketStringToPacketHalf(nextPacketLine)
			packets = append(packets, PacketPtr{left: left, right: right})
			i++
		}
	}
	return packets
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

func validatePtrPacketV1(left *IntOrSliceInt, right *IntOrSliceInt) int {
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

			validity := validatePtrPacketV1(&lefti, &righti)
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
		return validatePtrPacketV1(left, right)
	}
	return 0
}

func validatePtrPacketV2(left *IntOrSliceInt, right *IntOrSliceInt) int {
	leftArr := left.arr
	rightArr := right.arr
	var lowerLen int = 0
	if len(*leftArr) < len(*rightArr) {
		lowerLen = len(*leftArr)
	} else {
		lowerLen = len(*rightArr)
	}
	validity := 0
	for i := 0; i < lowerLen; i++ {
		lefti := (*leftArr)[i]
		leftIsArray := isArray(&lefti)
		righti := (*rightArr)[i]
		rightIsArray := isArray(&righti)

		if !leftIsArray && !rightIsArray {
			validity = *lefti.v - *righti.v
		} else if !leftIsArray && rightIsArray {
			convertToArr(&lefti)
			validity = validatePtrPacketV2(&lefti, &righti)
		} else if leftIsArray && !rightIsArray {
			convertToArr(&righti)
			validity = validatePtrPacketV2(&lefti, &righti)
		} else if leftIsArray && rightIsArray {
			validity = validatePtrPacketV2(&lefti, &righti)
		}

		if validity != 0 {
			break
		}
	}

	if validity == 0 {
		if len(*leftArr) > len(*rightArr) {
			validity = 1
		} else if len(*leftArr) < len(*rightArr) {
			validity = -1
		}
	}

	return validity
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
