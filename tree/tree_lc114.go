package tree

/*
Sub Task

Traverse function return the `most right leaf` node on the sub-tree
on the post-order position
1. link the right to the left node
2. link the right of return value to the right node
3. pay attention, when the right sub tree is nil. (return the left brother)
*/
func flattenSubTaskP1(root *TreeNode) {

	var traverse func(*TreeNode) *TreeNode

	// return the most right node on the sub tree
	traverse = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		if root.Left == nil && root.Right == nil {
			return root
		}

		rightMostOnRight := traverse(root.Right)
		rightMostOnLeft := traverse(root.Left)
		if rightMostOnLeft != nil {
			rightMostOnLeft.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}

		// important
		if rightMostOnRight == nil {
			rightMostOnRight = rightMostOnLeft
		}

		return rightMostOnRight
	}

	traverse(root)
}
