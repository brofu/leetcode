package graph

import (
	"sort"
)

func minimumCost(n int, connections [][]int) int {

	sum := 0
	uf := NewUF(n + 1)

	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	for _, con := range connections {
		a := con[0]
		b := con[1]
		if uf.Connected(a, b) {
			continue
		}
		uf.Union(a, b)
		sum += con[2]
	}

	if uf.Count() == 2 {
		return sum
	}

	return -1
}
