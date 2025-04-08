package mst

import (
	"sort"

	uflib "github.com/brofu/leetcode/graph/union_find"
)

func minimumCost(n int, connections [][]int) int {

	uf := uflib.NewPathCompactedUF(n + 1)

	sum := 0
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	for _, connection := range connections {
		if !uf.Connected(connection[0], connection[1]) {
			uf.Connect(connection[0], connection[1])
			sum += connection[2]
		}
	}

	if uf.Count() != 2 {
		return -1
	}
	return sum
}
