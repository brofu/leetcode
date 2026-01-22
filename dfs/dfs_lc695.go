package dfs

import "github.com/brofu/leetcode/common"

/*
*
KP.
 1. Calculate the number of sub lands, similar to the post-order traverse of binary tree
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

func maxAreaOfIslandDFS(grid [][]int) int {

	result, temp := 0, 0
	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right

	var dfs func(int, int)

	dfs = func(i, j int) {

		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0
		temp++

		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			temp = 0
			dfs(i, j)
			if result < temp {
				result = temp
			}
		}
	}

	return result
}

func maxAreaOfIslandBFS(grid [][]int) int {

	result, temp := 0, 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right

	m, n := len(grid), len(grid[0])

	bfs := func(i, j int) {

		grid[i][j] = 0
		temp++

		q := [][]int{{i, j}}

		for len(q) > 0 {
			l := len(q)
			for idx := 0; idx < l; idx++ {
				current := q[idx]
				for _, direction := range directions {
					i := current[0] + direction[0]
					j := current[1] + direction[1]
					if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
						continue
					}
					grid[i][j] = 0
					temp++
					q = append(q, []int{i, j})
				}
			}
			q = q[l:]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			temp = 0
			bfs(i, j)
			if result < temp {
				result = temp
			}
		}
	}

	return result
}

func maxAreaOfIslandBFSV2(grid [][]int) int {

	result, temp := 0, 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right

	m, n := len(grid), len(grid[0])

	bfs := func(i, j int) {

		grid[i][j] = 0
		temp++

		q := [][]int{{i, j}}

		idx := 0
		for idx < len(q) {
			current := q[idx]
			for _, direction := range directions {
				i := current[0] + direction[0]
				j := current[1] + direction[1]
				if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
					continue
				}
				grid[i][j] = 0
				temp++
				q = append(q, []int{i, j})
			}
			idx++
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			temp = 0
			bfs(i, j)
			if result < temp {
				result = temp
			}
		}
	}

	return result
}

func maxAreaOfIslandDFSV2(grid [][]int) int {

	var (
		result, tempCount int
		m, n              = len(grid), len(grid[0])
		directions        = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		dfs               func(int, int)
	)

	dfs = func(i, j int) {

		// base case
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0
		tempCount++

		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				tempCount = 0
				dfs(i, j)
				if tempCount > result {
					result = tempCount
				}
			}
		}
	}
	return result
}
