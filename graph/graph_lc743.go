package graph

import (
	"container/heap"
	"math"

	"github.com/brofu/leetcode/common"
	heaplib "github.com/brofu/leetcode/common/heap"
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

/*
Key Point

1. Calculate the MIN distance from `start` to all the nodes SO FAR, and if there is SMALLER distance, update it
2. BFS approach to traverse the nodes, when there is SMALLER distance found (to some node)
*/

func networkDelayTimePV1(times [][]int, n int, k int) int {

	graph := make([][][]int, n+1)
	dist := make([]int, n+1)
	for idx := range dist {
		dist[idx] = math.MaxInt32
	}

	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		graph[from] = append(graph[from], []int{to, weight})
	}

	pq := heaplib.NewInterfacePriorityQueue(func(data []interface{}, i, j int) bool {
		return data[i].([]int)[1] < data[j].([]int)[1]
	})

	dist[k] = 0
	heap.Push(pq, []int{k, 0})

	for pq.Len() > 0 {
		cNode := heap.Pop(pq).([]int)
		to, curDist := cNode[0], cNode[1]

		if dist[to] < curDist { // there is already a smaller distance
			continue
		}

		for _, nextNode := range graph[to] {
			t, d := nextNode[0], nextNode[1]
			nextNodeDistance := curDist + d
			if dist[t] == math.MaxInt32 || nextNodeDistance < dist[t] {
				dist[t] = nextNodeDistance
				heap.Push(pq, []int{t, nextNodeDistance})
			}
		}
	}

	result := 0
	for _, time := range dist[1:] {
		if time == math.MaxInt32 {
			return -1
		}
		result = common.MaxInt(result, time)
	}

	return result
}
