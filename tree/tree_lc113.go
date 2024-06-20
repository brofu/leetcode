package tree

/**
KP
	1.	Similar to backtrack.
	2.	Backtrack is actually DFS.
*/
func pathSum(root *TreeNode, targetSum int) [][]int {

	result := [][]int{}
	var bt func(*TreeNode, []int, int, *[][]int)

	bt = func(root *TreeNode, track []int, remain int, result *[][]int) {

		// base case
		if root == nil {
			return
		}

		// base case
		if root.Left == nil && root.Right == nil && remain == root.Val {
			track = append(track, root.Val)
			*result = append(*result, append([]int(nil), track...))
			return
		}

		// choose, next layer and cancel choose
		bt(root.Left, append(track, root.Val), remain-root.Val, result)
		bt(root.Right, append(track, root.Val), remain-root.Val, result)
	}

	bt(root, []int{}, targetSum, &result)
	return result
}
