package graph

import (
	"sort"

	"github.com/brofu/leetcode/common"
)

func minCostConnectPoints(points [][]int) int {

	cost := 0
	edges := [][]int{}
	n := len(points)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ix, iy := points[i][0], points[i][1]
			jx, jy := points[j][0], points[j][1]
			edges = append(edges, []int{i, j, common.AbsIntSub(ix, jx) + common.AbsIntSub(iy, jy)})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	uf := NewUF(n)

	for _, edge := range edges {
		x := edge[0]
		y := edge[1]
		if uf.Connected(x, y) {
			continue
		}
		uf.Union(x, y)
		cost += edge[2]
	}

	return cost
}
