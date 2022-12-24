package day15

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/brianjquinn/adventofcode-2022/utils"
)

type Sensor struct {
	pos []int64
	r   int64
}

func BeaconExclusionZonePart2() {
	fmt.Println("Day 15 Part 2: Beacon Exclusion Zone")

	sensorsAndBeacons := utils.ReadFileLinesToStringSlice("day15/sensors-and-beacons.txt")
	var re = regexp.MustCompile(`(?mi)x=(-*\d+), y=(-*\d+)`)
	sensors := make([]Sensor, 0)
	min := int64(0)
	max := int64(4000001)
	for _, sbLine := range sensorsAndBeacons {
		matches := re.FindAllStringSubmatch(sbLine, -1)
		sx, _ := strconv.ParseInt(matches[0][1], 10, 64)
		sy, _ := strconv.ParseInt(matches[0][2], 10, 64)
		bx, _ := strconv.ParseInt(matches[1][1], 10, 64)
		by, _ := strconv.ParseInt(matches[1][2], 10, 64)
		manhattanDistance := utils.Abs(sx-bx) + utils.Abs(sy-by)
		sensors = append(sensors, Sensor{pos: []int64{sx, sy}, r: manhattanDistance})
	}
	done := false
	var x int64
	var y int64
	for _, sensor := range sensors {
		// walk the perimeter 1 step outside of the sensor's range
		x = sensor.pos[0]
		y = sensor.pos[1] - sensor.r - 1
		direction := [][]int64{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}}
		innerLoopMax := int(sensor.r) + 1
		for i := 0; i < 4; i++ {
			for j := 0; j < innerLoopMax; j++ {
				xX := x + direction[i][0]
				yY := y + direction[i][1]
				x = xX
				y = yY
				if xX >= min && xX < max && yY >= min && yY < max {
					// check if manhattan dist from point to sensor is greater
					// than the sensor's range
					coveredCount := 0
					for _, sensor := range sensors {
						pointToSensorMDist := utils.Abs(xX-sensor.pos[0]) + utils.Abs(yY-sensor.pos[1])
						if pointToSensorMDist <= sensor.r {
							coveredCount++
						}
					}

					if coveredCount == 0 {
						done = true
						break
					}
				}
			}
			if done {
				break
			}
		}
		if done {
			break
		}
	}

	fmt.Printf("The tuning frequency for the beacon emitting the distress signal is %d\n", x*4000000+y)
}
