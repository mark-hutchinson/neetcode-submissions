func numIslands(grid [][]byte) int {
	seen := make([][]bool, len(grid))
	for row :=0; row < len(grid); row++ {
		seen[row] = make([]bool, len(grid[row]))
	}

	numIslands := 0
	for row :=0; row < len(grid); row++ {
		for col :=0; col < len(grid[row]); col++ {
			if grid[row][col] == '1' && !seen[row][col] {
				// found a new island.
				numIslands++
				// explore the island to mark all land squares as seen
				seen = explore(grid, row, col, seen)
			}

		}
	}

	return numIslands
}

func explore(grid [][]byte, row int, col int, seen [][]bool) [][]bool {
	if row < 0 || row >= len(grid) {
		return seen // row out of bounds
	}
	if col < 0 || col >= len(grid[row]) {
		return seen // col out of bounds
	}
	if grid[row][col] == '0' {
		return seen
	}
	if seen[row][col] {
		return seen
	}
	seen[row][col] = true

	seen = explore(grid, row, col-1, seen) // left
	seen = explore(grid, row, col+1, seen) // right
	seen = explore(grid, row-1, col, seen) // up
	seen = explore(grid, row+1, col, seen) // down
	return seen
}