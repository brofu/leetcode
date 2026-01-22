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

func numEnclavesPV1(grid [][]int) int {

	result := 0
	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(int, int, bool)
	dfs = func(i, j int, count bool) {

		// base case
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 0 {
			return
		}

		// count result
		if count {
			result++
		}
		// change the land to sea
		grid[i][j] = 0

		// change the connected lands
		for _, d := range directions {
			dfs(i+d[0], j+d[1], count)
		}
	}

	// change land to sea on the board
	for _, col := range []int{0, n - 1} {
		for raw := 0; raw < m; raw++ {
			dfs(raw, col, false)
		}
	}

	for _, raw := range []int{0, m - 1} {
		for col := 0; col < n; col++ {
			dfs(raw, col, false)
		}
	}

	// loop the board to find the result
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, true)
		}
	}

	return result
}

func numEnclavesDFS(grid [][]int) int {

	result := 0
	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}
	}

	for _, row := range []int{0, m - 1} {
		for idx := range grid[row] {
			if grid[row][idx] == 1 {
				dfs(row, idx)
			}
		}
	}
	for _, col := range []int{0, n - 1} {
		for idx := 0; idx < m; idx++ {
			if grid[idx][col] == 1 {
				dfs(idx, col)
			}
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 {
				result++
			}
		}
	}
	return result
}

func numEnclavesBFS(grid [][]int) int {

	result := 0
	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	bfs := func(i, j int) {

		grid[i][j] = 0
		q := make([][]int, 1)
		q[0] = []int{i, j}

		for len(q) > 0 {
			l := len(q)
			for idx := 0; idx < l; idx++ {
				for _, direction := range directions {
					i := q[idx][0] + direction[0]
					j := q[idx][1] + direction[1]
					if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
						continue
					}
					q = append(q, []int{i, j})
					grid[i][j] = 0
				}
			}
			q = q[l:]
		}
	}

	for _, row := range []int{0, m - 1} {
		for idx := range grid[row] {
			if grid[row][idx] == 1 {
				bfs(row, idx)
			}
		}
	}
	for _, col := range []int{0, n - 1} {
		for idx := 0; idx < m; idx++ {
			if grid[idx][col] == 1 {
				bfs(idx, col)
			}
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 {
				result++
			}
		}
	}
	return result
}

func numEnclavesV2(grid [][]int) int {

	var (
		result     int
		m, n       = len(grid), len(grid[0])
		directions = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
		dfs        func(int, int, bool)
	)

	dfs = func(i, j int, record bool) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0
		if record {
			result++
		}
		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1], record)
		}
	}

	for i := 0; i < m; i++ {
		for _, j := range []int{0, n - 1} {
			if grid[i][j] == 1 {
				dfs(i, j, false)
			}
		}
	}

	for _, i := range []int{0, m - 1} {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j, false)
			}
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 {
				dfs(i, j, true)
			}
		}
	}
	return result
}
