package tree

/*
Sub Task

Key Point
  - Need to check if it's a skewed tree, if yes, put it on the left or the right. Both is OK
*/

func constructFromPrePostSubTaskPV1(preorder []int, postorder []int) *TreeNode {

	var traverse func([]int, []int) *TreeNode

	traverse = func(preorder, postorder []int) *TreeNode {
		if len(preorder) == 0 || len(postorder) == 0 || len(preorder) != len(postorder) {
			return nil
		}

		cur := preorder[0]
		root := &TreeNode{
			Val: cur,
		}

		// leaf node
		if len(preorder) == 1 {
			return root
		}

		leftVal := preorder[1]
		rightVal := postorder[len(postorder)-2]

		// only 1 child
		if leftVal == rightVal {
			root.Left = traverse(preorder[1:], postorder[:len(postorder)-1])
			// or put the right location. Both are ok
			// root.Right = traverse(preorder[1:], postorder[:len(postorder)-1])
			return root
		}

		// 2 child
		indexLeft := -1
		for idx, v := range postorder {
			if v == leftVal {
				indexLeft = idx
				break
			}
		}
		indexRight := -1
		for idx, v := range preorder {
			if v == rightVal {
				indexRight = idx
				break
			}
		}

		root.Left = traverse(preorder[1:indexRight], postorder[:indexLeft+1])
		root.Right = traverse(preorder[indexRight:], postorder[indexLeft+1:len(postorder)-1])

		return root
	}

	return traverse(preorder, postorder)
}
