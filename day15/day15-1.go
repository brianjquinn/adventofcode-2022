package day15

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

func BeaconExclusionZonePart1() {
	fmt.Println("Day 15 Part 1: Beacon Exclusion Zone")

	sensorsAndBeacons := utils.ReadFileLinesToStringSlice("day15/sensors-and-beacons.txt")
	var re = regexp.MustCompile(`(?mi)x=(-*\d+), y=(-*\d+)`)
	sensorPositions := make([][]int, 0)
	beaconPositions := make([][]int, 0)
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for _, sbLine := range sensorsAndBeacons {
		matches := re.FindAllStringSubmatch(sbLine, -1)
		sx, _ := strconv.Atoi(matches[0][1])
		sy, _ := strconv.Atoi(matches[0][2])
		sensorPositions = append(sensorPositions, []int{sx, sy})
		bx, _ := strconv.Atoi(matches[1][1])
		by, _ := strconv.Atoi(matches[1][2])
		beaconPositions = append(beaconPositions, []int{bx, by})

		if sx < minX {
			minX = sx
		}
		if bx < minX {
			minX = bx
		}

		if sx > maxX {
			maxX = sx
		}
		if bx > maxX {
			maxX = bx
		}

		if sy < minY {
			minY = sy
		}
		if by < minY {
			minY = by
		}

		if sy > maxY {
			maxY = sy
		}
		if by > maxY {
			maxY = by
		}
	}

	yOfInterest := 2000000
	coveredXPositions := make(map[int]bool)
	for i := 0; i < len(sensorPositions); i++ {
		sensor := sensorPositions[i]
		beacon := beaconPositions[i]
		manhattanDist := int(utils.Abs(int64(sensor[0])-int64(beacon[0])) + utils.Abs(int64(sensor[1])-int64(beacon[1])))
		diff := int(utils.Abs(int64(yOfInterest - sensor[1])))
		width := (manhattanDist-diff)*2 + 1
		// figure out the starting x index in the target row and iterate
		// in order to produce the x indexes that would be covered by this
		// sensor in the target row - store those indexes in a map so we get
		// a unique count of the indexes in the target row covered by all sensors
		start := sensor[0] - (width / 2)
		for i := 0; i < width; i++ {
			coveredXPositions[start+i] = true
		}
	}

	// remove out beacons that are in the target row
	for _, beacon := range beaconPositions {
		if coveredXPositions[beacon[0]] && beacon[1] == yOfInterest {
			delete(coveredXPositions, beacon[0])
		}
	}
	fmt.Printf("The number of positions that cannot contain a beacon with y = %d is %d\n\n", yOfInterest, len(coveredXPositions))
}
