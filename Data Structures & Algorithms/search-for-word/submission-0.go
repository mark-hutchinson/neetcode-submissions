type Point struct {
	Row int
	Col int
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}
	
	for row := range board {
		for col := range board[row] {
			seenMap := make(map[Point]bool)
			if search(row, col, board, []byte(word), seenMap) {
				return true
			}
		}
	}

	return false
}

func search(row int, col int, board [][]byte, word []byte, seenMap map[Point]bool) bool {
	if len(word) == 0 {
		return true // nothing left to search for
	}
	if row < 0 || row >= len(board) {
		return false // out of bounds
	}
	if col <0 || col >= len(board[row]) {
		return false // out of bounds
	}
	if seenMap[Point{Row: row, Col:col}] == true {
		return false // don't consider same square twice
	}

	if board[row][col] == word[0] {
		// this character matches
		remainingWord := word[1:]
		seenMap[Point{Row: row, Col:col}] = true

		// look left
		found := search(row, col-1, board, remainingWord, seenMap)
		if found{
			return true
		}
		// look right
		found = search(row, col+1, board, remainingWord, seenMap)
		if found{
			return true
		}
		// look up
		found = search(row-1, col, board, remainingWord, seenMap)
		if found{
			return true
		}
		// look down
		found = search(row+1, col, board, remainingWord, seenMap)
		if found{
			return true
		}

		// if this square is not part of solution, remove from seen
		seenMap[Point{Row: row, Col:col}] = false
	}

	return false
}