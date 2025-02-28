package tree

/*
Traverse + BFS
*/

func rightSideView(root *TreeNode) []int {

	res := []int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, q[size-1].Val)
		q = q[size:]
	}

	return res
}
