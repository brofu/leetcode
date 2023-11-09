package tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return MaxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
