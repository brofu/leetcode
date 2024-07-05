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
	size   []int // how many sub-nodes in the tree with node x as root. Balance the tree
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
Time Complexity: O(1). How? Basically, we don't care about the struct of the tree, we just care about the `ROOT` of the tree. So, we may `compress the path` (make all the nodes have the `root` node as its parent node)
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
		uf.parent[x] = uf.find(uf.parent[x]) // compress the path here.
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

/**
Priority Queue of `state`
*/
type statePriorityQueue []*state

func (this statePriorityQueue) Less(i, j int) bool {
	return this[i].distanceFromStart < this[j].distanceFromStart
}

func (this statePriorityQueue) Len() int {
	return len(this)
}

func (this statePriorityQueue) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *statePriorityQueue) Push(x any) {
	*this = append(*this, x.(*state))
}

func (this *statePriorityQueue) Pop() any {
	x := (*this)[0]
	*this = (*this)[1:]
	return x
}

type stateXY struct {
	x, y              int
	distanceFromStart int
}

/**
Priority Queue of `stateXY`
*/
type stateXYPriorityQueue []*stateXY

func (this stateXYPriorityQueue) Less(i, j int) bool {
	return this[i].distanceFromStart < this[j].distanceFromStart
}

func (this stateXYPriorityQueue) Len() int {
	return len(this)
}

func (this stateXYPriorityQueue) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *stateXYPriorityQueue) Push(x any) {
	*this = append(*this, x.(*stateXY))
}

func (this *stateXYPriorityQueue) Pop() any {
	x := (*this)[0]
	*this = (*this)[1:]
	return x
}

type stateFloat64 struct {
	id          int
	probability float64
}

/**
Priority Queue of `stateFloat64`
*/
type stateFloat64PriorityQueue []*stateFloat64

func (this stateFloat64PriorityQueue) Less(i, j int) bool {
	return this[i].probability > this[j].probability
}

func (this stateFloat64PriorityQueue) Len() int {
	return len(this)
}

func (this stateFloat64PriorityQueue) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *stateFloat64PriorityQueue) Push(x any) {
	*this = append(*this, x.(*stateFloat64))
}

func (this *stateFloat64PriorityQueue) Pop() any {
	length := this.Len()
	x := (*this)[length-1]
	*this = (*this)[:length-1]
	return x
}
