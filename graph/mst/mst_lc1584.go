package mst

import (
	"container/heap"
	"math"
	"sort"

	"github.com/brofu/leetcode/common"
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

Time Complexity: O(N^2logN)
 1. Construct a full graph, there is N^2 sides
 2. Priority Queue operations O(N^2logN)

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
		cut(to)
	}

	return sum
}

/*
minCostConnectPointsPrimOptimization optimize the Prim version

Time Complexity: O(n^2)

How does the optimization work?

 1. No need to construct the FULL graph for this scenario, since in this scenarios, we are sure there all the nodes should be in the graph
 2. So, we are use another description of `Prim` algorithm
    a. `dist[i]` means, the shorted distance between the node `i` and the `mst` (all the nodes in the `mst`)
    b. all the nodes should be in the `mst` tree, so, we have `for edge < n-1` loop
*/

func minCostConnectPointsPrimOptimization(points [][]int) int {

	n := len(points)

	visisted := make([]bool, n)
	dist := make([]int, n)
	for idx := range dist {
		dist[idx] = math.MaxInt32
	}

	edge, node, sum := 0, 0, 0

	for edge < n-1 {

		visisted[node] = true // put node into mst
		nextNode := -1

		for idx := 0; idx < n; idx++ {

			if visisted[idx] {
				continue
			}

			// calculate the distance between i and node
			currentDistance := calculateDistance(points[idx][0], points[idx][1], points[node][0], points[node][1])

			// update d[i], if there is new shorter distance
			dist[idx] = common.MinInt(dist[idx], currentDistance)

			// get the minimum distance  from all the un-visted nodes
			if nextNode == -1 || dist[idx] < dist[nextNode] {
				nextNode = idx
			}
		}

		edge++
		sum += dist[nextNode]
		node = nextNode
	}

	return sum
}

func calculateDistance(x1, y1, x2, y2 int) int {
	return Abs(x1, x2) + Abs(y1, y2)
}
