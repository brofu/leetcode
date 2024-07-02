package graph

func solve(board [][]byte) {

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])

	visted := make([][]bool, m)
	for i := 0; i < m; i++ {
		visted[i] = make([]bool, n)
	}

	var dfs func(int, int, bool)
	dfs = func(i, j int, needChange bool) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if board[i][j] == byte('X') || visted[i][j] {
			return
		}

		visted[i][j] = true
		if needChange {
			board[i][j] = byte('X')
		}

		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1], needChange)
		}
	}

	// makr all the land on the boundary
	for _, i := range []int{0, m - 1} {
		for j := 0; j < n; j++ {
			if board[i][j] == 79 {
				dfs(i, j, false)
			}
		}
	}
	for _, j := range []int{0, n - 1} {
		for i := 0; i < m; i++ {
			if board[i][j] == 79 {
				dfs(i, j, false)
			}
		}
	}

	// change the `0` with `x`
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 79 || !visted[i][j] {
				dfs(i, j, true)
			}
		}
	}
}
