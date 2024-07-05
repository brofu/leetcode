package graph

import (
	"container/heap"
)

type tuple struct {
	to  int
	pro float64
}

/**
KP.
	Why this problem can be resolve by Dijkstra? The pre-conditions of Dijkstra are:
		1.	Directed Weighted Graph
		2.	No Negative Weight
*/
func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {

	// construct graph
	graph := make([][]tuple, n)
	for i := 0; i < n; i++ {
		graph[i] = []tuple{}
	}
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		pro := succProb[i]
		graph[from] = append(graph[from], tuple{to, pro})
		graph[to] = append(graph[to], tuple{from, pro})
	}

	// construct dp
	probabilityTo := make([]float64, n)
	for i := range probabilityTo {
		probabilityTo[i] = -1
	}
	probabilityTo[start_node] = 1

	pq := make(stateFloat64PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &stateFloat64{start_node, 1})

	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*stateFloat64)
		currentId := currentNode.id
		currentPro := currentNode.probability

		if currentId == end_node {
			return currentPro
		}

		if currentPro < probabilityTo[currentId] { // there is a better path
			continue
		}

		for _, node := range graph[currentId] { // check the neighbors
			nextId := node.to
			nextPro := currentPro * node.pro
			if probabilityTo[nextId] < 0 || nextPro > probabilityTo[nextId] {
				probabilityTo[nextId] = nextPro
				heap.Push(&pq, &stateFloat64{nextId, nextPro})
			}
		}

	}
	return 0
}
