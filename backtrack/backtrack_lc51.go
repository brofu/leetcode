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

func solveNQueensPV2(n int) [][]string {

	result := make([][]string, 0, n)
	board := make([]string, n)
	for i := 0; i < len(board); i++ {
		board[i] = strings.Repeat(".", n)
	}

	valid := func(board []string, row, col int) bool {

		// check same col
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}

		// check up-left
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		// check up-right
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		return true
	}

	var tb func([]string, int)

	tb = func(board []string, row int) {
		// base case
		if row == n {
			temp := append([]string{}, board...)
			result = append(result, temp)
			return
		}

		for col := 0; col < n; col++ {
			// pruning.
			if !valid(board, row, col) {
				continue
			}

			// choose
			raw := []byte(board[row])
			raw[col] = 'Q'
			board[row] = string(raw)

			// next layer
			tb(board, row+1)
			// cancel choose
			raw[col] = '.'
			board[row] = string(raw)
		}
	}

	tb(board, 0)
	return result
}

/*
Key Points:

How to control the flow?
 1. LC37, for the Sudoku problem, we need to choose a number for each `grad`
 2. N-Queen problem, we need to choose a location for each `row`
*/

func solveNQueensPV3(n int) [][]string {

	isValid := func(board []string, row, col int) bool {

		// check the col
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}

		// KEY POINTS:
		// Why not need to check the row?
		// => there is `cancel` operation at the post-order location
		// check the row
		//for i := 0; i < n; i++ {
		//	if board[row][i] == 'Q' {
		//		return false
		//	}
		//}

		// check the line
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

	result := [][]string{}

	var tb func([]string, int)

	tb = func(board []string, row int) {

		// base case
		if row == n {
			temp := make([]string, n)
			copy(temp, board)
			result = append(result, temp)
			return
		}

		for col := 0; col < n; col++ {
			// prune
			if !isValid(board, row, col) {
				continue
			}

			// choose
			rowData := []byte(board[row])
			rowData[col] = 'Q'
			board[row] = string(rowData)

			// next layer
			tb(board, row+1)

			// cancel choose
			rowData[col] = '.'
			board[row] = string(rowData)
		}
	}

	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}

	tb(board, 0)
	return result
}

/*

TC:
	1. Worst case is n^n (at each layer, n choices and there are n recursive layers), but actually, it's more near to O(n!).
		* The 1st layer has n choices, and
		* the 2nd layer has n-1 choices.(actually should be less and n-1, the same column and the distinct)
		* So, n * (n-1) * ... * 1 = n!
	2. For each choice, need to check isValid,
		* check column at each row, O(N)
		* check distinct, O(N)
		* overall O(N)
	3. So overall for search ??
	3. For each result, we need to copy, which is n^2

SC:
	1. Recursive depth is n
	2. For each result found,


*/
func solveNQueensPV4(n int) [][]string {

	var (
		result [][]string
		bt     func(int)
		board  = make([][]string, n)
	)

	for idx := range board {
		board[idx] = make([]string, n)
		for i := 0; i < n; i++ {
			board[idx][i] = "."
		}
	}

	isValid := func(row, col int) bool {

		for i := 0; i < row; i++ {
			if board[i][col] == "Q" {
				return false
			}
		}

		for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == "Q" {
				return false
			}
		}

		for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == "Q" {
				return false
			}
		}
		return true
	}

	bt = func(layer int) {

		// base case
		if layer == n {
			temp := make([]string, n)
			for idx := range board {
				temp[idx] = strings.Join(board[idx], "")
			}
			result = append(result, temp)
			return
		}

		for i := 0; i < n; i++ {

			//pruning
			if !isValid(layer, i) {
				continue
			}

			// choose
			board[layer][i] = "Q"
			// explore
			bt(layer + 1)

			// cancel choose
			board[layer][i] = "."
		}
	}

	bt(0)

	return result
}
