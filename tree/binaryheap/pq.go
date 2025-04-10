package binaryheap

import "fmt"

type PriorityQueue interface {
	Push(interface{})
	Pop() interface{}
	Size() int
	Debug()
}

/*
Key Points:
1. `Array` based solution requires a `Complete Tree`
2. When `Push`, append the value to the `end` of the array (to match the requirement of `Complete Tree`), and do `swim` operation.
3. When `Pop`, return the `head` value of the array, and move the `end` of the array to the `head` location, and do `sink` operation
*/

type PQBasedOnBinaryHeap struct {
	slice        []interface{}
	size         int
	lessThanFunc func([]interface{}, int, int) bool
}

func NewPQBasedOnBH(size int, f func([]interface{}, int, int) bool) *PQBasedOnBinaryHeap {
	pq := &PQBasedOnBinaryHeap{
		size: 0,
	}
	pq.slice = make([]interface{}, size+1, size+1)
	pq.lessThanFunc = f
	return pq
}

func (pq *PQBasedOnBinaryHeap) Size() int {
	return pq.size
}

func (pq *PQBasedOnBinaryHeap) Peak() interface{} {
	if pq.size == 0 {
		return nil
	}
	return pq.slice[1]
}

func (pq *PQBasedOnBinaryHeap) Push(val interface{}) {
	if pq.size == len(pq.slice) {
		return
	}
	pq.size++
	pq.slice[pq.size] = val

	pq.swim(pq.size)
}

func (pq *PQBasedOnBinaryHeap) Pop() interface{} {

	if pq.size == 0 {
		return nil
	}

	val := pq.slice[1]

	pq.slice[1] = pq.slice[pq.size]
	pq.slice[pq.size] = nil
	pq.size--
	pq.sink(1)

	return val
}

func (pq *PQBasedOnBinaryHeap) sink(index int) {

	leftIndex, rightIndex := pq.left(index), pq.right(index)
	for leftIndex <= pq.size && rightIndex <= pq.size {
		minIndex := leftIndex
		if pq.lessThanFunc(pq.slice, rightIndex, leftIndex) {
			minIndex = rightIndex
		}
		if pq.lessThanFunc(pq.slice, index, minIndex) {
			break
		}
		pq.swap(index, minIndex)
		index = minIndex
		leftIndex = pq.left(index)
		rightIndex = pq.right(index)
	}
	if leftIndex <= pq.size {
		if pq.lessThanFunc(pq.slice, leftIndex, index) {
			pq.swap(leftIndex, index)
		}
	}
}

func (pq *PQBasedOnBinaryHeap) swim(index int) {
	for parent := pq.parent(index); parent > 0 && pq.lessThanFunc(pq.slice, index, parent); {
		pq.swap(index, parent)
		index = parent
		parent = pq.parent(index)
	}
}

func (pq *PQBasedOnBinaryHeap) parent(index int) int {
	return index / 2
}

func (pq *PQBasedOnBinaryHeap) left(index int) int {
	return index * 2
}

func (pq *PQBasedOnBinaryHeap) right(index int) int {
	return index*2 + 1
}

func (pq *PQBasedOnBinaryHeap) swap(i, j int) {
	pq.slice[i], pq.slice[j] = pq.slice[j], pq.slice[i]
}

func (pq *PQBasedOnBinaryHeap) Debug() {
	fmt.Println(pq.slice)
}
