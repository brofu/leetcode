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

/*
Sub Task

Utilize the property of BST
*/

func lowestCommonAncestorBSTSubTaskPV1(root, p, q *TreeNode) *TreeNode {

	var traverse func(*TreeNode) *TreeNode

	traverse = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		if root.Val == p.Val || root.Val == q.Val {
			return root
		}

		if p.Val < root.Val && q.Val < root.Val {
			return traverse(root.Left)
		}
		if p.Val > root.Val && q.Val > root.Val {
			return traverse(root.Right)
		}
		return root
	}

	return traverse(root)
}

/*
The steps taken are also similar to approach DFS.
The only difference is instead of recursively calling the function, we traverse down the tree iteratively.
This is possible without using a stack or recursion since we don't need to backtrace to find the LCA node.
In essence of it the problem is iterative, it just wants us to find the split point.
The point from where p and q won't be part of the same subtree or when one is the parent of the other.
*/
func lowestCommonAncestorBSTIteratePV1(root, p, q *TreeNode) *TreeNode {

	node := root
	for node != nil {
		if node.Val > p.Val && node.Val > q.Val {
			node = node.Left
		} else if node.Val < p.Val && node.Val < q.Val {
			node = node.Right
		} else {
			return node
		}
	}
	return node
}
