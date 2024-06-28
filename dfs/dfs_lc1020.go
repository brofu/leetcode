package dfs

func numEnclaves(grid [][]int) int {

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

		grid[i][j] = 0

		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1])
		}
	}
	// change the lands on boundary to water
	for _, i := range []int{0, m - 1} {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 { // land
				dfs(i, j)
			}
		}
	}
	for _, j := range []int{0, n - 1} {
		for i := 0; i < m; i++ {
			if grid[i][j] == 1 { // land
				dfs(i, j)
			}
		}
	}

	// count the results
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 { // land
				result++
			}
		}
	}

	return result
}
