package backtrack

/*
Key Points

* TC
1. If the length of the of the word is L, the recursive depth would be 4^L, (4 choices at each recursive layer)
2. And for each M*N, need to check.
3. Over all, the TC is M*N*4^L

* SC
1. The recursive depth is L.
2. If need visited[M][N] to record the visited, the SC would be O(M*N+L)
3. If use board to record the visited info, the SC is O(L)
*/
func exist(board [][]byte, word string) bool {

	m, n := len(board), len(board[0])

	result := false

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var bt func(int, int, int)

	bt = func(x, y, idx int) {

		if result {
			return
		}

		if idx == len(word) {
			result = true
			return
		}

		temp := board[x][y]
		board[x][y] = '#'
		for _, direction := range directions {
			i := x + direction[0]
			j := y + direction[1]

			// pruning
			if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == '#' || board[i][j] != word[idx] {
				continue
			}

			// choose
			// next layer
			bt(i, j, idx+1)
			// cancel
		}
		board[x][y] = temp

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				bt(i, j, 1)
				if result {
					break
				}
			}
		}
	}
	return result
}
