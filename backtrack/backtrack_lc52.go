package backtrack

import (
	"strings"
)

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

func totalNQueensPV1(n int) int {

	result := 0

	board := make([][]int, n, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n, n)
	}

	var backtrack func([][]int, int, *int)

	isValid := func(board [][]int, raw, col int) bool {

		for r := 0; r < raw; r++ {
			if board[r][col] == 1 {
				return false
			}
		}

		for r, c := raw-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
			if board[r][c] == 1 {
				return false
			}
		}

		for r, c := raw-1, col+1; r >= 0 && c < len(board); r, c = r-1, c+1 {
			if board[r][c] == 1 {
				return false
			}
		}

		return true
	}

	backtrack = func(board [][]int, raw int, result *int) {

		// ending of choose
		if len(board) == raw {
			*result += 1
			return
		}

		for col := 0; col < len(board); col++ {

			if !isValid(board, raw, col) {
				continue
			}

			// choose
			board[raw][col] = 1

			// depth + 1
			backtrack(board, raw+1, result)

			// cancel choose
			board[raw][col] = 0
		}
	}

	backtrack(board, 0, &result)
	return result
}
