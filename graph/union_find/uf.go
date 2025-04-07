package uf

/*
Object: O(1) time complexity with the following APIs
*/

type UF interface {
	Connected(p, q int) bool
	Connect(q, p int)
	Count() int
}

/*
A basic version of Union-Find without performance optimization
*/
// Complexity O(N)
type BasicUF struct {
	parent []int
	count  int
}

func NewBasicUF(n int) UF {
	uf := BasicUF{
		parent: make([]int, n),
		count:  n,
	}

	for idx := range uf.parent {
		uf.parent[idx] = idx
	}

	return &uf
}

func (buf *BasicUF) Count() int {
	return buf.count
}

// O(N)
func (buf *BasicUF) Connected(p, q int) bool {
	rootP := buf.find(p)
	rootQ := buf.find(q)
	return rootP == rootQ
}

// O(N)
func (buf *BasicUF) find(p int) int {

	for buf.parent[p] != p {
		p = buf.parent[p]
	}
	return p
}

// O(N)
func (buf *BasicUF) Connect(p, q int) {
	rootP := buf.parent[p]
	rootQ := buf.parent[q]
	if rootP == rootQ {
		return
	}
	buf.parent[rootP] = rootQ
	buf.count--
}

// WeightedUF use an array to make the `weights` (how many sub-nodes one node has)
// Complexity O(lgN)
type WeightedUF struct {
	parent []int
	child  []int
	count  int
}

func NewWeightedUF(n int) UF {
	uf := WeightedUF{
		parent: make([]int, n),
		count:  n,
	}
	for idx := range uf.parent {
		uf.parent[idx] = idx
		uf.child[idx] = 1
	}
	return &uf
}

func (uf *WeightedUF) Connected(p, q int) bool {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	return rootP == rootQ
}

func (uf *WeightedUF) Connect(p, q int) {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	if rootP == rootQ {
		return
	}
	uf.count--
	if uf.child[rootP] < uf.child[rootQ] {
		uf.parent[rootP] = rootQ
		uf.child[rootQ] += uf.child[rootP]
	} else {
		uf.parent[rootQ] = rootP
		uf.child[rootP] += uf.child[rootQ]
	}
}

func (uf *WeightedUF) Count() int {
	return uf.count
}

func (uf *WeightedUF) find(p int) int {

	for uf.parent[p] != p {
		p = uf.parent[p]
	}
	return p
}

// PathComapctedUF use an array to make the `weights` (how many sub-nodes one node has)
// Complexity O(lgN)
type PathCompactedUF struct {
	parent []int
	count  int
}

func NewPathCompactedUF(n int) UF {
	uf := PathCompactedUF{
		parent: make([]int, n),
		count:  n,
	}
	for idx := range uf.parent {
		uf.parent[idx] = idx
	}
	return &uf
}

func (uf *PathCompactedUF) Connected(p, q int) bool {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	return rootP == rootQ
}

// Iterative version of `Path Compacting`
func (uf *PathCompactedUF) find(x int) int {

	for x != uf.parent[x] {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}

	return uf.parent[x]
}

// Recursive version of `Path Compacting`
func (uf *PathCompactedUF) findV2(x int) int {
	if x != uf.parent[x] {
		uf.parent[x] = uf.findV2(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *PathCompactedUF) Connect(p, q int) {
	rootP := uf.findV2(p)
	rootQ := uf.findV2(q)
	if rootP == rootQ {
		return
	}
	uf.parent[rootP] = rootQ
	uf.count--
}

func (uf *PathCompactedUF) Count() int {
	return uf.count
}
