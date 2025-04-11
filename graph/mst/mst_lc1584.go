package mst

import (
	"container/heap"
	"fmt"
	"sort"

	heaplib "github.com/brofu/leetcode/common/heap"
	uflib "github.com/brofu/leetcode/graph/union_find"
)

/*
Kruskal

Key Point:

1. How to present point info? Use the index in the `points`

Time Complexity
O(N^2logN^2) = O(2N^2logN)
*/
func minCostConnectPoints(points [][]int) int {

	n := len(points)
	cost := [][]int{}
	sum := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cost = append(cost, []int{i, j, Abs(points[i][0], points[j][0]) + Abs(points[i][1], points[j][1])})
		}
	}

	sort.Slice(cost, func(i, j int) bool {
		return cost[i][2] < cost[j][2]
	})

	uf := uflib.NewPathCompactedUF(n)
	for _, c := range cost {
		if !uf.Connected(c[0], c[1]) {
			uf.Connect(c[0], c[1])
			sum += c[2]
		}
	}

	return sum
}

func Abs(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

/*
Prime

Key Point:
1. Use the index in the points array as the node
2. MUST construct the FULL graph (ALL the side between any two nodes. Otherwise, there would be sides missing. Refer to the case in case 1
*/
func minCostConnectPointsPrim(points [][]int) int {

	// build graph
	graph := make([][][]int, len(points))
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			cost := Abs(points[i][0], points[j][0]) + Abs(points[i][1], points[j][1])
			graph[i] = append(graph[i], []int{j, cost})
			graph[j] = append(graph[j], []int{i, cost})
		}
	}

	fmt.Println("flag", points)
	fmt.Println("flag", graph)
	pq := heaplib.NewInterfacePriorityQueue(func(data []interface{}, i, j int) bool {
		return data[i].([]int)[2] < data[j].([]int)[2]
	})
	inMST := make([]bool, len(points))
	sum := 0

	var cut = func(n int) {
		for _, neigher := range graph[n] {
			if len(neigher) == 0 {
				continue
			}
			to, weight := neigher[0], neigher[1]
			if inMST[to] {
				continue
			}
			heap.Push(pq, []int{n, to, weight})
		}
	}

	// start from 0
	inMST[0] = true
	cut(0)

	for pq.Len() > 0 {
		current := heap.Pop(pq).([]int)
		to, weight := current[1], current[2]
		if inMST[to] {
			continue
		}
		inMST[to] = true
		sum += weight
		fmt.Println("flag2", sum, current)
		cut(to)
	}

	return sum
}

func minCostConnectPointsPrimV2(points [][]int) int {

	return 0
}
