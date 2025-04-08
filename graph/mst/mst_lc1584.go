package mst

import (
	"sort"

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
