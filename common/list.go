package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateListNodeFromSlice(input []int) *ListNode {

	dummy := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := dummy
	for _, s := range input {
		node := &ListNode{
			Val:  s,
			Next: nil,
		}
		p.Next = node
		p = p.Next
	}

	return dummy.Next
}

func GenerateSliceFromListNode(node *ListNode) []int {
	result := []int{}
	for ; node != nil; node = node.Next {
		result = append(result, node.Val)
	}
	return result
}
