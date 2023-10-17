package list

/*
Problem 141
Key Points
	*	Multiple pointer
*/

func hasCycle(head *ListNode) bool {

	p, q := head, head

	for p != nil && p.Next != nil && p.Next.Next != nil {
		p = p.Next.Next
		q = q.Next

		if p == q {
			return true
		}
	}

	return false

}
