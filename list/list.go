package list

import (
	"errors"

	"github.com/brofu/leetcode/common"
	"github.com/brofu/leetcode/tree/tree_common"
)

func ConstructIntersectedListNodesFromSlice(sliceA, sliceB []int, skipA, skipB int) (*ListNode, *ListNode, error) {

	var (
		lengthA = len(sliceA)
		lengthB = len(sliceB)

		dummy1 = &ListNode{
			-1,
			nil,
		}
		dummy2 = &ListNode{
			-1,
			nil,
		}
		p1 = dummy1
		p2 = dummy2
	)

	if skipA < 0 || skipA > lengthA || skipB < 0 || skipB > lengthB {
		return nil, nil, errors.New("wrong skip parameters")
	}

	i := 0
	for ; i < skipA; i += 1 {
		p1.Next = &ListNode{
			sliceA[i],
			nil,
		}
		p1 = p1.Next
	}

	intersectedNode := p1

	for ; i < lengthA; i += 1 {
		p1.Next = &ListNode{
			sliceA[i],
			nil,
		}
		p1 = p1.Next
	}

	for j := 0; j < skipB; j += 1 {
		p2.Next = &ListNode{
			sliceB[j],
			nil,
		}
		p2 = p2.Next
	}

	p2.Next = intersectedNode.Next

	return dummy1.Next, dummy2.Next, nil
}

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

/*
Problem 160 Intersection of Two Linked Lists
Key Points
	*	Performance
		*	Map with memory O(M+N) or
		*	Multiple pointer with time O(M+N) and memory O(1)
	*	Implement
		*	More clear implementation
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	n, n1, n2 := 0, 0, 0

	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	for p1 != nil {
		p1 = p1.Next
		n1 += 1
	}
	for p2 != nil {
		p2 = p2.Next
		n2 += 1
	}

	if n1 > 0 {
		n = n1
		p1 = headA
		p2 = headB
	} else {
		n = n2
		p1 = headB
		p2 = headA
	}

	for i := 0; i < n; i += 1 {
		p1 = p1.Next
	}

	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil
}

func getIntersectionNodeV2(headA, headB *ListNode) *ListNode {

	pointerAFaster, pointerASlower := headA, headA
	pointerBFaster, pointerBSlower := headB, headB

	for pointerAFaster != nil && pointerBFaster != nil {
		pointerAFaster = pointerAFaster.Next
		pointerBFaster = pointerBFaster.Next
	}

	for pointerAFaster != nil {
		pointerAFaster = pointerAFaster.Next
		pointerASlower = pointerASlower.Next
	}

	for pointerBFaster != nil {
		pointerBFaster = pointerBFaster.Next
		pointerBSlower = pointerBSlower.Next
	}

	for pointerASlower != nil && pointerBSlower != nil {
		if pointerASlower == pointerBSlower {
			return pointerASlower
		}
		pointerASlower = pointerASlower.Next
		pointerBSlower = pointerBSlower.Next
	}

	return nil
}

/*
Problem 23. Merge k Sorted Lists
Versions
	*	V3 merge the list by 2-2. The time consumption would be higher than version 1 and 2
Key Points
	*	Use Binary Heap for sort
*/

// CompareTo Let *ListNode implement the Comparable interface
// But it's not allowed by leetcode
func (ln *ListNode) CompareTo(destination common.Comparable) int {
	destinationListNode, _ := destination.(*ListNode)
	if ln.Val > destinationListNode.Val {
		return 1
	}
	if ln.Val == destinationListNode.Val {
		return 0
	}
	return -1
}

func mergeKLists(lists []*ListNode) *ListNode {

	dummy := &ListNode{
		-1,
		nil,
	}

	p := dummy

	pq := NewMinListNodePQ(lists)

	node := pq.Pop()
	for node != nil {
		p.Next = node
		p = p.Next
		if node.Next != nil {
			pq.Push(node.Next)
		}
		node = pq.Pop()
	}

	return dummy.Next
}

type MinListNodePQ struct {
	binaryHeap *tree_common.BinaryHeap
}

func NewMinListNodePQ(nodes []*ListNode) *MinListNodePQ {
	pq := &MinListNodePQ{}

	//TODO: better approach
	elements := make([]common.Comparable, len(nodes))
	for i, node := range nodes {
		elements[i] = node
	}
	hb := tree_common.NewBinaryHeap(elements, pq.greaterThan)
	pq.binaryHeap = hb
	return pq
}
func (pq *MinListNodePQ) Push(node *ListNode) {
	pq.binaryHeap.Push(node)
}

func (pq *MinListNodePQ) Pop() (node *ListNode) {
	element := pq.binaryHeap.Pop()
	node, _ = element.(*ListNode)
	return
}

func (pq *MinListNodePQ) greaterThan(source, destination common.Comparable) bool {
	return source.CompareTo(destination) > 0
}

func mergeKListsV2(lists []*ListNode) *ListNode {

	dummy := &ListNode{
		-1,
		nil,
	}

	p := dummy

	pq := NewMinListNodePQV2(lists)

	node := pq.Pop()
	for node != nil {
		p.Next = node
		p = p.Next
		if node.Next != nil {
			pq.Push(node.Next)
		}
		node = pq.Pop()
	}

	return dummy.Next
}

type MinListNodePQV2 struct {
	array []*ListNode
	size  int
}

func NewMinListNodePQV2(lists []*ListNode) *MinListNodePQV2 {
	pq := &MinListNodePQV2{
		array: make([]*ListNode, len(lists)+1),
		size:  0,
	}
	for _, list := range lists {
		if list != nil {
			pq.Push(list)
		}
	}
	return pq
}

func (pq *MinListNodePQV2) Push(list *ListNode) {
	pq.size += 1
	pq.array[pq.size] = list
	pq.swim(pq.size)
}

func (pq *MinListNodePQV2) swim(index int) {
	for index > 1 && pq.greaterThan(pq.parent(index), index) {
		pq.swap(pq.parent(index), index)
		index = pq.parent(index)
	}
}

func (pq *MinListNodePQV2) swap(sourceIndex, destinationIndex int) {
	pq.array[sourceIndex], pq.array[destinationIndex] = pq.array[destinationIndex], pq.array[sourceIndex]
}

func (pq *MinListNodePQV2) parent(index int) int {
	return index / 2
}

func (pq *MinListNodePQV2) greaterThan(sourceIndex, destinationIndex int) bool {
	return pq.array[sourceIndex].Val > pq.array[destinationIndex].Val
}

func (pq *MinListNodePQV2) Pop() *ListNode {

	if pq.size == 0 {
		return nil
	}

	list := pq.array[1]

	pq.swap(1, pq.size)
	pq.array[pq.size] = nil
	pq.size -= 1

	pq.sink(1)
	return list
}

func (pq *MinListNodePQV2) sink(index int) {

	for pq.left(index) <= pq.size { // there is left child
		minIndex := pq.left(index)
		rightIndex := pq.right(index)
		if rightIndex <= pq.size && pq.greaterThan(minIndex, rightIndex) { // there is right child and compare
			minIndex = rightIndex
		}

		if !pq.greaterThan(index, minIndex) {
			break
		}
		pq.swap(index, minIndex)
		index = minIndex
	}
}

func (pq *MinListNodePQV2) left(index int) int {
	return index * 2
}

func (pq *MinListNodePQV2) right(index int) int {
	return index*2 + 1
}

func mergeKListsV3(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2

	left := mergeKListsV3(lists[:mid])
	right := mergeKListsV3(lists[mid:])

	return MergeTwoLists(left, right)
}
