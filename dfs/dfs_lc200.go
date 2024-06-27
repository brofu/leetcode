package dfs

/**
KP
	1.	Note the input is '1' 0r '0', instead of 1 or 0.
*/
func numIslands(grid [][]byte) int {

	result := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right
	m, n := len(grid), len(grid[0])

	var dfs func(int, int)
	dfs = func(i, j int) {

		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		if grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0 // change the land to sea

		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				result += 1
				// change all the lands of this island to sea
				dfs(i, j)
			}
		}
	}

	return result
}
