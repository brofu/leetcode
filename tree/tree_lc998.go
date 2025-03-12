package tree

/*
Sub Tasks
*/
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {

	var traverse func(*TreeNode, int) *TreeNode

	traverse = func(root *TreeNode, val int) *TreeNode {

		if root == nil {
			return &TreeNode{Val: val}
		}

		if root.Val < val {
			node := &TreeNode{
				Val:  val,
				Left: root,
			}
			return node
		}

		root.Right = traverse(root.Right, val)
		return root
	}

	return traverse(root, val)
}
