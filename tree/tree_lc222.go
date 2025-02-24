package tree

import "math"

func countNodesPV1(root *TreeNode) int {

	var traverse func(*TreeNode) int

	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := 1
		leftC := root.Left
		for leftC != nil {
			left++
			leftC = leftC.Left
		}

		right := 1
		rightC := root.Right
		for rightC != nil {
			right++
			rightC = rightC.Right
		}

		if left == right {
			return int(math.Pow(float64(2), float64(left))) - 1
		}

		return traverse(root.Left) + traverse(root.Right) + 1
	}
	return traverse(root)
}
