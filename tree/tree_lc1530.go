package tree

func countPairs(root *TreeNode, distance int) int {

	var dfs func(*TreeNode, int, int) (int, []int)

	calculate := func(leftLeavies, rightLeavies []int, parentHeight, distance int) int {
		count := 0
		for _, ll := range leftLeavies {
			for _, rl := range rightLeavies {
				if ll-parentHeight+rl-parentHeight <= distance {
					count++
				}
			}
		}
		return count
	}

	dfs = func(root *TreeNode, distance int, height int) (int, []int) {

		// base case
		if root == nil {
			return 0, nil
		}

		if root.Left == nil && root.Right == nil { // return the height of current node
			return 0, []int{height}
		}

		leftCount, leftLeavies := dfs(root.Left, distance, height+1)
		rightCount, rightLeavies := dfs(root.Right, distance, height+1)

		count := calculate(leftLeavies, rightLeavies, height, distance)
		return leftCount + rightCount + count, append(leftLeavies, rightLeavies...)
	}

	res, _ := dfs(root, distance, 0)

	return res
}
