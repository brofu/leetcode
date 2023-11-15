package tree

func connect117(root *Node) *Node {
	if root == nil {
		return root
	}

	connectTraverse117([]*Node{root.Left, root.Right})

	return root
}

func connectTraverse117(nodes []*Node) {

	nextLayerNodes := []*Node{}
	var firstNonNillNode *Node

	for _, node := range nodes {
		if node == nil {
			continue
		}

		nextLayerNodes = append(nextLayerNodes, node.Left, node.Right)
		if firstNonNillNode == nil {
			firstNonNillNode = node
		} else {
			firstNonNillNode.Next = node
			firstNonNillNode = node
		}
	}
	if len(nextLayerNodes) > 0 {
		connectTraverse117(nextLayerNodes)
	}
}

// connectMulPointer
func connectMulPointer(root *Node) *Node {

	sentinel := &Node{Val: -1, Next: root}

	for sentinel.Next != nil {
		cur := sentinel.Next
		layerPointer := sentinel

		for cur != nil {

			if cur.Left != nil {
				layerPointer.Next = cur.Left
				layerPointer = cur.Left
			}

			if cur.Right != nil {
				layerPointer.Next = cur.Right
				layerPointer = cur.Right
			}
			cur = cur.Next
		}

		layerPointer.Next = nil // Really need this.

	}
	return root
}
