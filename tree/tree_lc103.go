package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
KP
	1. bfs
*/
func zigzagLevelOrder(root *TreeNode) [][]int {

	result := [][]int{}

	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	direction := 1

	for len(queue) > 0 {
		sz := len(queue)
		index := 0
		numbers := []int{}

		for ; index < sz; index += 1 {
			if direction == 1 {
				numbers = append(numbers, queue[index].Val)
			} else {
				numbers = append(numbers, queue[sz-index-1].Val)
			}
			node := queue[index]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, numbers)
		queue = queue[index:]
		direction = -direction
	}

	return result
}

/**
KP
	1. DFS
*/
func zigzagLevelOrderDFS(root *TreeNode) [][]int {

	result := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(result) == level {
			result = append(result, []int{root.Val})
		} else {
			result[level] = append(result[level], root.Val)
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	revert := func(result [][]int) {
		for index, list := range result {
			if index%2 == 1 {
				for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
					list[i], list[j] = list[j], list[i]
				}
			}
		}
	}

	revert(result)
	return result
}
