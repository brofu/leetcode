package array

type MyCircularQueue struct {
	data             []int
	capability, size int
	start, end       int
}

func Constructor622(k int) MyCircularQueue {

	return MyCircularQueue{
		data:       make([]int, k, k),
		size:       0,
		capability: k,
		start:      0,
		end:        0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.end] = value
	this.end = (this.end + 1) % this.capability
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.start = (this.start + 1) % this.capability
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.start]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.end-1+this.capability)%this.capability]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capability
}
