package tree

func isCousins(root *TreeNode, x int, y int) bool {

	var helper func(*TreeNode, int, int) (*TreeNode, int)

	helper = func(root *TreeNode, target, height int) (*TreeNode, int) {

		if root == nil {
			return nil, 0
		}

		if root.Left != nil && root.Left.Val == target {
			return root, height + 1
		}

		if root.Right != nil && root.Right.Val == target {
			return root, height + 1
		}

		if parent, h := helper(root.Left, target, height+1); parent != nil {
			return parent, h
		}

		return helper(root.Right, target, height+1)
	}

	parentX, heightX := helper(root, x, 0)
	parentY, heightY := helper(root, y, 0)
	return parentX != parentY && heightX == heightY
}
