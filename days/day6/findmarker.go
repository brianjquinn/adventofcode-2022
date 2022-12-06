package day6

import (
	"errors"
	"strconv"
)

func FindEndIdxOfMarkerWithUniqLength(length int, input string) (int, error) {
	trackerMap := make(map[rune]int)
	var uniqCount int = 0
	for i, char := range input {
		lastSeen, present := trackerMap[char]
		if !present || lastSeen < i-uniqCount {
			uniqCount++
			trackerMap[char] = i
		} else {
			uniqCount = i - lastSeen
			trackerMap[char] = i
		}

		if uniqCount == length {
			return i + 1, nil
		}
	}

	return -1, errors.New("marker of length " + strconv.Itoa(length) + " does not exist")
}
