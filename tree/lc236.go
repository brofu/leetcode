package tree

// traverse the tree and get the ancestors
// compare the ancestors
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var traverse func(*TreeNode, int, *[]*TreeNode) bool

	traverse = func(root *TreeNode, key int, ancester *[]*TreeNode) bool {

		if root == nil {
			return false
		}

		if key == root.Val {
			*ancester = append(*ancester, root)
			return true
		}

		if traverse(root.Left, key, ancester) {
			*ancester = append(*ancester, root)
			return true
		}
		if traverse(root.Right, key, ancester) {
			*ancester = append(*ancester, root)
			return true
		}

		return false // this should NOT happen
	}

	pAncestors, qAncestors := []*TreeNode{}, []*TreeNode{}

	traverse(root, p.Val, &pAncestors)
	traverse(root, q.Val, &qAncestors)

	i, j := len(pAncestors)-1, len(qAncestors)-1

	for ; i >= 0 && j >= 0 && pAncestors[i].Val == qAncestors[j].Val; i, j = i-1, j-1 {
	}
	return pAncestors[i+1]
}

func lowestCommonAncestorV2(root, p, q *TreeNode) *TreeNode {

	var traverse func(*TreeNode, *TreeNode, *TreeNode) *TreeNode

	traverse = func(root, p, q *TreeNode) *TreeNode {
		if root == nil || root.Val == p.Val || root.Val == q.Val {
			return root
		}

		left := traverse(root.Left, p, q)
		right := traverse(root.Right, p, q)

		if left != nil && right != nil {
			return root
		}

		if left != nil {
			return left
		}

		return right
	}

	return traverse(root, p, q)
}
