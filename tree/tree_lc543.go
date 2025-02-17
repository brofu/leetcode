package tree

import "github.com/brofu/leetcode/common"

/*
每一条二叉树的「直径」长度，就是某个节点的左右子树的最大深度之和。

问题转换为，求二叉树中每个节点的左右子树最大深度之和

后序逻辑。需要知道每个子树的左右子树最大深度
*/
func diameterOfBinaryTreeSubProblemP1(root *TreeNode) int {

	diameter := 0
	var traverse func(*TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftDepth := traverse(root.Left)
		rightDepth := traverse(root.Right)

		diameter = common.MaxInt(leftDepth+rightDepth, diameter)
		return common.MaxInt(leftDepth, rightDepth) + 1
	}

	traverse(root)
	return diameter
}
