package graph

func countComponents(n int, edges [][]int) int {

	uf := NewUFWithSize(n)
	//uf := NewUF(n)
	for _, connection := range edges {
		uf.Union(connection[0], connection[1])
	}

	return uf.Count()
}
