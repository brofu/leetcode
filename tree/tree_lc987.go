package tree

import (
	"sort"
)

func verticalTraversal(root *TreeNode) [][]int {
	result := [][]int{}

	type node struct {
		row int
		col int
		val int
	}

	nodes := []node{}

	var traverse func(*TreeNode, int, int)

	traverse = func(root *TreeNode, row, col int) {
		if root == nil {
			return
		}

		nodes = append(nodes, node{row: row, col: col, val: root.Val})
		traverse(root.Left, row+1, col-1)
		traverse(root.Right, row+1, col+1)
	}

	traverse(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {

		if nodes[i].col == nodes[j].col && nodes[i].row == nodes[j].row {
			return nodes[i].val < nodes[j].val
		}
		if nodes[i].col == nodes[j].col {
			return nodes[i].row < nodes[j].row
		}
		return nodes[i].col < nodes[j].col
	})

	preCol := nodes[0].col
	temp := []int{}
	for _, node := range nodes {
		if node.col == preCol {
			temp = append(temp, node.val)
			continue
		}
		// a new col
		result = append(result, temp)
		temp = []int{node.val}
		preCol = node.col // update
	}
	result = append(result, temp)
	return result
}
