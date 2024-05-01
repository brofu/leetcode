package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {

	result := [][]int{}

	if root == nil {
		return nil
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
