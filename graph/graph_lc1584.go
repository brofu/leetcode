package graph

import (
	"math"
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

/**
KP.
	Version without Union Find
*/
func minCostConnectPointsWithoutUF(points [][]int) int {

	sum := 0
	n := len(points)

	distances := make([]int, n)
	for i := range distances {
		distances[i] = math.MaxInt
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			distances[j] = common.MinInt(distances[j], common.AbsIntSub(points[i][0], points[j][0])+common.AbsIntSub(points[i][1], points[j][1]))
			if distances[i+1] > distances[j] {
				distances[i+1], distances[j] = distances[j], distances[i+1]
				points[i+1], points[j] = points[j], points[i+1]
			}
		}
		sum += distances[i+1]
	}
	return sum
}
