package tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else { // assume there is NO `equal`
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
