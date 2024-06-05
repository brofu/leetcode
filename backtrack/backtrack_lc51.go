package backtrack

import (
	"strings"
)

func solveNQueens(n int) [][]string {

	result := [][]string{}

	board := make([]string, n)

	for i := 0; i < n; i += 1 {
		board[i] = strings.Repeat(".", n)
	}

	backTrackNQueen(board, 0, &result)

	return result
}

func backTrackNQueen(board []string, row int, result *[][]string) {

	// base
	if row == len(board) {
		temp := make([]string, len(board))
		copy(temp, board)
		*result = append(*result, temp)
	}

	for col := 0; col < len(board); col += 1 {

		if !isValidNQueen(board, row, col) {
			continue
		}

		// choose a path
		line := []byte(board[row])
		line[col] = 'Q'
		board[row] = string(line)

		// next choose
		backTrackNQueen(board, row+1, result)

		// cancel choose
		line[col] = '.'
		board[row] = string(line)
	}
}

func isValidNQueen(board []string, row, col int) bool {

	// check same column
	for index := 0; index < row; index += 1 {
		if board[index][col] == 'Q' {
			return false
		}
	}

	// check right-up diagonal
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// check left-up diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func solveNQueensPV1(n int) [][]string {

	result := make([][]string, 0, n)
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}

	isValid := func(board []string, raw, col int) bool {

		// check same col
		for r := 0; r < raw; r++ {
			if board[r][col] == 'Q' {
				return false
			}
		}

		// check up-left
		for r, c := raw-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
			if board[r][c] == 'Q' {
				return false
			}
		}

		// check up-right
		for r, c := raw-1, col+1; r >= 0 && c < len(board); r, c = r-1, c+1 {
			if board[r][c] == 'Q' {
				return false
			}
		}

		return true
	}

	var trackback func([]string, int, *[][]string)

	trackback = func(board []string, raw int, result *[][]string) {

		// ending of choose
		if len(board) == raw {
			temp := make([]string, len(board))
			copy(temp, board)
			*result = append(*result, temp)
			return
		}

		for col := 0; col < len(board); col++ {

			if !isValid(board, raw, col) {
				continue
			}

			// choose
			line := []byte(board[raw])
			line[col] = 'Q'
			board[raw] = string(line)

			// depth + 1
			trackback(board, raw+1, result)

			// remove from list
			line[col] = '.'
			board[raw] = string(line)
		}
	}

	trackback(board, 0, &result)
	return result
}
