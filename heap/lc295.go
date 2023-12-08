package stream

import (
	"container/heap"
)

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type MedianFinder struct {
	maxHeap *cHeap
	minHeap *cHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: newCHeap(func(i, j int) bool { return i >= j }),
		minHeap: newCHeap(func(i, j int) bool { return i < j }),
	}
}

func (this MedianFinder) AddNum(num int) {

	maxHeapLength, minHeapLength := this.maxHeap.Len(), this.minHeap.Len()

	if maxHeapLength == 0 && minHeapLength == 0 {
		heap.Push(this.maxHeap, num)
		return
	}
	if maxHeapLength <= minHeapLength {
		if num <= this.minHeap.Top() {
			heap.Push(this.maxHeap, num)
		} else {
			heap.Push(this.maxHeap, heap.Pop(this.minHeap))
			heap.Push(this.minHeap, num)
		}
		return
	}

	if num >= this.maxHeap.Top() {
		heap.Push(this.minHeap, num)
	} else {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
		heap.Push(this.maxHeap, num)
	}

}

func (this MedianFinder) FindMedian() float64 {

	maxHeapLength, minHeapLength := this.maxHeap.Len(), this.minHeap.Len()

	if (maxHeapLength+minHeapLength)%2 == 0 {
		return float64(this.maxHeap.Top()+this.minHeap.Top()) / 2.0
	}
	return float64(this.maxHeap.Top())

}

func newCHeap(less func(int, int) bool) *cHeap {
	cHeap := &cHeap{}

	cHeap.less = func(i, j int) bool {
		i, j = cHeap.data[i], cHeap.data[j]
		return less(i, j)
	}

	return cHeap
}

type cHeap struct {
	data []int
	less func(int, int) bool
}

func (h *cHeap) Top() int {
	if h.Len() > 0 {
		return h.data[0]
	}
	return 0
}

func (h *cHeap) Push(num any) {
	h.data = append(h.data, num.(int))
}

func (h *cHeap) Pop() any {
	num := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return num
}

func (h *cHeap) Len() int {
	return len(h.data)
}

func (h *cHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *cHeap) Less(i, j int) bool {
	return h.less(i, j)
}
