package dfs

import "github.com/brofu/leetcode/common"

/**
KP.
	1.	Calculate the number of sub lands, similar to the post-order traverse of binary tree
*/
func maxAreaOfIsland(grid [][]int) int {
	result := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right
	m, n := len(grid), len(grid[0])

	var dfs func(int, int) int
	dfs = func(i, j int) int {

		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}

		if grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0
		result := 1
		for _, dir := range directions {
			result += dfs(i+dir[0], j+dir[1])
		}
		return result
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				result = common.MaxInt(result, dfs(i, j))
			}
		}
	}
	return result
}
