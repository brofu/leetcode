package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)

	return root
}

func invertTreeTraverse(root *TreeNode) *TreeNode {
	return invertTreeTraverseWrapper(root)
}

func invertTreeTraverseWrapper(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	// inorder
	root.Left, root.Right = root.Right, root.Left

	invertTreeTraverseWrapper(root.Left)
	invertTreeTraverseWrapper(root.Right)

	return root
}
