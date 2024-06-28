package graph

/**
KP.
	1.	How to transfer the problem to Bipartite problem?
	如果把每个人看做图中的节点，相互讨厌的关系看做图中的边，那么 dislikes 数组就可以构成一幅图；又因为题目说互相讨厌的人不能放在同一组里，相当于图中的所有相邻节点都要放进两个不同的组； 那就回到了「双色问题」，如果能够用两种颜色着色所有节点，且相邻节点颜色都不同，那么按照颜色把这些节点分成两组就行了。 所以解法是，dislikes 构造成一幅图，然后执行二分图的判定算法即可
*/
func possibleBipartition(n int, dislikes [][]int) bool {

	result := true
	visited := make([]bool, n+1)
	color := make([]bool, n+1)

	// construct the graph
	graph := make([][]int, n+1)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}

	for _, r := range dislikes {
		p := r[0]
		q := r[1]
		graph[p] = append(graph[p], q)
		graph[q] = append(graph[q], p)
	}

	var dfs func(int)
	dfs = func(node int) {

		if visited[node] || !result {
			return
		}

		visited[node] = true

		for _, next := range graph[node] {
			if !visited[next] {
				color[next] = !color[node]
				dfs(next)
			} else {
				if color[next] == color[node] {
					result = false
				}
			}
		}
	}

	for i := 1; i < n+1; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	return result
}
