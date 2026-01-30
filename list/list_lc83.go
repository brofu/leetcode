package list

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	slow, fast := head, head

	for fast != nil {
		if slow.Val == fast.Val {
			fast = fast.Next
			continue
		}
		slow.Next = fast
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = fast

	return head
}
