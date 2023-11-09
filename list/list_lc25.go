package list

/*
Problem 25. Reverse Nodes in k-Group
Versions:
	*	Recursive
Key Points
	*	Recursive
		*	Break down the tasks
		*	Find the base case
		*	Find the middle pattern
*/

// Recursive
// T: O(N)
// M: O(N)
func reverseKGroup(head *ListNode, k int) *ListNode {

	p := head

	for index := 0; index < k; index++ {
		if p == nil { // base case: less than k nodes
			return head
		}
		p = p.Next
	}

	newHead := reverseByNode(head, p)
	head.Next = reverseKGroup(p, k)
	return newHead
}

// Key Points:
// 1.	target should one of nodes in the head. Add the `!= nil` to protect
// 2.	[head, b) is reversed
func reverseByNode(head *ListNode, target *ListNode) *ListNode {

	var pre, nxt *ListNode

	for head != target && head != nil {
		nxt = head.Next
		head.Next = pre
		pre = head
		head = nxt
	}

	return pre
}

func reverseByNodeWithTail(head *ListNode, target *ListNode) (*ListNode, *ListNode) {

	var pre, nxt *ListNode

	for head != target && head != nil {
		nxt = head.Next
		head.Next = pre
		pre = head
		head = nxt
	}

	return pre, nil
}

// Iterate
// T O(2N)
// M O(1)
func reverseKGroupV2(head *ListNode, k int) *ListNode {

	var (
		index      = 0
		p, s, tail = head, head, head
		newHead    *ListNode
	)

	for ; p != nil; index++ {
		p = p.Next
	}

	groups := index / k

	for index, p = 0, head; index < groups; index++ {
		for i := 0; i < k; i++ {
			p = p.Next
		}
		t := reverseByNode(s, p)
		if newHead == nil {
			newHead = t
		} else {
			tail.Next = t
		}
		tail = s
		s = p // adjust the start
	}

	tail.Next = p

	return newHead
}
