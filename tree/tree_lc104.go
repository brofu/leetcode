package tree

import "github.com/brofu/leetcode/common"

/*
二叉树题目的递归解法可以分两类思路，第一类是遍历一遍二叉树得出答案，第二类是通过分解问题计算出答案，
这两类思路分别对应着回溯算法核心框架 和 动态规划核心框架。

traverse the tree to get the answer
*/
func maxDepthTraverseP1(root *TreeNode) int {

	maxDepth := 0
	depth := 0

	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {

		if root == nil {
			return
		}

		// pre-order
		depth++
		if root.Left == nil && root.Right == nil {
			maxDepth = common.MaxInt(maxDepth, depth)
		}

		traverse(root.Left)
		traverse(root.Right)

		// post-order
		depth--
	}

	traverse(root)
	return maxDepth
}

/*
二叉树题目的递归解法可以分两类思路，第一类是遍历一遍二叉树得出答案，第二类是通过分解问题计算出答案，
这两类思路分别对应着回溯算法核心框架 和 动态规划核心框架。

Resolve the sub problems
*/

func maxDepthSubProblemP1(root *TreeNode) int {

	var traverse func(*TreeNode) int

	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return common.MaxInt(traverse(root.Left)+1, traverse(root.Right)+1)
	}

	return traverse(root)
}
