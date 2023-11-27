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

func flattenV2(root **TreeNode) {
	if root == nil {
		return
	}

	dummy := &TreeNode{
		Val:   -1,
		Left:  nil,
		Right: nil,
	}

	flattenTraverseV2(*root, dummy)
	*root = dummy.Right
}

func flattenTraverseV2(root, dummy *TreeNode) *TreeNode {

	dummy.Right = &TreeNode{
		Val: root.Val,
	}

	dummy = dummy.Right

	if root.Left != nil {
		dummy = flattenTraverseV2(root.Left, dummy)
	}
	if root.Right != nil {
		dummy = flattenTraverseV2(root.Right, dummy)
	}

	return dummy
}
