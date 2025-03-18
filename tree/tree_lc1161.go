package tree

import "math"

/*
BFS
*/
func maxLevelSum(root *TreeNode) int {

	result := 0
	if root == nil {
		return 0
	}

	sum := math.MinInt
	q := []*TreeNode{root}
	depth := 0
	for len(q) > 0 {
		depth++
		size := len(q)
		temp := 0
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			temp += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum < temp {
			sum = temp
			result = depth
		}
		q = q[size:]
	}
	return result
}

/*
DFS

Key Points

List or Map?
1. List, don't know the length of the numbers
2. Map, the order of loop is not fixed.
*/

func maxLevelSumDFS(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var traverse func(*TreeNode, int)

	sumMap := make(map[int]int)

	traverse = func(root *TreeNode, depth int) {

		val, exists := sumMap[depth]
		if !exists {
			val = root.Val
		} else {
			val += root.Val
		}
		sumMap[depth] = val
		if root.Left != nil {
			traverse(root.Left, depth+1)
		}
		if root.Right != nil {
			traverse(root.Right, depth+1)
		}
	}

	traverse(root, 1)

	maxSum := math.MinInt
	result := 0
	for k, v := range sumMap {
		if maxSum < v {
			maxSum = v
			result = k
		}
	}

	return result
}
