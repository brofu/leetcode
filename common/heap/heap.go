package heap

import "fmt"

type IntHeap []int

func (this IntHeap) Len() int {
	return len(this)
}

func (this IntHeap) Less(i, j int) bool {
	return this[i] > this[j]
}

func (this IntHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this IntHeap) Top() int {
	if this.Len() > 0 {
		return this[0]
	}
	return 0
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// BinaryHeapInt  is full implmentation of binary heap on int
// No need the `capability` property
type BinaryHeapInt struct {
	data []int
}

func NewBinaryHeapInt() *BinaryHeapInt {
	return &BinaryHeapInt{
		data: []int{},
	}
}

func (this *BinaryHeapInt) Empty() bool {
	return len(this.data) == 0
}

func (this *BinaryHeapInt) Size() int {
	return len(this.data)
}

func (this *BinaryHeapInt) Push(val int) {
	this.data = append(this.data, val)
	this.swim(len(this.data) - 1)
}

func (this *BinaryHeapInt) Pop() int {
	this.swap(0, len(this.data)-1)
	val := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	this.sink(0)
	return val
}

func (this *BinaryHeapInt) Debug(msg string) {
	fmt.Println(msg, "size: ", len(this.data), "data:", this.data)
}

func (this *BinaryHeapInt) sink(index int) {
	smallest := index
	left, right := 2*index+1, 2*index+2

	if left < len(this.data) && this.data[smallest] > this.data[left] {
		smallest = left
	}
	if right < len(this.data) && this.data[smallest] > this.data[right] {
		smallest = right
	}

	if smallest != index {
		this.swap(index, smallest)
		this.sink(smallest)
	}
}

func (this *BinaryHeapInt) swim(index int) {
	parent := (index - 1) / 2
	if parent < 0 {
		return
	}
	if this.data[index] < this.data[parent] {
		this.swap(index, parent)
		this.swim(parent)
	}
}

func (this *BinaryHeapInt) swap(i, j int) {
	this.data[i], this.data[j] = this.data[j], this.data[i]
}

func NewInterfacePriorityQueue(lessFunc func([]interface{}, int, int) bool) *InterfacePriorityQueue {
	return &InterfacePriorityQueue{
		slice:    make([]interface{}, 0),
		lessFunc: lessFunc,
	}
}

type InterfacePriorityQueue struct {
	slice    []interface{}
	lessFunc func([]interface{}, int, int) bool
}

func (ipq *InterfacePriorityQueue) Len() int {
	return len(ipq.slice)
}

func (ipq *InterfacePriorityQueue) Swap(i, j int) {
	ipq.slice[i], ipq.slice[j] = ipq.slice[j], ipq.slice[i]
}

func (ipq *InterfacePriorityQueue) Less(i, j int) bool {
	return ipq.lessFunc(ipq.slice, i, j)
}

func (ipq *InterfacePriorityQueue) Push(x any) {
	ipq.slice = append(ipq.slice, x)
}

func (ipq *InterfacePriorityQueue) Pop() any {
	v := ipq.slice[len(ipq.slice)-1]
	ipq.slice = ipq.slice[:len(ipq.slice)-1]
	return v
}
