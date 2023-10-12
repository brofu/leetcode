package list

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
// 1. protect the situation: head is nil or n > length of the list
func FindFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return head
	}

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

// removeNthFromEnd
// Problem 19
// Note the constraints
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
