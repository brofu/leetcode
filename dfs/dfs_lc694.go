package dfs

import (
	"fmt"
	"strings"
)

/**
KP.
	1.	How to check if 2 islands are same?
	很显然我们得想办法把二维矩阵中的「岛屿」进行转化，变成比如字符串这样的类型，然后利用 HashSet 这样的数据结构去重，最终得到不同的岛屿的个数。如果想把岛屿转化成字符串，说白了就是序列化，序列化说白了就是遍历嘛，前文 二叉树的序列化和反序列化 讲了二叉树和字符串互转，这里也是类似的。 首先，对于形状相同的岛屿，如果从同一起点出发，dfs 函数遍历的顺序肯定是一样的。需要记录撤销操作的路径
	2.	Actually, this is `backtrack`
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
