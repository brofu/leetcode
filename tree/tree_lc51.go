package tree

import "strings"

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
