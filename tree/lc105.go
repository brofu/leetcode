package tree

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return root
	}

	index := FindIndex(inorder, preorder[0]) // index should NOT equals to -1

	root.Left = buildTree(preorder[1:1+index], inorder[0:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return root
}
