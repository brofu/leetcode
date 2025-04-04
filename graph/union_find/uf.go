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
type BasicUF struct {
	parent []int
	count  int
}

func NewBasicUF(n int) UF {
	uf := BasicUF{
		parent: make([]int, n),
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
}
