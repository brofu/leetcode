package graph

type UnionFind interface {
	Union(int, int)
	Connected(int, int) bool
	Count() int
}

/**
An implementation of algorithm of Union Find.
Time Complexity: O(lgN)
*/
type UFWithSize struct {
	count  int
	parent []int // use array to implement the connection of the trees
	size   []int // how many sub-nodes in the tree with node x as root
}

func (uf *UFWithSize) Union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX == rootY {
		return
	}

	if uf.size[x] < uf.size[y] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX] // use uf.size to balance the tree
	} else {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	}

	uf.count--
}

func (uf *UFWithSize) find(x int) int {
	for uf.parent[x] != x {
		x = uf.parent[x]
	}
	return x
}

func (uf *UFWithSize) Connected(x, y int) bool {
	rootX := uf.find(x)
	rootY := uf.find(y)

	return rootX == rootY
}

func (uf *UFWithSize) Count() int {
	return uf.count
}

func NewUFWithSize(n int) UnionFind {
	uf := UFWithSize{
		count:  n,
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return &uf

}

/**
An implementation of algorithm of Union Find.
Time Complexity: O(1)
	*  based on the `path compression`
*/
type UF struct {
	count  int
	parent []int // use array to implement the connection of the trees
}

func (uf *UF) Union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX != rootY {
		uf.parent[rootX] = rootY
		uf.count--
	}
}

func (uf *UF) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UF) Connected(x, y int) bool {
	rootX := uf.find(x)
	rootY := uf.find(y)

	return rootX == rootY
}

func (uf *UF) Count() int {
	return uf.count
}

func NewUF(n int) UnionFind {
	uf := UF{
		count:  n,
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return &uf
}
