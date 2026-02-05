package heap

// An implementataion to heap.Interface
type IntHeap []int

func NewHeap(n int) IntHeap {
	return make([]int, 0, n)
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	e := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return e
}

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Peek() int {
	if len(*h) == 0 {
		return -1
	}
	return (*h)[0]
}
