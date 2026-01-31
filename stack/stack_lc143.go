package stack

import (
	"github.com/brofu/leetcode/list"
)

type ListNode = list.ListNode

func reorderList(head *ListNode) {

	if head == nil {
		return
	}

	st := []*ListNode{}

	p := head
	for p != nil {
		st = append(st, p)
		p = p.Next
	}

	p = head

	for i, j := 0, len(st)-1; i <= j; i, j = i+1, j-1 {
		next := st[j]
		next.Next = p.Next
		p.Next = next
		p = p.Next.Next
	}
	p.Next = nil
}
