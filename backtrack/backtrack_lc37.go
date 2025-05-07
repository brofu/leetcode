package backtrack

/*

Key points to control the workflow. Refer to the code
*/

func solveSudoku(board [][]byte) {

	found := false
	var backtrack func(int)

	backtrack = func(index int) {

		// found a solution
		if found == true {
			return
		}

		// base case
		if index == 81 {
			found = true
			return
		}

		i, j := index/9, index%9
		// KEY POINT:
		// how to remove to next grad if there is already number in current grad
		// need to return
		if board[i][j] != '.' {
			backtrack(index + 1)
			return
		}

		for ch := '1'; ch <= '9'; ch++ {
			if !isValid(board, i, j, byte(ch)) {
				continue
			}
			board[i][j] = byte(ch)
			backtrack(index + 1)
			// KEY POINT:
			// In the post-order location
			// if found one resolution, return directly,
			// otherwise the board would be reverted.
			if found {
				return
			}
			board[i][j] = byte('.')
		}
	}

	backtrack(0)
}

func isValid(board [][]byte, i, j int, ch byte) bool {

	// check the row
	for _, num := range board[i] {
		if num == ch {
			return false
		}
	}

	// check the column
	for idx := 0; idx < 9; idx++ {
		if board[idx][j] == ch {
			return false
		}
	}

	startR, endR := i/3*3, (i/3+1)*3
	startC, endC := j/3*3, (j/3+1)*3
	for k := startR; k < endR; k++ {
		for l := startC; l < endC; l++ {
			if board[k][l] == ch {
				return false
			}
		}
	}
	return true
}
