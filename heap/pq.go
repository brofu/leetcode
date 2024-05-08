package heap

import (
	"fmt"

	"github.com/brofu/leetcode/common"
)

/**
A min binary heap
*/

type ListNodeBHeap struct {
	size  int
	nodes []*common.ListNode
}

func NewListNodeBHeap(lists []*common.ListNode) *ListNodeBHeap {
	h := &ListNodeBHeap{
		size:  0,
		nodes: make([]*common.ListNode, len(lists)+1), // why len(lists) + 1? easy to caculate parenet
	}

	for _, node := range lists {
		if node != nil {
			h.Push(node)
		}
	}

	return h
}

func (this *ListNodeBHeap) Debug(msg string) {
	for i, s := range this.nodes {
		if i == 0 {
			continue
		}
		if s != nil {
			fmt.Println(msg, i, s.Val)
		}
	}
}
func (this *ListNodeBHeap) Pop() *common.ListNode {

	if this.size == 0 {
		return nil
	}

	node := this.nodes[1]
	this.swap(1, this.size)
	this.nodes[this.size] = nil
	this.size -= 1
	this.sink(1)

	return node
}

func (this *ListNodeBHeap) Push(node *common.ListNode) {
	this.size += 1
	this.nodes[this.size] = node
	this.swim(this.size)
}

func (this *ListNodeBHeap) swim(index int) {

	if index == 1 {
		return
	}

	parent := index / 2

	if this.getValue(parent) > this.getValue(index) {
		this.swap(parent, index)
		this.swim(parent)
	}
}

func (this *ListNodeBHeap) sink(index int) {
	smaller := index
	left := 2 * index
	right := 2*index + 1

	if left <= this.size && this.getValue(left) < this.getValue(smaller) {
		smaller = left
	}
	if right <= this.size && this.getValue(right) < this.getValue(smaller) {
		smaller = right
	}

	if smaller != index { // other the left child or the right child is smaller than index
		this.swap(index, smaller)
		this.sink(smaller)
	}
}

func (this *ListNodeBHeap) getValue(index int) int {
	return this.nodes[index].Val
}

func (this *ListNodeBHeap) swap(src, dest int) {
	this.nodes[src], this.nodes[dest] = this.nodes[dest], this.nodes[src]
}
