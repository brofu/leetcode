package tree

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	}

	return root
}
