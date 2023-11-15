package tree

func flatten(root *TreeNode) {
	_ = flattenTraverse(root)
}

//flattenTraverse post order
func flattenTraverse(root *TreeNode) *TreeNode {

	if root == nil || root.Left == nil && root.Right == nil { // leaf node
		return root
	}

	// traverse left
	leftLeaf := flattenTraverse(root.Left)

	// traverse right
	rightLeaf := flattenTraverse(root.Right)

	// current node
	if leftLeaf != nil {
		leftLeaf.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	if rightLeaf == nil {
		return leftLeaf
	} else {
		return rightLeaf
	}
}
