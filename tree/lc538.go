package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	convertBSTWrapper(root, 0)
	return root
}

func convertBSTWrapper(root *TreeNode, count int) (res int) {

	if root == nil {
		return 0
	}

	res += count

	if root.Right != nil {
		res = convertBSTWrapper(root.Right, count)
	}

	res += root.Val

	root.Val = res

	if root.Left == nil {
		return res
	}

	return convertBSTWrapper(root.Left, res)
}

func convertBSTV2(root *TreeNode) *TreeNode {
	sum := 0
	traverse := func(*TreeNode) {}
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Right)

		sum += root.Val
		root.Val = sum

		traverse(root.Left)
	}
	traverse(root)
	return root
}
