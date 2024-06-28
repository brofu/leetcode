package graph

func allPathsSourceTarget(graph [][]int) [][]int {

	result := [][]int{}
	path := []int{}

	var dfs func([]int, int, *[][]int)
	dfs = func(path []int, node int, result *[][]int) {

		path = append(path, node)

		if node == len(graph)-1 { // hit the n-1 node
			*result = append(*result, append([]int{}, path...))
			return
		}

		for _, i := range graph[node] {
			dfs(path, i, result)
		}

		path = path[:len(path)-1]
	}

	dfs(path, 0, &result)

	return result
}
