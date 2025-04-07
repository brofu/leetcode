package uf

/*
Validating Tree with Graph

Key Points:

1. There is NO ring
2. There is ONLY ONE connected component
*/

func validTree(n int, edges [][]int) bool {

	uf := NewPathCompactedUF(len(edges))

	for _, edge := range edges {
		if uf.Connected(edge[0], edge[1]) {
			return false
		}
		uf.Connect(edge[0], edge[1])
	}
	return uf.Count() == 1
}
