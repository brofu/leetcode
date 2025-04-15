package mst

import (
	"container/heap"
	"math"

	"github.com/brofu/leetcode/common"
	heaplib "github.com/brofu/leetcode/common/heap"
)

/*
Key Point

1. Calculate the MIN distance from `start` to all the nodes SO FAR, and if there is SMALLER distance, update it
2. BFS approach to traverse the nodes, when there is SMALLER distance found (to some node)
*/

func networkDelayTime(times [][]int, n int, k int) int {

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
