package tree

import "math"

// O(n)
func countNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// if a perfect binary tree
// O(logN)
func countNodesPerfectBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	h := 0
	for ; root != nil; root = root.Left {
		h += 1
	}
	return int(math.Pow(2, float64(h))) - 1
}

// O(lgN*lgN)
func countNodesCompleteBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left, right := 0, 0

	for lt := root; lt != nil; lt = lt.Left {
		left += 1
	}

	for rt := root; rt != nil; rt = rt.Right {
		right += 1
	}

	if left == right {
		return int(math.Pow(2, float64(left))) - 1
	}

	return 1 + countNodesCompleteBinaryTree(root.Left) + countNodesCompleteBinaryTree(root.Right)
}
