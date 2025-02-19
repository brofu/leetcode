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

/*
Sub Task
Construct left tree, and right tree, and the root
*/
func buildTreeSubTaskPV2(preorder []int, inorder []int) *TreeNode {

	var traverse func([]int, []int) *TreeNode

	traverse = func(preorder, inorder []int) *TreeNode {

		if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
			return nil
		}

		index := -1
		for idx, val := range inorder {
			if val == preorder[0] {
				index = idx
			}
		}

		root := &TreeNode{
			Val:   preorder[0],
			Left:  traverse(preorder[1:index+1], inorder[:index]),
			Right: traverse(preorder[index+1:], inorder[index+1:]),
		}

		return root
	}
	return traverse(preorder, inorder)
}
