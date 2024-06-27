package dfs

/**
KP
	1.	Similar as 200. Only need to change the islands on boarding to sea
*/
func closedIsland(grid [][]int) int {
	result := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right
	m, n := len(grid), len(grid[0])

	var dfs func(int, int)

	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		if grid[i][j] == 1 {
			return
		}

		grid[i][j] = 1 // change the land to sea

		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1])
		}
	}

	// change the lands on boarder to water
	for _, i := range []int{0, m - 1} {
		for j := 0; j < n; j++ {
			dfs(i, j)
		}
	}
	for _, j := range []int{0, n - 1} {
		for i := 0; i < m; i++ {
			dfs(i, j)
		}
	}

	for i := 0; i < m; i = i + 1 {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				result += 1
				dfs(i, j)
			}
		}
	}
	return result
}
