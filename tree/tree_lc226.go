package tree

/*
traverse approach
*/
func invertTreeTraverserP1(root *TreeNode) *TreeNode {

	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)

	return root
}
