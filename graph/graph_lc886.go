package graph

/*
*
KP.
 1. How to transfer the problem to Bipartite problem?
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

func possibleBipartitionDFSPV1(n int, dislikes [][]int) bool {

	// construct un-directed graph
	graph := make([][]int, n+1)
	for _, record := range dislikes {
		from, to := record[0], record[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	result := true

	// traverse
	var traverse func(v int)

	visited := make([]bool, n+1)
	color := make([]bool, n+1)

	traverse = func(v int) {

		// base case
		if !result {
			return
		}

		if visited[v] {
			return
		}

		// pre-order location
		visited[v] = true

		for _, next := range graph[v] {
			if visited[next] {
				if color[next] == color[v] {
					result = false
					return
				}
			} else {
				color[next] = !color[v]
				traverse(next)
			}
		}
	}

	for idx := range graph {
		traverse(idx)
	}

	return result
}

func possibleBipartitionBFSPV1(n int, dislikes [][]int) bool {

	// construct un-directed graph
	graph := make([][]int, n+1)
	for _, record := range dislikes {
		from, to := record[0], record[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	result := true
	visited := make([]bool, n+1)
	color := make([]bool, n+1)

	traverse := func(v int) {

		if !result {
			return
		}

		q := []int{v}
		visited[v] = true

		for len(q) > 0 {
			size := len(q)
			for idx := 0; idx < size; idx++ {
				v := q[idx]

				for _, next := range graph[v] {
					if visited[next] {
						if color[next] == color[v] {
							result = false
							return
						}
					} else {
						color[next] = !color[v]
						visited[next] = true
						q = append(q, next)
					}
				}
			}
			q = q[size:]
		}
	}

	for idx := range graph {
		if !visited[idx] {
			traverse(idx)
		}
	}

	return result

}
