package graph

/**
KP.
	How to separate the graph and tree?
		1. 对于添加的这条边，如果该边的两个节点本来就在同一连通分量里，那么添加这条边会产生环；反之，如果该边的两个节点不在同一连通分量里，则添加这条边不会产生环。
		2. There is ONLY 1 tree finally. `uf.Count() == 1`
*/
func validTree(n int, edges [][]int) bool {

	uf := NewUF(n)

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		if uf.Connected(a, b) {
			return false
		}
		uf.Union(a, b)
	}

	return uf.Count() == 1
}
