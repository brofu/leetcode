package list

func ConstructListNodeFromSliceWithCycle(input []int, n int) *ListNode {

	if n > len(input)-1 {
		return nil
	}

	dummy := &ListNode{-1, nil}
	p := dummy
	t := dummy.Next

	index := 0

	for _, val := range input {
		node := &ListNode{val, nil}
		p.Next = node
		p = p.Next
		if index == n {
			t = node
		}
		index += 1
	}

	p.Next = t

	return dummy.Next
}

func ConstructListNodeFromSlice(input []int) *ListNode {
	dummy := &ListNode{-1, nil}
	p := dummy
	for _, val := range input {
		p.Next = &ListNode{val, nil}
		p = p.Next
	}
	return dummy.Next
}

func GetSliceFromList(list *ListNode) []int {
	s := []int{}
	for list != nil {
		s = append(s, list.Val)
		list = list.Next
	}

	return s
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists
// Problem 21
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := &ListNode{-1, nil}
	p := dummy
	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

// partition
// Problem 86
// Key Points:
// 1. Double pointer
// 2. Avoid point cycle
func partition(head *ListNode, x int) *ListNode {

	dummy := &ListNode{
		-1,
		nil,
	}
	dummy2 := &ListNode{
		-1,
		nil,
	}

	p := dummy
	s := dummy2

	for head != nil {
		if head.Val < x {
			p.Next = head
			p = p.Next
		} else {
			s.Next = head
			s = s.Next
		}
		head = head.Next

	}

	s.Next = nil
	p.Next = dummy2.Next

	return dummy.Next
}

// FindFromEnd
// Key Point
// Pay attention to the contraints
func FindFromEnd(head *ListNode, n int) *ListNode {

	p := head

	for i := 0; i < n; i += 1 {
		p = p.Next
	}

	s := head

	for p != nil {
		p = p.Next
		s = s.Next
	}

	return s
}

// FindFromEndWithProtection
// Key Point
// 1. protect the edges cases:
// 1.1 head is nil
// 1.2 n > length of the list
func FindFromEndWithProtection(head *ListNode, n int) *ListNode {

	// head is nil
	if head == nil {
		return head
	}

	p := head

	var (
		i int
		s *ListNode
	)

	// if n > length of head
	for ; i < n && p != nil; i += 1 {
		p = p.Next
	}

	// means n > length of head
	if i < n {
		return s
	}

	s = head
	for p != nil {
		p = p.Next
		s = s.Next
	}

	return s
}

/*
	Problem 19
	Note the constraints.
		*	input list as nil or
		*	n > list length
	Key Points
		*	Boundary control
			*	Dummy and the beginning
			*	nil at the end
		*	The idea of reusing the FindFromEnd
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{
		-1,
		head,
	}

	t := dummy
	p := head

	for i := 0; i < n; i += 1 {
		p = p.Next
	}

	s := head

	//  n == length of the list
	if p == nil {
		return s.Next
	}

	t.Next = s

	for p != nil {
		p = p.Next
		s = s.Next
		t = t.Next
	}

	t.Next = s.Next

	return dummy.Next
}

func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		-1,
		head,
	}
	x := FindFromEnd(dummy, n+1)

	x.Next = x.Next.Next

	return dummy.Next
}

func removeNthFromEndV3(head *ListNode, n int) *ListNode {

	dummy := &ListNode{
		-1,
		head,
	}

	t, p := dummy, head

	for i := 0; i < n; i += 1 {
		p = p.Next
	}

	for p != nil {
		p = p.Next
		t = t.Next
	}

	t.Next = t.Next.Next

	return dummy.Next
}

// middleNode
// Problem 876
// Key Points
// 1. Fast and slow pointer
func middleNode(head *ListNode) *ListNode {

	f := head
	s := head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}

	return s
}

/*
Problem 2095
Key Points
	*	Based on the constraints, should use 1 loop or 2 loop?
	*	2 or 3 pointer?
		*	Present the close-flowing notes as x.Next and x.Next.Next, so that, we can have only 1 pointer
	*	The constraints ==> when the nodes number is large, it's better to use more than 1 loop
		*	The number of nodes in the list is in the range [1, 105].
		*	1 <= Node.val <= 105
*/

func deleteMiddle(head *ListNode) *ListNode {

	// this should not happen with the constraints
	//if head == nil {
	//	return nil
	//}

	f := head
	s := head

	dummy := &ListNode{
		-1,
		nil,
	}
	t := dummy

	if f != nil && f.Next != nil {
		t.Next = s
	}

	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
		t = t.Next
	}

	t.Next = s.Next

	return dummy.Next
}

func deleteMiddleWithTwoLoop(head *ListNode) *ListNode {
	n := 0

	dummy := &ListNode{-1, head}
	s, p := dummy, head

	for ; p != nil; p = p.Next {
		n += 1
	}

	for i := 0; i < n/2; i++ {
		s = s.Next
	}

	s.Next = s.Next.Next

	return dummy.Next
}

func deleteMiddleV2(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	slow, fast := dummy, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

/*
Problem 142
Key Points
	*	Util Map for the check
	*	Multiple pointers
		*	The special idea
*/
func detectCycle(head *ListNode) *ListNode {

	hit := map[*ListNode]struct{}{}

	for ; head != nil; head = head.Next {
		if _, ok := hit[head]; !ok {
			hit[head] = struct{}{}
		} else {
			return head
		}
	}
	return nil
}

func detectCycleV2(head *ListNode) *ListNode {

	p, s := head, head

	for p != nil && p.Next != nil {
		p = p.Next.Next
		s = s.Next

		if s == p {
			break
		}
	}

	if p == nil || p.Next == nil {
		return nil
	}

	s = head
	for s != p {
		s = s.Next
		p = p.Next
	}

	return s
}
