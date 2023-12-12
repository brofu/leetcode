package tree

import "math"

func isValidBST(root *TreeNode) bool {

	valid := true
	spaceHolder := int(1 - math.Pow(2, 31))

	var traverse func(*TreeNode) (int, int)

	traverse = func(root *TreeNode) (min, max int) {

		if root == nil || !valid {
			return spaceHolder, spaceHolder
		}

		min, max = root.Val, root.Val
		leftMax, rightMin := spaceHolder, spaceHolder

		if root.Left != nil {
			min, leftMax = traverse(root.Left)
			if leftMax != spaceHolder && root.Val <= leftMax {
				valid = false
				return
			}
		}

		if root.Right != nil {
			rightMin, max = traverse(root.Right)
			if rightMin != spaceHolder && root.Val >= rightMin {
				valid = false
				return
			}
		}
		return
	}

	_, _ = traverse(root)

	return valid

}

// better one
func isValidBSTV2(root *TreeNode) bool {

	var traverse func(*TreeNode, *int, *int) bool

	traverse = func(root *TreeNode, min, max *int) bool {
		if root == nil {
			return true
		}

		if min != nil && root.Val <= *min {
			return false
		}

		if max != nil && root.Val >= *max {
			return false
		}

		return traverse(root.Left, min, &root.Val) && traverse(root.Right, &root.Val, max)
	}

	return traverse(root, nil, nil)
}
