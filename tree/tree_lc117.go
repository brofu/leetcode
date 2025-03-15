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

/*
Follow-up: You may only use constant extra space.

Question:
How to resolve the problem with constant extra space?

Pointers

	Key Points
	1. 3 important pointers
		1. The `leftmost` node at each layer
		2. The `current` node within each layer
		3. The `pre` node of `current` node

	2. Connect the `children` at the `parent` layer and then, loop the `children` layer
*/
func connectWtihConstantSpace(root *Node) *Node {

	if root == nil {
		return root
	}

	leftmost := root

	for leftmost != nil {
		cur := leftmost
		leftmost = nil
		var pre *Node
		for cur != nil {
			pre, leftmost = processChildren(cur.Left, pre, leftmost)
			pre, leftmost = processChildren(cur.Right, pre, leftmost)
			cur = cur.Next
		}
	}

	return root
}

func processChildren(cur, pre, leftmost *Node) (*Node, *Node) {

	if cur == nil {
		return pre, leftmost
	}

	if pre != nil {
		pre.Next = cur
	} else {
		leftmost = cur
	}
	pre = cur
	return pre, leftmost
}
