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

/**
Union Find version
*/

func solveUF(board [][]byte) {

	m, n := len(board), len(board[0])
	uf := NewUF(m*n + 1)

	// connect the 'O' on boundary with a dummy root
	dummy := m * n
	for _, i := range []int{0, m - 1} {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				uf.Union(i*n+j, dummy)
			}
		}
	}
	for _, j := range []int{0, n - 1} {
		for i := 0; i < m; i++ {
			if board[i][j] == 'O' {
				uf.Union(i*n+j, dummy)
			}
		}
	}

	// connect the 'O' with arounding 'O'
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'X' {
				continue
			}
			for _, dir := range directions {
				x := i + dir[0]
				y := j + dir[1]
				if board[x][y] == 'O' {
					uf.Union(i*n+j, x*n+y)
				}
			}
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' && !uf.Connected(n*i+j, dummy) {
				board[i][j] = 'X'
			}
		}
	}
}
