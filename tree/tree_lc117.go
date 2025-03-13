package tree

/*
BFS
*/
func connectPV1(root *Node) *Node {

	if root == nil {
		return root
	}

	q := []*Node{root}

	for len(q) > 0 {

		size := len(q)
		for idx := 0; idx < size-1; idx++ {
			node := q[idx]
			node.Next = q[idx+1]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		lastNode := q[size-1]
		if lastNode.Left != nil {
			q = append(q, lastNode.Left)
		}
		if lastNode.Right != nil {
			q = append(q, lastNode.Right)
		}
		q = q[size:]
	}

	return root
}
