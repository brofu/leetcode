package list

/*
Problem 206
Key Point
Versions
	*	Iterate version 1. Bad performance
	*	Iterate version 2. Multiple pointer and the controlling logic
	*	Recursive version 3
	*	Recursive version 4. Key point is, how to use 1 pointer, instead of 2?

Extensions
	*	Reverse the first N nodes.
		*	return 2 parameters or use global var
		*	The edges cases
			*	n > list length
*/

// Iterate verion 1.
func reverseList(head *ListNode) *ListNode {

	nodes := []*ListNode{}
	dummy := &ListNode{
		-1,
		nil,
	}
	p := dummy

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	for index := len(nodes) - 1; index >= 0; index -= 1 {
		p.Next = nodes[index]
		p = p.Next
	}

	p.Next = nil

	return dummy.Next
}

// Iterate verion 2.
func reverseListV2(head *ListNode) *ListNode {
	var before, after *ListNode
	for head != nil {
		after = head.Next
		head.Next = before

		before = head
		head = after
	}
	return before
}

// Recursive version 3
func reverseListV3(head *ListNode) *ListNode {

	if head != nil && head.Next != nil {
		newHead, _ := reverseListHelper(head, head.Next)
		return newHead
	}

	return head
}

func reverseListHelper(current, next *ListNode) (*ListNode, *ListNode) {

	current.Next = nil
	if next.Next != nil {
		head, tail := reverseListHelper(next, next.Next)
		tail.Next = current
		return head, current
	}
	next.Next = current
	return next, current
}

// Recursive version 3, better implementation. More briefing.

func reverseListV4(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	last := reverseListV4(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// reverse version
func reverseListFirstN(head *ListNode, n int) *ListNode {
	last, _ := reverseListFirstNHelper(head, n)
	return last
}

func reverseListFirstNHelper(head *ListNode, n int) (*ListNode, *ListNode) {
	if head == nil {
		return head, nil
	}

	if n <= 1 || head.Next == nil {
		return head, head.Next
	}

	last, tail := reverseListFirstNHelper(head.Next, n-1)
	head.Next.Next = head
	head.Next = tail

	return last, tail
}

func reverseListBetween(head *ListNode, m, n int) *ListNode {

	if m > n {
		return nil
	}

	dummy := &ListNode{
		-1,
		head,
	}

	p := dummy

	for index := 0; index < m-1 && p != nil; index++ {
		p = p.Next
	}

	if p == nil || p.Next == nil {
		return dummy.Next
	}

	last := reverseListFirstN(p.Next, n-m+1)

	p.Next = last

	return dummy.Next
}

// v2.totally reverse
func reverseListBetweenV2(head *ListNode, m, n int) *ListNode {

	if m > n {
		return nil
	}

	if m == 1 {
		return reverseListFirstN(head, n)
	}

	head.Next = reverseListBetweenV2(head.Next, m-1, n-1)

	return head
}
