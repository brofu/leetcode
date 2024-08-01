package array

import "github.com/brofu/leetcode/common"

func deleteDuplicates(head *common.ListNode) *common.ListNode {

	p, q := head, head

	for q != nil {
		if p.Val != q.Val {
			p.Next = q
			p = p.Next
		}
		q = q.Next
	}

	if p != nil { // for the case of head == nil
		p.Next = nil // need this
	}

	return head
}
