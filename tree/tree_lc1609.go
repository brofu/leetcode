package tree

func isEvenOddTree(root *TreeNode) bool {

	q := []*TreeNode{root}

	depth := 0
	for len(q) > 0 {
		size := len(q)
		flag := 0
		if depth&1 == 1 { // odd layer
			flag = 100001
		}
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			if depth&1 == 0 { // even layer
				if node.Val&1 == 0 { // should be odd
					return false
				}
				if node.Val <= flag { // should be increasing
					return false
				}
			} else { // odd layer
				if node.Val&1 == 1 { // should be even
					return false
				}
				if node.Val >= flag { // should be decreasing
					return false
				}
			}
			flag = node.Val

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
		depth++
	}

	return true
}
