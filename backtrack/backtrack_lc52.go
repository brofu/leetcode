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

func totalNQueensPV2(n int) int {

	result := 0
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}

	isValid := func(row, col int) bool {

		// check col
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}

		// check
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		return true
	}

	var bt func(int)

	bt = func(row int) {

		//base case
		if row == n {
			result++
			return
		}

		for col := 0; col < n; col++ {
			//prune
			if !isValid(row, col) {
				continue
			}

			//choose
			rowData := []byte(board[row])
			rowData[col] = 'Q'
			board[row] = string(rowData)

			// next layer
			bt(row + 1)

			// cancel choose
			rowData[col] = '.'
			board[row] = string(rowData)
		}
	}

	bt(0)
	return result
}

/*
Key Points

1. 主，副对角线优化，
    将 valid 函数时间复杂度降为O(1)
    空间复杂度降为O(N), 如果维持棋盘的话，需要O(N^2)

2. Time Complexity
	1. O(N^N)
	2. 剪枝大概变成O(N!)
		Time complexity of prune is O(1)
*/

func totalNQueensPV3(n int) int {

	result := 0
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)

	isValid := func(row, col int) bool {

		// check col
		if cols[col] {
			return false
		}

		// check diag1
		if diag1[row-col+n-1] {
			return false
		}

		// check diag2
		if diag2[row+col] {
			return false
		}
		return true
	}

	var bt func(int)

	bt = func(row int) {

		//base case
		if row == n {
			result++
			return
		}

		for col := 0; col < n; col++ {
			//prune
			if !isValid(row, col) {
				continue
			}

			//choose
			cols[col] = true
			diag1[row-col+n-1] = true
			diag2[row+col] = true

			// next layer
			bt(row + 1)

			// cancel choose
			cols[col] = false
			diag1[row-col+n-1] = false
			diag2[row+col] = false
		}
	}

	bt(0)
	return result
}
