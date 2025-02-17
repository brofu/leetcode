package tree

/*
二叉树题目的递归解法可以分两类思路，第一类是遍历一遍二叉树得出答案，第二类是通过分解问题计算出答案，
这两类思路分别对应着回溯算法核心框架 和 动态规划核心框架。

Resolve the sub problems
*/

func preorderTraversalSubProblemP1(root *TreeNode) []int {

	var traverse func(*TreeNode) []int

	traverse = func(root *TreeNode) []int {
		if root == nil {
			return []int{}
		}
		result := []int{root.Val}

		return append(append(result, traverse(root.Left)...), traverse(root.Right)...)
	}

	return traverse(root)
}
