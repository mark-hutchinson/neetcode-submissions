func searchMatrix(matrix [][]int, target int) bool {
	numRows := len(matrix)
	if numRows == 0 {
		return false
	}
	numCols := len(matrix[0])
	if numCols == 0 {
		return false
	}
	left :=0
	right := numRows - 1

	// First, we need to find the possible row using Binary Search
	for left <= right {
		mid := left + (right - left) / 2
		// fmt.Printf("searching rows: left=%d, right=%d, mid=%d\n", left, right, mid)
		if target >= matrix[mid][0] && target <= matrix[mid][numCols - 1] {
			// fmt.Printf("searching row %d\n", mid)
			return searchRow(matrix[mid], target)
		} else if target < matrix[mid][0] {
			// search left
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func searchRow(row []int, target int) bool {
	left := 0
	right := len(row)
	for left <= right {
		mid := left + (right - left) / 2
		// fmt.Printf("searching row: left=%d, right=%d, mid=%d\n", left, right, mid)

		if row[mid] == target {
			return true
		} else if target < row[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
