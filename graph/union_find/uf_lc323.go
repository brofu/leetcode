package uf

func countComponents(n int, edges [][]int) int {

	uf := NewPathCompactedUF(n)
	for _, edge := range edges {
		uf.Connect(edge[0], edge[1])
	}

	return uf.Count()
}
