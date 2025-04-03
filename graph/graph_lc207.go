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

/*
Key Point

The order of checking `visited` and `onPath` is important
*/

func canFinishPV1(numCourses int, prerequisites [][]int) bool {

	// construct the graph
	graph := make([][]int, numCourses)
	for _, record := range prerequisites {
		from, to := record[1], record[0]
		graph[from] = append(graph[from], to)
	}

	var traverse func(int)

	hasCycle := false
	onPath := make([]bool, len(graph))
	visited := make([]bool, len(graph))

	traverse = func(val int) {

		// base case
		if hasCycle {
			return
		}

		// get ring
		if onPath[val] {
			hasCycle = true
			return
		}

		if visited[val] {
			return
		}

		// pre-order location
		onPath[val] = true
		visited[val] = true
		for _, number := range graph[val] {
			traverse(number)
		}

		// post-order location
		onPath[val] = false
	}

	for index := range graph {
		traverse(index)
	}

	return !hasCycle
}
