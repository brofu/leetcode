package uf

/*
The Union-Find way to resolve the problem

# Key Points

1. The `Dummy` node
2. Divide the `O` into 2 groups. One group should be replaced and the other one should NOT
*/
func solve(board [][]byte) {

	m := len(board)
	n := len(board[0])

	uf := NewPathCompactedUF(m*n + 1)

	dumy := m * n
	for j := 0; j < n; j++ {
		if string(board[0][j]) == "O" {
			uf.Connect(dumy, j)
		}
		if string(board[m-1][j]) == "O" {
			uf.Connect(dumy, (m-1)*n+j)
		}
	}
	for i := 0; i < m; i++ {
		if string(board[i][0]) == "O" {
			uf.Connect(dumy, i*n)
		}
		if string(board[i][n-1]) == "O" {
			uf.Connect(dumy, i*n+n-1)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if string(board[i][j]) == "O" {
				arounds := [][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
				for _, node := range arounds {
					if string(board[node[0]][node[1]]) == "O" {
						uf.Connect(i*n+j, node[0]*n+node[1])
					}
				}
			}
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !uf.Connected(dumy, i*n+j) {
				board[i][j] = 'X'
			}
		}
	}
}
