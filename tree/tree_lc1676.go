package tree

/*
1. Sub Task
2. Optimization at in-order location

Key Points
1. Same principle as find LCA of 2 nodes
*/
func lowestCommonAncestorIVSubTaskPV1(root *TreeNode, nodes []*TreeNode) *TreeNode {

	var lca *TreeNode
	nodesMap := make(map[int]struct{})
	for _, node := range nodes {
		nodesMap[node.Val] = struct{}{}
	}

	var traverse func(*TreeNode) *TreeNode
	traverse = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		if _, exits := nodesMap[root.Val]; exits {
			return root
		}
		left := traverse(root.Left)

		if lca != nil {
			return lca
		}

		right := traverse(root.Right)

		if left != nil && right != nil {
			return root
		}
		if left != nil {
			return left
		}
		return right
	}

	return traverse(root)
}
