package graph

/**
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
