package dfs

/*
*
KP
 1. Similar as 200. Only need to change the islands on boarding to sea

Time Complexity
* Similar with 200

Space Complexity
* Similar with 200
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

func closedIslandPV1(grid [][]int) int {

	result := 0
	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right

	var dfs func(int, int)

	dfs = func(i, j int) {

		//base case
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 1 {
			return
		}

		grid[i][j] = 1

		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}

	}

	for _, row := range []int{0, m - 1} {
		for idx := range grid[row] {
			if grid[row][idx] == 1 {
				continue
			}
			dfs(row, idx)
		}
	}
	for _, col := range []int{0, n - 1} {
		for idx := 0; idx < m; idx++ {
			if grid[idx][col] == 1 {
				continue
			}
			dfs(idx, col)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 {
				continue
			}
			result++
			dfs(i, j)
		}
	}

	return result
}

/*

TC:
	1. Clean. 2*(m+n)
	2. Flood Fill. 4(m-1)*(n-1). So, overall is O(m*n)

*/
func closedIslandBFS(grid [][]int) int {

	result := 0
	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left and right

	bfs := func(i, j int) {

		grid[i][j] = 1

		queue := make([][]int, 1)
		queue[0] = []int{i, j}

		for len(queue) > 0 {
			l := len(queue)

			for idx := 0; idx < l; idx++ {
				node := queue[idx]
				for _, direction := range directions {
					i := node[0] + direction[0]
					j := node[1] + direction[1]
					if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 1 {
						continue
					}
					queue = append(queue, []int{i, j})
					grid[i][j] = 1
				}
			}
			queue = queue[l:]
		}
	}

	for _, row := range []int{0, m - 1} {
		for idx := range grid[row] {
			if grid[row][idx] == 1 {
				continue
			}
			bfs(row, idx)
		}
	}
	for _, col := range []int{0, n - 1} {
		for idx := 0; idx < m; idx++ {
			if grid[idx][col] == 1 {
				continue
			}
			bfs(idx, col)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 {
				continue
			}
			result++
			bfs(i, j)
		}
	}

	return result
}

func closedIslandV2(grid [][]int) int {

	var (
		result     int
		m, n       = len(grid), len(grid[0])
		dfs        func(int, int)
		directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	)

	dfs = func(i, j int) {

		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 1 {
			return
		}

		grid[i][j] = 1

		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}
	}

	for _, i := range []int{0, m - 1} { //remove the un-closed islands
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for _, j := range []int{0, n - 1} {
			if grid[i][j] == 0 {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				result++
				dfs(i, j)
			}
		}
	}

	return result
}
