package tree

/*
BFS
*/
func levelOrder(root *TreeNode) [][]int {

	result := [][]int{}
	if root == nil {
		return result
	}

	q := []*TreeNode{root}

	for len(q) > 0 {

		size := len(q)
		temp := make([]int, size)
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			temp[idx] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, temp)
		q = q[size:]
	}

	return result
}
