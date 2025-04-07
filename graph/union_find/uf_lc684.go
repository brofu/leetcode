package uf

func findRedundantConnection(edges [][]int) []int {

	uf := NewPathCompactedUF(len(edges) + 1)

	duplicated := [][]int{}
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if !uf.Connected(from, to) {
			uf.Connect(from, to)
		} else {
			duplicated = append(duplicated, edge)
		}
	}
	if len(duplicated) > 0 {
		return duplicated[len(duplicated)-1]
	}
	return []int{}
}
