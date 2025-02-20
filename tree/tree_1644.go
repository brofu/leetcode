package tree

/*
1. Sub Task

Key Points
1.  p, or q NOT necessary in the tree, so, make 2 flags to mark if the nodes exists
2.  then, need to make sure, each node must be traversed, so,
3.	Put the checking logic `if root.value == p.value || root.value == q.value` to `post-order` location
*/

func westCommonAncestorIISubTaskPV1(root, p, q *TreeNode) *TreeNode {

	flagP, flagQ := false, false
	var traverse func(*TreeNode, *TreeNode, *TreeNode) *TreeNode
	traverse = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return root
		}

		left := traverse(root.Left, p, q)
		right := traverse(root.Right, p, q)

		if root.Val == p.Val {
			flagP = true
			return root
		}
		if root.Val == q.Val {
			flagQ = true
			return root
		}

		if left != nil && right != nil {
			return root
		}

		if left != nil {
			return left
		}
		return right
	}

	res := traverse(root, p, q)
	if flagP && flagQ {
		return res
	}
	return nil
}
