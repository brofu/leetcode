package tree

func lowestCommonAncestorIV(root *TreeNode, nodes []*TreeNode) *TreeNode {

	var foundNode *TreeNode

	var traverse func(*TreeNode) int

	traverse = func(root *TreeNode) int {

		foundCount := 0
		if root == nil {
			return 0
		}

		left := traverse(root.Left)
		right := traverse(root.Right)

		foundCount += left + right
		// check current node
		for _, node := range nodes {
			if node == root {
				foundCount += 1
				break
			}
		}

		if len(nodes) == foundCount && foundNode == nil {
			foundNode = root
		}

		return foundCount
	}

	traverse(root)

	return foundNode
}
