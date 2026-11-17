func maxAreaOfIsland(grid [][]int) int {
	// // construct seen matrix, to track what we have seen so far
	// seen := make([][]struct{}, len(grid))
	// for i, row := range grid {
	// 	seen[i] := make([]struct{}, len(grid[i]))
	// }

	maxSize := 0

	for row, _ := range grid {
		for col, _ := range grid[row] {

			if grid[row][col] == 1 {
				// Found an island, explore it
				islandSize := exploreIsland(grid, row, col, 0)
				if islandSize > maxSize {
					maxSize = islandSize
				}
			}

		}
	}

	return maxSize
}

func exploreIsland(grid [][]int, row int, col int, size int) int {
	if row < 0 || row >= len(grid) {
		return size
	}
	if col < 0 || col >= len(grid[row]) {
		return size
	}

	if grid[row][col] == 1 {
		// This is land, increase island size
		size = size + 1
		grid[row][col] = 2 // mark as explored, avoid double counting

		// explore left
		size = exploreIsland(grid, row, col-1, size)
		// explore right
		size = exploreIsland(grid, row, col+1, size)
		// explore up
		size = exploreIsland(grid, row-1, col, size)
		// explore down
		size = exploreIsland(grid, row+1, col, size)
	}
	return size
}
