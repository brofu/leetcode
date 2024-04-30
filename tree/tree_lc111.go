package tree

/*
BFS version
T: O(n)
M: O(n)
*/
func minDepth(root *TreeNode) int {

	depth := 0

	if root == nil {
		return depth
	}

	queue := []*TreeNode{root}
	depth = 1

	for len(queue) > 0 {

		sz := len(queue)
		numberOfLayer := 0
		for ; numberOfLayer < sz; numberOfLayer += 1 {
			node := queue[numberOfLayer]
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth += 1
		queue = queue[numberOfLayer:]
	}

	return depth
}

/*
DFS: in-order
*/

func minDepthV2(root *TreeNode) int {

	depth := 0

	var helper func(*TreeNode, int)

	helper = func(root *TreeNode, d int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			if depth > d || depth == 0 {
				depth = d
			}
		}

		if root.Left != nil {
			helper(root.Left, d+1)
		}

		if root.Right != nil {
			helper(root.Right, d+1)
		}
	}

	helper(root, 1)

	return depth
}
