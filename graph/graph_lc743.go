package graph

import (
	"container/heap"
)

type state struct {
	id                int
	distanceFromStart int
}

func networkDelayTime(times [][]int, n int, k int) int {

	longestTime := 0

	// init graph
	graph := make([][][]int, n+1)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([][]int, 0)
	}
	for _, latency := range times {
		from, to, t := latency[0], latency[1], latency[2]
		graph[from] = append(graph[from], []int{to, t})
	}

	// call dijkstra
	dij := func(graph [][][]int, start int) []int {

		distanceTo := make([]int, len(graph))
		for i := range distanceTo {
			distanceTo[i] = -1
		}
		distanceTo[start] = 0

		pq := make(statePriorityQueue, 0)
		heap.Init(&pq)
		startState := &state{start, 0}
		heap.Push(&pq, startState)

		for len(pq) > 0 {
			currentNode := heap.Pop(&pq).(*state)
			currentId := currentNode.id
			currentDistance := currentNode.distanceFromStart

			if currentDistance > distanceTo[currentId] {
				continue
			}

			for _, neighbor := range graph[currentId] { // calculate the neighbor
				nextId := neighbor[0]
				latency := neighbor[1]
				nextDistance := distanceTo[currentId] + latency
				if distanceTo[nextId] < 0 || nextDistance < distanceTo[nextId] {
					distanceTo[nextId] = nextDistance
					heap.Push(&pq, &state{nextId, nextDistance})
				}
			}
		}

		return distanceTo
	}

	pathes := dij(graph, k)

	for _, path := range pathes[1:] {
		if path == -1 {
			return -1
		}
		if longestTime < path {
			longestTime = path
		}
	}

	return longestTime
}
