package bfs

/*
KP

* TC
1. Nodes need to check, N = n^2, So, up bound is O(N). But actually O(L), L is the number of cells which is 0
2. At each node, need to check the neighbors, of which the number is 8.
3. Over all, is O(N)

* SC
1. Queue number: O(N), but actually O(L)
2. Use grid as visited also. O(N)
3. Over all O(N)
*/
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	n := len(grid)
	result := 1
	grid[0][0] = 1

	q := [][]int{{0, 0}}

	for len(q) > 0 {
		l := len(q)
		for idx := 0; idx < l; idx++ {
			current := q[idx]
			if current[0] == n-1 && current[1] == n-1 {
				return result
			}
			for _, direction := range directions {
				x := current[0] + direction[0]
				y := current[1] + direction[1]
				if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] == 1 {
					continue
				}
				grid[x][y] = 1
				q = append(q, []int{x, y})
			}
		}
		result++
		q = q[l:]
	}

	return -1
}
