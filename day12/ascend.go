package day12

type TerrainSquare struct {
	row      int
	col      int
	distance int
}

func ascend(row int, col int, terrainMap [][]rune) int {
	visited := make([][]bool, len(terrainMap))
	for i := 0; i < len(terrainMap); i++ {
		visited[i] = make([]bool, len(terrainMap[0]))
	}

	var queue []*TerrainSquare = []*TerrainSquare{{row: row, col: col, distance: 0}}
	visited[row][col] = true
	for len(queue) > 0 {
		currTerrainSq := queue[0]
		if len(queue) == 1 {
			queue = []*TerrainSquare{}
		} else {
			queue = queue[1:]
		}

		currElevation := terrainMap[currTerrainSq.row][currTerrainSq.col]

		if currElevation == 'E' {
			return currTerrainSq.distance
		}

		// up down left right
		var rowMods [4]int = [4]int{-1, 1, 0, 0}
		var colMods [4]int = [4]int{0, 0, -1, 1}

		// iterate through the row/col modifications in order to see if
		// we can enqueue one of the neighbors (up, down, left, right)
		for i := 0; i < 4; i++ {
			newRow := currTerrainSq.row + rowMods[i]
			newCol := currTerrainSq.col + colMods[i]

			if newRow >= 0 && newRow < len(terrainMap) && newCol >= 0 && newCol < len(terrainMap[row]) && !visited[newRow][newCol] {
				neighborElevation := terrainMap[newRow][newCol]
				if neighborElevation-currElevation <= 1 || currElevation == 'S' {
					queue = append(queue, &TerrainSquare{row: newRow, col: newCol, distance: currTerrainSq.distance + 1})
					visited[newRow][newCol] = true
				}
			}
		}
	}
	return -1
}
