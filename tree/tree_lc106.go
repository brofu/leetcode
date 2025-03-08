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

/*
Sub Task
Use map to reduce the time complexity
*/
func buildTreeSubTaskPV3(inorder []int, postorder []int) *TreeNode {

	var traverse func([]int, []int) *TreeNode

	traverse = func(inorder, postorder []int) *TreeNode {
		if len(postorder) == 0 {
			return nil
		}

		root := &TreeNode{
			Val: postorder[len(postorder)-1],
		}
		index := -1
		for idx, num := range inorder {
			if num == root.Val {
				index = idx
				break
			}
		}
		root.Left = traverse(inorder[:index], postorder[:index])
		root.Right = traverse(inorder[index+1:], postorder[index:len(postorder)-1])
		return root
	}
	return traverse(inorder, postorder)
}
