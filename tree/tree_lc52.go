package tree

import "strings"

func totalNQueens(n int) int {

	result := 0

	board := make([]string, n)

	for i := 0; i < n; i += 1 {
		board[i] = strings.Repeat(".", n)
	}

	backTrackNQueenII(board, 0, &result)
	return result
}

func backTrackNQueenII(board []string, row int, result *int) {

	if row == len(board) {
		*result += 1
	}

	for col := 0; col < len(board); col += 1 {
		if !isValidNQueen(board, row, col) {
			continue
		}

		// choose
		line := []byte(board[row])
		line[col] = 'Q'
		board[row] = string(line)

		// next
		backTrackNQueenII(board, row+1, result)

		// back
		line[col] = '.'
		board[row] = string(line)
	}
}
