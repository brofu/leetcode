package tree

import (
	"math"

	"github.com/brofu/leetcode/common"
)

/*
BFS
*/
func largestValues(root *TreeNode) []int {

	result := []int{}
	if root == nil {
		return result
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		maxNum := math.MinInt
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			maxNum = common.MaxInt(maxNum, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, maxNum)

		q = q[size:]
	}
	return result
}
