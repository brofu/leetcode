package tree

/*
wrong version
*/

type TreeNodeWithIndex struct {
	*TreeNode
	Index int
}

/*
BFS

Key Points
There scenarios:
1. [1, null, 3]
2. [1, 2, 3, 4, null, 6]
3. [1, 2, null, 4]
*/
func isCompleteTree(root *TreeNode) bool {

	if root == nil {
		return true
	}

	q := []TreeNodeWithIndex{{root, 1}}
	sum := 0
	maxIndex := 0

	for len(q) > 0 {
		size := len(q)
		sum += size
		hasNull := false
		for idx := 0; idx < size; idx++ {
			node := q[idx]

			if node.Left != nil {
				q = append(q, TreeNodeWithIndex{node.Left, 2 * node.Index})
			}
			if node.Right != nil {
				q = append(q, TreeNodeWithIndex{node.Right, 2*node.Index + 1})
			}

			if node.Left == nil && node.Right != nil {
				return false
			}
			if node.Left == nil || node.Right == nil {
				hasNull = true
				continue
			}
			if hasNull && (node.Left != nil || node.Right != nil) {
				return false
			}
		}
		maxIndex = q[size-1].Index
		q = q[size:]
	}

	return sum == maxIndex
}
