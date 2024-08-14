package tree

/**
KP.
	traverse the tree and get the ancestors
	compare the ancestors
	Why need to loop with reversed order?
*/
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

/**
KP
Search the nodes in sub-tree. 2 situations:
1. One on the left tree and one on the right tree
2. One is the parent of the other
*/

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

func lowestCommonAncestorV2PV1(root, p, q *TreeNode) *TreeNode {

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

func lowestCommonAncestorPV1(root, p, q *TreeNode) *TreeNode {

	var traverse func(*TreeNode, *TreeNode, *[]*TreeNode) bool

	traverse = func(root, target *TreeNode, track *[]*TreeNode) bool {

		if root == nil {
			return false
		}

		if root.Val == target.Val {
			*track = append(*track, root)
			return true
		}

		if traverse(root.Left, target, track) {
			*track = append(*track, root)
			return true
		}

		if traverse(root.Right, target, track) {
			*track = append(*track, root)
			return true
		}

		return false
	}

	var parentsP, parentsQ []*TreeNode
	traverse(root, p, &parentsP)
	traverse(root, q, &parentsQ)
	i, j := len(parentsP)-1, len(parentsQ)-1
	for ; i >= 0 && j >= 0 && parentsP[i].Val == parentsQ[j].Val; i, j = i-1, j-1 {
	}

	return parentsP[i+1]
}

func lowestCommonAncestorPV2(root, p, q *TreeNode) *TreeNode {

	var traverse func(root, p, q *TreeNode) *TreeNode

	traverse = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		if root != nil && (root.Val == p.Val || root.Val == q.Val) {
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
		if right != nil {
			return right
		}

		return nil
	}

	return traverse(root, p, q)
}

func lowestCommonAncestorPV3(root, p, q *TreeNode) *TreeNode {

	var dfs func(*TreeNode, *TreeNode, *TreeNode) *TreeNode

	dfs = func(root, p, q *TreeNode) *TreeNode {

		// base case
		if root == nil {
			return nil
		}
		if root.Val == p.Val || root.Val == q.Val {
			return root
		}

		left := dfs(root.Left, p, q)   // search left
		right := dfs(root.Right, p, q) // search right

		if left != nil && right != nil { // p, q in different child tree
			return root
		}

		if left != nil {
			return left
		}

		return right
	}

	result := dfs(root, p, q)
	return result
}
