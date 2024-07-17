package tree

func buildTreePV1(preorder []int, inorder []int) *TreeNode {

	var traverse func([]int, []int) *TreeNode

	traverse = func(preorder, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}

		val := preorder[0]
		node := &TreeNode{Val: val}

		var i int
		for i = 0; i < len(inorder) && inorder[i] != val; i++ {
		}
		node.Left = traverse(preorder[1:i+1], inorder[0:i])
		node.Right = traverse(preorder[i+1:], inorder[i+1:])
		return node
	}

	node := traverse(preorder, inorder)

	return node
}
