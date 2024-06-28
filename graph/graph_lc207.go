package graph

func canFinish(numCourses int, prerequisites [][]int) bool {

	ring := false
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)

	// construct the graph
	graph := make([][]int, numCourses)
	for i := 0; i < len(graph); i++ {
		graph[i] = []int{}
	}
	for _, dep := range prerequisites {
		from, to := dep[1], dep[0]
		graph[from] = append(graph[from], to)
	}

	var dfs func(int)

	dfs = func(node int) {
		if onPath[node] {
			ring = true
			return
		}

		if visited[node] || ring {
			return
		}

		visited[node] = true

		// similar as backtrack
		onPath[node] = true
		for _, next := range graph[node] {
			dfs(next)
		}
		onPath[node] = false
	}

	for i := 0; i < len(graph); i++ {
		dfs(i)
	}
	return !ring
}
