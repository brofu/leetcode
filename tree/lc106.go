package tree

func buildTreeV2(inorder []int, postorder []int) *TreeNode {

	if len(inorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]

	root := &TreeNode{
		Val: val,
	}

	if len(inorder) == 1 {
		return root
	}

	index := FindIndex(inorder, val)
	root.Left = buildTreeV2(inorder[0:index], postorder[0:index])
	root.Right = buildTreeV2(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
