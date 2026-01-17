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

/*

Key Model:
	1. A simple problem: 函数的输入是一个正整数 n，返回所有长度为 n 的二进制数（0、1 组成），可以按任意顺序返回答案。
	2. 这道题可以认为是数独游戏和 N 皇后问题的简化版：这道题相当于对一个长度为 n 的一维数组中的每一个位置进行穷举，其中每个位置可以填 0 或 1。
	3. 数独游戏相当于对一个 9x9 的二维数组中的每个位置进行穷举，其中每个位置可以是数字 1~9，且同一行、同一列和同一个 3x3 的九宫格内数字不能重复。
	4. N 皇后问题相当于对一个 N x N 的二维数组中的每个位置进行穷举，其中每个位置可以不放皇后或者放置皇后（相当于 0 或 1），且不能存在多个皇后在同一行、同一列或同一对角线上。
	5. 只要把这道简化版的题目的穷举过程搞明白，其他问题都迎刃而解了，无非是规则多了一些而已。

TC:
	1. Similar as 81 cells, of which one there are 9 choices. So, for an total empty one, is around 9^81
	2. But, actually, there are some celled filled in advance. If the empty number is E, then, it's around 9^E
	3. For each round, call the isValid, 9 + 9 + 9 = 27 checks
	4. So, the overall TC is 9^E

SC:
	1. Recursive Depth: 81
	2. Reuse the board. No extra space
	3. So, overall is O(1)
*/

func solveSudokuPV1(board [][]byte) {

	var (
		found = false
	)

	var bt func(int, int)

	bt = func(x, y int) {

		if found {
			return
		}

		//if 9*x+y == 81 {
		if x == 9 { // this more clear
			found = true
			return
		}

		if board[x][y] != byte('.') { // existing number, no choose/cancel-choose, explore next layer directly.
			x, y = getNext(x, y)
			bt(x, y)
			return
		}

		for num := '1'; num <= '9'; num++ {
			if !isValidPV1(board, x, y, byte(num)) { // pruning
				continue
			}

			// choose
			board[x][y] = byte(num)
			if x == 9 {
				found = true
			}

			// explore
			i, j := getNext(x, y) // important. x, y should keep same, for canceling choose.
			bt(i, j)

			// cancel choose
			if found { // reuse board, so, if found, return directly.
				return
			}
			board[x][y] = byte('.')
		}
	}

	bt(0, 0)
}

func getNext(x, y int) (int, int) {
	if y+1 == 9 {
		return x + 1, 0
	}
	return x, y + 1
}

func isValidPV1(board [][]byte, x, y int, num byte) bool {

	for i := 0; i < 9; i++ {
		if board[i][y] == num {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if board[x][j] == num {
			return false
		}
	}

	for i := (x / 3) * 3; i < (x/3+1)*3; i++ {
		for j := (y / 3) * 3; j < (y/3+1)*3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	return true
}
