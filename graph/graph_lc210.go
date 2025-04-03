package graph

import (
	"github.com/brofu/leetcode/common/array"
)

/*
*
KP
 1. Direction between the nodes is imported. `Depending on` or `Depended By`?
    `graph[from] = append(graph[from], to)` is `Depended By`. Need to reverse the order
    `graph[to] = append(graph[to], from)` is `depending on`. No need to reverse the order.
 2. PostOrder is used. ==> what about preorder?
*/
func findOrder(numCourses int, prerequisites [][]int) []int {

	result := []int{}
	isRing := false
	visisted := make([]bool, numCourses)
	onPath := make([]bool, numCourses)

	// construct graph
	graph := make([][]int, numCourses)
	for i := 0; i < len(graph); i++ {
		graph[i] = []int{}
	}
	for _, dep := range prerequisites { // the order is important here.
		from, to := dep[1], dep[0]
		graph[from] = append(graph[from], to)
	}

	// check if there is ring
	var dfs func(int)
	dfs = func(node int) {
		if onPath[node] {
			isRing = true
			return
		}
		if visisted[node] || isRing {
			return
		}

		visisted[node] = true
		onPath[node] = true
		for _, next := range graph[node] {
			dfs(next)
		}
		onPath[node] = false
		result = append(result, node)
	}

	for i := 0; i < len(graph); i++ {
		dfs(i)
	}

	if isRing {
		return []int{}
	}

	// reverse
	array.ReverseInt(result)

	return result
}

/*
*
KP
 1. WRONG
 2. The `preorder` version is WRONG. Because, we need to
*/
func findOrderPerOrder(numCourses int, prerequisites [][]int) []int {

	result := []int{}
	isRing := false
	visisted := make([]bool, numCourses)
	onPath := make([]bool, numCourses)

	// construct graph
	graph := make([][]int, numCourses)
	for i := 0; i < len(graph); i++ {
		graph[i] = []int{}
	}
	for _, dep := range prerequisites { // the order is important here.
		from, to := dep[1], dep[0]
		graph[from] = append(graph[from], to)
	}

	// check if there is ring
	var dfs func(int)
	dfs = func(node int) {
		if onPath[node] {
			isRing = true
			return
		}
		if visisted[node] || isRing {
			return
		}

		visisted[node] = true
		onPath[node] = true
		result = append(result, node)
		for _, next := range graph[node] {
			dfs(next)
		}
		onPath[node] = false
	}

	for i := 0; i < len(graph); i++ {
		dfs(i)
	}

	if isRing {
		return []int{}
	}

	return result
}

func findOrderPV1(numCourses int, prerequisites [][]int) []int {

	// construct graph
	graph := make([][]int, numCourses)

	for _, record := range prerequisites {
		from, to := record[1], record[0]
		graph[from] = append(graph[from], to)
	}

	var traverse func(val int)

	hasCycle := false
	onPath := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	path := []int{}

	traverse = func(val int) {

		// base case
		if hasCycle {
			return
		}

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

		//post-order location
		onPath[val] = false
		path = append(path, val)
	}

	for index := range graph {
		traverse(index)
	}

	if hasCycle {
		return []int{}
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func findOrderBFS(numCourses int, prerequisites [][]int) []int {

	// construct graph
	graph := make([][]int, numCourses)

	for _, record := range prerequisites {
		from, to := record[1], record[0]
		graph[from] = append(graph[from], to)
	}

	// construct in-degree array
	indegree := make([]int, numCourses)

	for _, record := range prerequisites {
		to := record[0]
		indegree[to]++
	}

	// initialize q
	q := []int{}

	for root, ind := range indegree {
		if ind == 0 {
			q = append(q, root)
		}
	}

	result := []int{}
	count := 0
	for len(q) > 0 {
		size := len(q)
		for idx := 0; idx < size; idx++ {
			count++
			val := q[idx]
			result = append(result, val)
			for _, next := range graph[val] {
				indegree[next]--
				if indegree[next] == 0 { // mean no pre node any more
					q = append(q, next)
				}
			}
		}
		q = q[size:]
	}

	if count == numCourses {
		return result
	}
	return []int{}
}
