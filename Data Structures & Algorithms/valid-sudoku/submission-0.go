func isValidSudoku(board [][]byte) bool {

	// Check each row and column
	for i:=0; i<9; i++ {
		row := getRow(board, i)
		if !isValid(row) {
			return false
		}
		col := getCol(board, i)
		if !isValid(col) {
			return false
		}
	}

	// check the squares
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			square := getSquare(board, i*3, j*3)
			if !isValid(square) {
				return false
			}
		}
	}

	return true
}

func getRow(board [][]byte, row int) []byte{
	return board[row]
}

func getCol(board [][]byte, col int) []byte {
	result := make([]byte, 9)
	for i:=0; i<9;i++ {
		result[i] = board[i][col]
	}
	return result
}

func getSquare(board [][]byte, offsetX int, offsetY int) []byte {
	result := make([]byte, 0)
	for i:= offsetX; i < offsetX +3; i++ {
		for j:= offsetY; j < offsetY+3; j++ {
			result = append(result, board[i][j])
		}
	}
	return result
}

func isValid(inputs []byte) bool {
	counts := make([]int, 10)
	for _, input := range inputs {
		if input == '.' {
			continue
		}
		numeric := input - '0'

		counts[numeric]++
		if counts[numeric] > 1 {
			return false
		}
	}
	return true
}