package tree_common

import "github.com/brofu/leetcode/common"

type checkSwapFunc func(source, destination common.Comparable) bool

type BinaryHeap struct {
	array         []common.Comparable
	size          int
	checkSwapFunc checkSwapFunc
}

// NewBinaryHeap set up a new Binary Heap object
// Need to init
func NewBinaryHeap(elements []common.Comparable, checkSwapFunc checkSwapFunc) *BinaryHeap {

	length := len(elements)

	bh := &BinaryHeap{
		// Assume the total number of the queue <= len(elements) when initial
		array:         make([]common.Comparable, length+1, length+1),
		size:          0,
		checkSwapFunc: checkSwapFunc,
	}

	for _, element := range elements {
		bh.Push(element)
	}

	return bh
}

func (bh *BinaryHeap) parent(index int) int {
	return index / 2
}

func (bh *BinaryHeap) left(index int) int {
	return index * 2
}

func (bh *BinaryHeap) right(index int) int {
	return index*2 + 1
}

func (bh *BinaryHeap) swap(indexA, indexB int) {
	bh.array[indexA], bh.array[indexB] = bh.array[indexB], bh.array[indexA]
}

// Push pushes a new element into the queue
// Assume the total number of the queue <= len(elements) when initial
func (bh *BinaryHeap) Push(element common.Comparable) {
	bh.size += 1
	bh.array[bh.size] = element
	bh.swim(bh.size)
}
func (bh *BinaryHeap) Pop() common.Comparable {
	if bh.size == 0 {
		return nil
	}
	element := bh.array[1]

	bh.swap(1, bh.size)
	bh.array[bh.size] = nil
	bh.size -= 1

	bh.sink(1)
	return element
}

func (bh *BinaryHeap) swim(index int) {
	for index > 1 {
		parent := bh.array[bh.parent(index)]
		current := bh.array[index]
		if bh.checkSwapFunc(parent, current) {
			bh.swap(bh.parent(index), index)
			index = bh.parent(index)
		}
	}
}

func (bh *BinaryHeap) sink(index int) {
	for bh.left(index) <= bh.size { // left node exists

		maxIndex := bh.left(index)

		if bh.right(index) <= bh.size {
			leftMax := bh.array[maxIndex]
			rightMax := bh.array[bh.right(index)]
			if bh.checkSwapFunc(leftMax, rightMax) {
				maxIndex = bh.right(index)
			}
		}

		current := bh.array[index]
		max := bh.array[maxIndex]

		if !bh.checkSwapFunc(current, max) {
			break
		}

		bh.swap(index, maxIndex)
		index = maxIndex
	}
}

//MaxPQ is Max Priority Queue based on binary heap
type MaxPQ struct {
	*BinaryHeap
}

func NewMaxPQ(elements []common.Comparable) *MaxPQ {

	pq := &MaxPQ{}
	bh := NewBinaryHeap(elements, pq.lessThan)
	pq.BinaryHeap = bh

	return pq
}

// Queue return the internal queue
// For testing purpose
func (pq *MaxPQ) Queue() []common.Comparable {
	return pq.BinaryHeap.array
}

func (pq *MaxPQ) Push(element common.Comparable) {
	pq.BinaryHeap.Push(element)
}

func (pq *MaxPQ) Pop() common.Comparable {
	return pq.BinaryHeap.Pop()
}

func (pq *MaxPQ) lessThan(source, destination common.Comparable) bool {
	return source.CompareTo(destination) < 0
}
