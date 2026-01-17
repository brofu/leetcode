package backtrack

import (
	"github.com/brofu/leetcode/tree"
)

func pathOfDeepestOfBinaryTree(root *tree.TreeNode) []int {

	if root == nil {
		return []int{}
	}

	var (
		maxDept int
		result  []int
		bt      func(*tree.TreeNode, int, []int)
	)

	bt = func(root *tree.TreeNode, depth int, track []int) {

		if root.Left == nil && root.Right == nil {
			if depth > maxDept { // a new deeper depth
				maxDept = depth
				temp := make([]int, len(track))
				copy(temp, track)
				result = temp
			}
			return
		}

		// An updated version of back track. No need cancel choose
		for _, node := range []*tree.TreeNode{root.Left, root.Right} {
			if node == nil {
				continue
			}
			bt(node, depth+1, append(track, node.Val))
		}
	}

	bt(root, 1, []int{root.Val})

	return result
}

func pathOfDeepestOfBinaryTreeClassicVersion(root *tree.TreeNode) []int {

	var bt func(*tree.TreeNode)

	var (
		depth, maxDepth int
		track, result   []int
	)

	bt = func(root *tree.TreeNode) {
		if root == nil {
			return
		}

		// choose
		depth++
		track = append(track, root.Val)

		if root.Left == nil && root.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				result = append([]int(nil), track...) // copy
			}
		}

		// next layer
		bt(root.Left)
		bt(root.Right)

		// cancel choose
		depth--
		track = track[0 : len(track)-1]
	}

	bt(root)

	return result
}
