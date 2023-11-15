package tree

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	connectTraverse(root.Left, root.Right)

	return root
}

func connectTraverse(nodeLeft, nodeRight *Node) {
	if nodeLeft == nil || nodeRight == nil {
		return
	}

	nodeLeft.Next = nodeRight
	connectTraverse(nodeLeft.Left, nodeLeft.Right)
	connectTraverse(nodeLeft.Right, nodeRight.Left)
	connectTraverse(nodeRight.Left, nodeRight.Right)
}

//connectBFS utils the BFS
// The number of nodes in the tree is in the range [0, 212 - 1].
func connectBFS(root *Node) *Node {

	if root == nil {
		return root
	}

	queue := []*Node{}
	layerCount, elementCount := 1, 1

	queue = append(queue, root)

	for len(queue) > 0 {

		node := queue[0]

		if node.Left != nil { // push stack
			queue = append(queue, node.Left, node.Right)
		}

		if Pow(2, layerCount) == elementCount+1 { // the most right element of layer
			// if math.Pow(float64(2), float64(layerCount)) == float64(elementCount+1) { // the most right element a tree layer
			node.Next = nil
			layerCount += 1 // change to next layer
		} else {
			node.Next = queue[1]
		}

		elementCount += 1 // update element pointer
		queue = queue[1:]
	}

	return root
}
