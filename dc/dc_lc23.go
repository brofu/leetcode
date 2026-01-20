package dc

import "github.com/brofu/leetcode/common"

type ListNode = common.ListNode

/*
KP:
	1. 为什么有些算法可以通过分治思想来优化时间复杂度呢？
	2. 把递归算法抽象成递归树，如果递归树节点的时间复杂度和树的深度相关，那么使用分治思想对问题进行二分，就可以使递归树尽可能平衡，进而优化总的时间复杂度。
	3. 反之，如果递归树节点的时间复杂度和树的深度无关，那么使用分治思想就没有意义，反而可能引入额外的空间复杂度。

TC:
	1. O(N*lgK). K is the number of list, and N is the total number of nodes in K lists
	2. The recursive tree depth is lgK, and on each layer, the nodes are always N.


SC:
	1. Recursive depth O(lgK)
	2. No extra space if not considering dumy, p1, p2, p, and the result


Think About:
	1. Why iteration would have higher time complexity?
	2. Merge 1 and 2, get, 1-2, and then, 1-2 and 3, get 1-2-3 and then...
	3. The total merge time is SAME, say that, k-1
	4. But, the earlier the list merged, the more times it would be merged repeatedly. For example, if k=8, then, list 1 and 2 would be merged 7 times, list 3 would be merged 7 times.
	5. It's N/K*2(K-1) + N/K(K-2) + N/K(K-3) + ... + N/K(K-1)

*/

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}

	return mergeKListsRecursive(lists, 0, len(lists)-1)
}

func mergeKListsRecursive(list []*ListNode, start, end int) *ListNode {

	if start == end {
		return list[start]
	}

	mid := (end-start)/2 + start
	left := mergeKListsRecursive(list, start, mid)
	right := mergeKListsRecursive(list, mid+1, end)

	return mergeTwoLists(left, right)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {

	dumy := &ListNode{
		Next: nil,
	}
	p := dumy
	p1, p2 := list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = (*ListNode)(p1.Next)
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

	return dumy.Next
}
