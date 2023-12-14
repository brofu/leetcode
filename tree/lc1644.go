package tree

func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {

	var traverse func(root, p, q *TreeNode) *TreeNode
	var foundP, foundQ bool

	traverse = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return root
		}

		left := traverse(root.Left, p, q)
		right := traverse(root.Right, p, q)

		if root.Val == p.Val {
			foundP = true
			return root
		}
		if root.Val == q.Val {
			foundQ = true
			return root
		}

		if left == nil {
			return right
		}
		if right == nil {
			return left
		}

		return root
	}

	result := traverse(root, p, q)

	if foundP && foundQ {
		return result
	}

	return nil
}
