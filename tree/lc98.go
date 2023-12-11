package tree

import "math"

func isValidBST(root *TreeNode) bool {

	valid := true
	spaceHolder := int(1 - math.Pow(2, 31))

	var traverse func(*TreeNode) int

	traverse = func(root *TreeNode) int {

		if root == nil {
			return spaceHolder
		}

		if root.Left != nil {
			leftMax := traverse(root.Left)
			if leftMax != spaceHolder && root.Val <= leftMax {
				valid = false
				return
			}
		}
	}

	traverse(root)

	return valid
}
