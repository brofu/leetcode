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
