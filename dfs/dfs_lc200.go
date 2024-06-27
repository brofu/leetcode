package dfs

/**
KP
	1.	Note the input is '1' 0r '0', instead of 1 or 0.
	2.	为什么每次遇到岛屿，都要用 DFS 算法把岛屿「淹了」呢？主要是为了省事，避免维护 visited 数组。 因为 dfs 函数遍历到值为 0 的位置会直接返回，所以只要把经过的位置都设置为 0，就可以起到不走回头路的作用。
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
