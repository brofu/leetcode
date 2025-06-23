package dfs

import (
	"fmt"
	"strings"
)

/*
KP.

 1. How to check if 2 islands are same?
    很显然我们得想办法把二维矩阵中的「岛屿」进行转化，变成比如字符串这样的类型，然后利用 HashSet 这样的数据结构去重，最终得到不同的岛屿的个数。如果想把岛屿转化成字符串，说白了就是序列化，序列化说白了就是遍历嘛，前文 二叉树的序列化和反序列化 讲了二叉树和字符串互转，这里也是类似的。 首先，对于形状相同的岛屿，如果从同一起点出发，dfs 函数遍历的顺序肯定是一样的。需要记录撤销操作的路径

 2. About BFS
    如何利用坐标确定岛屿的形状？ BFS v.s. DFS?

Time Complexity
1. Recursive call times O(N+4N) = O(5N)
2. strings.Builder.WirteString : worst case O(2*N)
3. out side for i, j loop: O(N)
Overall O(N)

Space Complexity
1. Recursive stack depth: worst O(N)
2. strings.Builder worst: O(2*N) = O(N)
3. The map depends. Should
*/
func numDistinctIslands(grid [][]int) int {
	directions := [][]int{{-1, 0, 1}, {1, 0, 2}, {0, -1, 3}, {0, 1, 4}} // directions and direction numbers
	m, n := len(grid), len(grid[0])

	var dfs func(int, int, *strings.Builder, int)
	dfs = func(i, j int, sb *strings.Builder, dirNum int) {

		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		if grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0 // change the land to sea

		sb.WriteString(fmt.Sprintf("%d", dirNum))
		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1], sb, dir[2])
		}
		sb.WriteString(fmt.Sprintf("%d", -dirNum))
	}

	islandsMap := make(map[string]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				sb := &strings.Builder{}
				dfs(i, j, sb, 5)
				islandsMap[sb.String()] = true
			}
		}
	}
	return len(islandsMap)
}

func numDistinctIslandsPV1(grid [][]int) int {

	directions := [][]int{{-1, 0, 1}, {1, 0, 2}, {0, -1, 3}, {0, 1, 4}} // directions and direction numbers
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int, sb *strings.Builder, dir int)
	dfs = func(i, j int, sb *strings.Builder, dir int) {

		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		sb.WriteString(fmt.Sprintf("%d,", dir))
		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1], sb, direction[2])
		}
		sb.WriteString(fmt.Sprintf("%d,", -dir))
	}

	islandsMap := make(map[string]struct{})
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				sb := &strings.Builder{}
				dfs(i, j, sb, 5)
				islandsMap[sb.String()] = struct{}{}
			}
		}
	}

	return len(islandsMap)
}

func numDistinctIslandsBFS(grid [][]int) int {

	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // directions and direction numbers
	islandsMap := make(map[string]struct{})

	bfs := func(i, j int) {

		buf := strings.Builder{}

		grid[i][j] = 0
		q := [][]int{{i, j}}

		for len(q) > 0 {
			l := len(q)
			for idx := 0; idx < l; idx++ {
				current := q[idx]
				for _, direction := range directions {
					ni := current[0] + direction[0]
					nj := current[1] + direction[1]
					if ni < 0 || ni >= m || nj < 0 || nj >= n || grid[ni][nj] == 0 {
						continue
					}
					grid[ni][nj] = 0
					q = append(q, []int{ni, nj})
					buf.WriteByte(byte(ni - i))
					buf.WriteByte(byte(nj - j))
				}
			}
			q = q[l:]
		}
		islandsMap[buf.String()] = struct{}{}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
			}
		}
	}

	return len(islandsMap)
}
