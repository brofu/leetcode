package tree

func buildTreeSubTaskPV1(inorder []int, postorder []int) *TreeNode {

	var traverse func([]int, []int) *TreeNode

	traverse = func(inorder, postorder []int) *TreeNode {
		if len(inorder) == 0 || len(postorder) == 0 || len(inorder) != len(postorder) {
			return nil
		}

		val := postorder[len(postorder)-1]
		index := -1
		for idx, v := range inorder {
			if v == val {
				index = idx
				break
			}
		}

		root := &TreeNode{
			Val:   val,
			Left:  traverse(inorder[:index], postorder[:index]),
			Right: traverse(inorder[index+1:], postorder[index:len(postorder)-1]),
		}

		return root
	}
	return traverse(inorder, postorder)
}
