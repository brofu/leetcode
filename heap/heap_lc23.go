package heap

import (
	"github.com/brofu/leetcode/common"
)

func mergeKLists(lists []*common.ListNode) *common.ListNode {

	dummy := &common.ListNode{
		Val:  -1,
		Next: nil,
	}

	p := dummy

	bHeap := NewListNodeBHeap(lists)

	for node := bHeap.Pop(); node != nil; node = bHeap.Pop() {
		p.Next = node
		p = p.Next
		if node.Next != nil {
			bHeap.Push(node.Next)
		}
	}

	return dummy.Next
}
