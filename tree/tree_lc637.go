package tree

func averageOfLevels(root *TreeNode) []float64 {

	result := []float64{}
	if root == nil {
		return result
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		sum := 0
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
		avarage := float64(sum) / float64(size)
		result = append(result, avarage)
	}
	return result
}
