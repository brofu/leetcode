package dfs

/*
*
KP
 1. Note the input is '1' 0r '0', instead of 1 or 0.
 2. 为什么每次遇到岛屿，都要用 DFS 算法把岛屿「淹了」呢？主要是为了省事，避免维护 visited 数组。 因为 dfs 函数遍历到值为 0 的位置会直接返回，所以只要把经过的位置都设置为 0，就可以起到不走回头路的作用。

Time Complexity
1. O(N+4N) = O(N) N = m*n
2. Loop all the node outside, m*n
3. Overall O(N)

Space Complexity
1. For DFS, Recursive stack depth: worst at `O(N)`
2. For BFS, Queue length: worst at `O(N)`
3. Worst case example, `m=1, n = N, and all equals 1`
4. Use `grid` as  visited, no extra spaces
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

func numIslandsPV1(grid [][]byte) int {

	count := 0
	m, n := len(grid), len(grid[0])

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		if grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0

		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				count += 1
				dfs(i, j)
			}
		}
	}

	return count
}

func numIslandsPV2(grid [][]byte) int {

	result := 0
	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

	var dfs func(int, int)

	dfs = func(i, j int) {

		// base case
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0

		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			result += 1
			dfs(i, j)
		}
	}
	return result
}

func numIslandsBFS(grid [][]byte) int {

	result := 0
	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

	bfs := func(i, j int) {

		grid[i][j] = 0

		queue := make([][]int, 1)
		queue[0] = []int{i, j}

		for len(queue) > 0 {
			l := len(queue)
			for idx := 0; idx < l; idx++ {
				node := queue[idx]
				for _, direction := range directions {
					i := node[0] + direction[0]
					j := node[1] + direction[1]
					if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
						continue
					}
					queue = append(queue, []int{i, j})
					grid[i][j] = 0
				}
			}
			queue = queue[l:]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			result += 1
			bfs(i, j)
		}
	}

	return result
}
