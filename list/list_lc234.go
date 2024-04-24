package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * 1. Find the half of the list
 * 2. Reverse the 2nd half
 * 3. Compare

 * O(n) + O(1)
 */
func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// if fast == nil, reverse from `slow`
	// else reverse from `slow.Next`
	if fast != nil {
		slow = slow.Next
	}

	head2 := reverseListV2(slow)

	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}
