package tree

import (
	"github.com/brofu/leetcode/common"
)

func widthOfBinaryTreeProblematicVersion(root *TreeNode) int {

	if root == nil {
		return 0
	}

	result := 1
	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) > 0 {
		size := len(queue)
		nextLevelCount := 0

		var maxIndex int
		for maxIndex = size; maxIndex > 0; maxIndex-- {
			node := queue[maxIndex-1]
			if node == nil {
				continue
			}
			if node.Right != nil {
				nextLevelCount = 2 * maxIndex
				break
			}
			if node.Left != nil {
				nextLevelCount = 2*maxIndex - 1
				break
			}
		}

		for i := 0; i < maxIndex; i++ {
			var left, right *TreeNode
			node := queue[i]
			if node == nil {
				queue = append(queue, []*TreeNode{left, right}...)
				continue
			}
			if node.Left != nil {
				left = node.Left
			}
			queue = append(queue, left)
			if node.Right != nil {
				right = node.Right
			}
			queue = append(queue, right)
		}
		result = common.MaxInt(result, nextLevelCount)
		queue = queue[size:]
	}

	return result
}

func widthOfBinaryTree(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}

	type queueNode struct {
		node *TreeNode
		col  int
	}

	queue := make([]queueNode, 1)
	queue[0] = queueNode{root, 1}

	for len(queue) > 0 {
		size := len(queue)

		var i int
		var qn queueNode
		for i = 0; i < size; i++ {
			qn = queue[i]
			col := qn.col
			node := qn.node
			if node.Left != nil {
				queue = append(queue, queueNode{node.Left, 2*col - 1})
			}
			if node.Right != nil {
				queue = append(queue, queueNode{node.Right, 2 * col})
			}
		}
		// KP. pay attention to the defination of `width`. Special case: case 4
		result = MaxInt(result, qn.col-queue[0].col+1)
		queue = queue[size:]
	}
	return result
}

func widthOfBinaryTreePV1(root *TreeNode) int {

	if root == nil {
		return 0
	}

	maxWidth := 0

	type QueueNode struct {
		node *TreeNode
		col  int
	}
	q := []QueueNode{QueueNode{root, 1}}

	for len(q) > 0 {
		size := len(q)
		for idx := 0; idx < size; idx++ {
			qn := q[idx]
			if qn.node.Left != nil {
				q = append(q, QueueNode{qn.node.Left, 2 * qn.col})
			}
			if qn.node.Right != nil {
				q = append(q, QueueNode{qn.node.Right, 2*qn.col + 1})
			}
		}
		maxWidth = common.MaxInt(maxWidth, q[size-1].col-q[0].col+1)
		q = q[size:]
	}

	return maxWidth
}
