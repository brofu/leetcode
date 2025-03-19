package tree

func deepestLeavesSum(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var result int

	q := []*TreeNode{root}

	for len(q) > 0 {
		result = 0
		size := len(q)
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			result += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
	}
	return result
}
