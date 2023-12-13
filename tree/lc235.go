package tree

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {

	var traverse func(*TreeNode, *TreeNode, *TreeNode) *TreeNode

	traverse = func(root, p, q *TreeNode) *TreeNode {

		if root == nil || root.Val == p.Val || root.Val == q.Val {
			return root
		}

		if p.Val < root.Val && q.Val > root.Val || p.Val > root.Val && q.Val < root.Val {
			return root
		}

		if p.Val < root.Val && q.Val < root.Val {
			return traverse(root.Left, p, q)
		}

		if p.Val > root.Val && q.Val > root.Val {
			return traverse(root.Right, p, q)
		}

		return nil // This should NOT happen
	}

	return traverse(root, p, q)
}
