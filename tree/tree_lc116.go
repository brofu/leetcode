package tree

/*
DFS traverse the tree

Key Point:
传统的 traverse 函数是遍历二叉树的所有节点，但现在我们想遍历的其实是两个相邻节点之间的「空隙」。
一棵二叉树被抽象成了一棵三叉树，三叉树上的每个节点就是原先二叉树的两个相邻节点。
*/
func connectTraverseXX(root *Node) *Node {
	if root == nil {
		return root
	}
	var traverse func(*Node, *Node)
	traverse = func(left, right *Node) {
		if left == nil || right == nil {
			return
		}
		left.Next = right
		traverse(left.Left, left.Right)
		traverse(right.Left, right.Right)
		traverse(left.Right, right.Left)
	}
	traverse(root.Left, root.Right)
	return root
}

/*
BFS traverse the tree
*/
func connectxx(root *Node) *Node {
	if root == nil {
		return root
	}
	q := []*Node{root}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size-1; i++ {
			q[i].Next = q[i+1]
			if q[i].Left != nil {
				q = append(q, q[i].Left, q[i].Right)
			}
		}
		if q[size-1].Left != nil {
			q = append(q, q[size-1].Left, q[size-1].Right)
		}
		q = q[size:]
	}
	return root
}
