package dfs

/*
*
KP.
 1. How to check if an island is the sub island to another one?
    当岛屿 B 中所有陆地在岛屿 A 中也是陆地的时候，岛屿 B 是岛屿 A 的子岛。反过来说，如果岛屿 B 中存在一片陆地，在岛屿 A 的对应位置是海水，那么岛屿 B 就不是岛屿 A 的子岛。那么，我们只要遍历 grid2 中的所有岛屿，把那些不可能是子岛的岛屿排除掉，剩下的就是子岛。
*/
func countSubIslands(grid1 [][]int, grid2 [][]int) int {

	result := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right
	m, n := len(grid1), len(grid1[0])

	var dfs func(int, int)

	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid2[i][j] == 0 {
			return
		}

		grid2[i][j] = 0

		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 { // island of grid2, but water in the grid1
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				result += 1
				dfs(i, j)
			}
		}
	}
	return result
}

func countSubIslandsPV1(grid1 [][]int, grid2 [][]int) int {

	result := 0
	m, n := len(grid1), len(grid1[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid2[i][j] == 0 {
			return
		}

		grid2[i][j] = 0
		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				result++
				dfs(i, j)
			}
		}
	}
	return result
}

func countSubIslandsBFS(grid1 [][]int, grid2 [][]int) int {
	result := 0
	m, n := len(grid1), len(grid1[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right

	bfs := func(i, j int) {
		grid2[i][j] = 0

		q := make([][]int, 1)
		q[0] = []int{i, j}

		for len(q) > 0 {
			l := len(q)
			for idx := 0; idx < l; idx++ {
				for _, direction := range directions {
					i := q[idx][0] + direction[0]
					j := q[idx][1] + direction[1]
					if i < 0 || i >= m || j < 0 || j >= n || grid2[i][j] == 0 {
						continue
					}
					q = append(q, []int{i, j})
					grid2[i][j] = 0
				}
			}
			q = q[l:]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				bfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				result++
				bfs(i, j)
			}
		}
	}
	return result

}
