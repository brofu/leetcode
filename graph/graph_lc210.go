package graph

import (
	"github.com/brofu/leetcode/common/array"
)

/**
KP
	1.	Direction between the nodes is imported. `Depending on` or `Depended By`?
		`graph[from] = append(graph[from], to)` is `Depended By`. Need to reverse the order
		`graph[to] = append(graph[to], from)` is `depending on`. No need to reverse the order.
	2.	PostOrder is used. ==> what about preorder?
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

/**
KP
	1.	WRONG
	2.  The `preorder` version is WRONG. Because, we need to
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
