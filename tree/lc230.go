package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Assume k <= nodes number of root
func kthSmallest(root *TreeNode, k int) int {
	if root != nil {
		_, val := kthSmallestWrapper(root, k)
		return val
	}
	return -1
}

func kthSmallestWrapper(root *TreeNode, index int) (int, int) {

	if root.Left == nil && root.Right == nil { // leaf node
		return index - 1, root.Val
	}

	val := -1
	if root.Left != nil {
		index, val = kthSmallestWrapper(root.Left, index)
	}

	// get the target val from left tree
	if index == 0 {
		return index, val
	}

	// get the target for root node itself
	if index == 1 {
		return 0, root.Val
	}

	// reduce the index
	index -= 1

	if root.Right == nil {
		return index, -1
	}

	// return from right subtree
	return kthSmallestWrapper(root.Right, index)
}

func kthSmallestV2(root *TreeNode, k int) int {

	val, index := -1, 0

	traverse := func(*TreeNode) {}
	traverse = func(root *TreeNode) {

		if root == nil {
			return
		}

		traverse(root.Left)

		index += 1
		if index == k {
			val = root.Val
			return
		}

		traverse(root.Right)
	}

	traverse(root)

	return val
}

// Assume k <= nodes number of root
func kthSmallestV3(root *TreeNodeWithSubtreeSum, k int) int {
	if root == nil {
		return 0
	}

	leftCount := Count(root.Left)
	if k <= leftCount {
		return kthSmallestV3(root.Left, k)
	} else if k > leftCount+1 {
		return kthSmallestV3(root.Right, k-leftCount-1)
	}

	return root.Val
}

func Insert(root *TreeNodeWithSubtreeSum, val int) *TreeNodeWithSubtreeSum {

	if root == nil {
		root = &TreeNodeWithSubtreeSum{
			Val: val,
			Sum: 1,
		}
		return root
	}

	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else {
		root.Right = Insert(root.Right, val)
	}

	root.Sum = 1 + Count(root.Left) + Count(root.Right)

	return root
}

func Delete(root *TreeNodeWithSubtreeSum, val int) *TreeNodeWithSubtreeSum {

	if root == nil {
		return nil
	}

	if val < root.Val {
		root.Left = Delete(root.Left, val)
	}
	if val > root.Val {
		root.Right = Delete(root.Right, val)
	}
	if val == root.Val {
		if root.Left == nil { // only has right tree
			return root.Right
		}
		if root.Right == nil { // only has left tree
			return root.Left
		}
		root.Val = Min(root.Right)
		root.Right = Delete(root.Right, root.Val)
	}

	root.Sum = 1 + Count(root.Left) + Count(root.Right)
	return root
}

func Count(root *TreeNodeWithSubtreeSum) int {
	if root == nil {
		return 0
	}
	return root.Sum
}

func Min(root *TreeNodeWithSubtreeSum) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}
