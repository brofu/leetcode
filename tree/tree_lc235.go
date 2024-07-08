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

func lowestCommonAncestorBSTPV1(root, p, q *TreeNode) *TreeNode {

	var traverse func(*TreeNode, *TreeNode, *TreeNode) *TreeNode

	traverse = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return root
		}

		small, big := p, q
		if p.Val > q.Val {
			big = p
			small = q
		}

		if root.Val >= small.Val && root.Val <= big.Val {
			return root
		}

		if root.Val < small.Val {
			return traverse(root.Right, p, q)
		} else {
			return traverse(root.Left, p, q)
		}
	}

	return traverse(root, p, q)
}
