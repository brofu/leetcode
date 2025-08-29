package bfs

/*
KP

* TC
1. O(M*N), M,N is the size of maze. But actually, should be L,  L is the cells the value is '.'
2. For each node, need to check the 4 directions
3. So, over all is O(4*L)

* SC
1. Queue size, O(L),
*/
func nearestExit(maze [][]byte, entrance []int) int {

	result := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(maze), len(maze[0])

	maze[entrance[0]][entrance[1]] = '+'
	q := [][]int{entrance}

	for len(q) > 0 {
		l := len(q)
		for idx := 0; idx < l; idx++ {

			i, j := q[idx][0], q[idx][1]
			if (i != entrance[0] || j != entrance[1]) && (i == 0 || i == m-1 || j == 0 || j == n-1) {
				return result
			}
			for _, direction := range directions {
				x := i + direction[0]
				y := j + direction[1]
				if x < 0 || x >= m || y < 0 || y >= n || maze[x][y] == '+' {
					continue
				}
				q = append(q, []int{x, y})
				maze[x][y] = '+'
			}
		}
		result++
		q = q[l:]
	}

	return -1
}
