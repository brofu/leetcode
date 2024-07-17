package tree

import (
	"github.com/brofu/leetcode/common"
)

// wrong version. Need BFS
func verticalOrderWrong(root *TreeNode) [][]int {

	data := make(map[int][]int)
	minCol, maxCol := 0, 0

	var traverse func(*TreeNode, int)

	traverse = func(root *TreeNode, col int) {
		if root == nil {
			return
		}

		if d, ok := data[col]; ok {
			d = append(d, root.Val)
			data[col] = d
		} else {
			data[col] = []int{root.Val}
		}
		minCol = common.MinInt(minCol, col)
		maxCol = common.MaxInt(maxCol, col)

		traverse(root.Left, col-1)
		traverse(root.Right, col+1)
	}

	traverse(root, 0)

	result := make([][]int, maxCol-minCol+1)
	for i := minCol; i <= maxCol; i++ {
		result[i-minCol] = data[i]
	}

	return result
}

/**
KP.
	1. Must BFS
	2. Edged case root == nil
*/
func verticalOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	data := make(map[int][]int)
	minCol, maxCol := 0, 0

	type reco struct {
		col  int
		node *TreeNode
	}
	queue := make([]reco, 1)
	queue[0] = reco{0, root}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {

			record := queue[i]
			col := record.col
			node := record.node

			minCol = common.MinInt(minCol, col)
			maxCol = common.MaxInt(maxCol, col)

			if d, ok := data[col]; ok {
				d = append(d, node.Val)
				data[col] = d
			} else {
				data[col] = []int{node.Val}
			}

			if node.Left != nil {
				queue = append(queue, reco{col - 1, node.Left})
			}
			if node.Right != nil {
				queue = append(queue, reco{col + 1, node.Right})
			}
		}
		queue = queue[size:]
	}

	result := make([][]int, maxCol-minCol+1)
	for i := minCol; i <= maxCol; i++ {
		result[i-minCol] = data[i]
	}

	return result
}
