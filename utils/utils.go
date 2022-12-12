package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileLinesToStringSlice(filename string) []string {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func MapStringArrToUInt64Arr(strArr []string) []uint64 {
	intArr := make([]uint64, len(strArr))
	for i, v := range strArr {
		intArr[i], _ = strconv.ParseUint(v, 10, 64)
	}
	return intArr
}
