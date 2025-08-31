package bfs

/*
KP

TC:

1. Filter all the rotten. O(m*n)
2. Setup fresh map, O(m*n)
3. Check all the node. O(m*n)
4. For each node, get the 4 neighbors. That's around O(4)
5. Check if there is still fresh. O(1)
6. Over all, O(m*n)

SC
1. Queue size. O(m*n)
2. Fresh map. O(m*n)
3. Over all O(m*n)
*/

func orangesRotting(grid [][]int) int {

	var (
		result     = 0
		m, n       = len(grid), len(grid[0])
		rottens    = [][]int{}
		freshes    = make(map[int]struct{})
		directions = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				rottens = append(rottens, []int{i, j})
				continue
			}
			if grid[i][j] == 1 {
				freshes[i*n+j] = struct{}{}
			}
		}
	}
	if len(freshes) == 0 {
		return result
	}
	if len(rottens) == 0 {
		return -1
	}

	q := rottens
	for len(q) > 0 {
		l := len(q)
		for idx := 0; idx < l; idx++ {
			current := q[idx]
			for _, direction := range directions {
				x := current[0] + direction[0]
				y := current[1] + direction[1]
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 2 || grid[x][y] == 0 {
					continue
				}
				grid[x][y] = 2
				delete(freshes, x*n+y)
				q = append(q, []int{x, y})
			}
		}
		result++
		if len(freshes) == 0 {
			return result
		}
		q = q[l:]
	}

	return -1
}

/*
Optimization:
1. User counter to record the number of fresh.
*/
func orangesRottingSpaceOptimized(grid [][]int) int {

	var (
		result, freshes int
		m, n            = len(grid), len(grid[0])
		rottens         = [][]int{}
		directions      = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				rottens = append(rottens, []int{i, j})
				continue
			}
			if grid[i][j] == 1 {
				freshes++
			}
		}
	}
	if freshes == 0 {
		return result
	}
	if len(rottens) == 0 {
		return -1
	}

	q := rottens
	for len(q) > 0 {
		l := len(q)
		for idx := 0; idx < l; idx++ {
			current := q[idx]
			for _, direction := range directions {
				x := current[0] + direction[0]
				y := current[1] + direction[1]
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 2 || grid[x][y] == 0 {
					continue
				}
				grid[x][y] = 2
				freshes--
				q = append(q, []int{x, y})
			}
		}
		result++
		if freshes == 0 {
			return result
		}
		q = q[l:]
	}

	return -1
}
