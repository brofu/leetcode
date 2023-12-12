package tree

func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Left == nil {
		return root.Right
	}

	if root.Right == nil {
		return root.Left
	}

	root.Val = GetMinFromBST(root.Right)
	root.Right = deleteNode(root.Right, root.Val)

	return root
}
