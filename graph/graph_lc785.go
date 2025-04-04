package graph

/*
*
KP.
 1. Pay attention that, the graph may be not `connected`
*/
func isBipartite(graph [][]int) bool {
	result := true
	colors := make([]bool, len(graph)) // true or false stands for the 2 colors
	visited := make([]bool, len(graph))

	var dfs func(int)

	dfs = func(node int) {

		if !result { //if one the sub graph is NOT bipartite, just return
			return
		}

		visited[node] = true

		for _, next := range graph[node] {
			if !visited[next] {
				colors[next] = !colors[node]
				dfs(next)
			} else {
				if colors[next] == colors[node] {
					result = false
					return
				}
			}
		}
	}

	for node := 0; node < len(graph); node++ {
		if !visited[node] {
			dfs(node)
		}
	}

	return result
}

func isBipartiteBFS(graph [][]int) bool {

	result := true
	visited := make([]bool, len(graph))
	colors := make([]bool, len(graph))

	var bfs func(int)

	bfs = func(node int) {

		visited[node] = true
		queue := []int{node}

		for len(queue) > 0 && result {
			node = queue[0]
			queue = queue[1:]
			for _, next := range graph[node] {
				if !visited[next] {
					visited[next] = true
					colors[next] = !colors[node]
					queue = append(queue, next)
				} else {
					if colors[next] == colors[node] {
						result = false
						return
					}
				}
			}
		}
	}

	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			bfs(i)
		}
	}
	return result
}

func isBipartiteDFSPV1(graph [][]int) bool {

	var traverse func(v int)

	isBipartite := true

	visited := make([]bool, len(graph))
	color := make([]bool, len(graph))

	traverse = func(v int) {

		// base case
		if !isBipartite {
			return
		}

		if visited[v] {
			return
		}

		// pre-order location
		visited[v] = true

		for _, next := range graph[v] {
			if !visited[next] {
				color[next] = !color[v]
				traverse(next)
			} else {
				if color[next] == color[v] {
					isBipartite = false
					return
				}
			}
		}
	}

	for idx := 0; idx < len(graph); idx++ {
		traverse(idx)
	}

	return isBipartite
}

/*
Key Point

1. The difference of BFS from problem `Cycle Detection` and `Topological Sorting` (Directed Graph)
2. This problems is UnDirecte Graph
*/

func isBipartiteBFSPV1(graph [][]int) bool {

	isBipartite := true
	visited := make([]bool, len(graph))
	color := make([]bool, len(graph))

	bfs := func(v int) {

		if !isBipartite {
			return
		}

		q := []int{v}
		visited[v] = true

		for len(q) > 0 {
			size := len(q)
			for idx := 0; idx < size; idx++ {
				v := q[idx]
				for _, next := range graph[v] {
					if !visited[next] {
						color[next] = !color[v]
						q = append(q, next)
						visited[next] = true
					} else {
						if color[next] == color[v] {
							isBipartite = false
							return
						}
					}
				}
			}
			q = q[size:]
		}
	}

	for idx := range graph {
		if !visited[idx] {
			bfs(idx)
		}
	}

	return isBipartite
}
