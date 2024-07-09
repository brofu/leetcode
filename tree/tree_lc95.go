package tree

/**
KP.
	How to process the edged case of `low > high` and `low == high`?
	2 different versions.
*/

func generateTrees(n int) []*TreeNode {

	var traverse func(int, int) []*TreeNode

	traverse = func(low, high int) []*TreeNode {

		if low > high {
			return nil
		}

		var trees []*TreeNode

		if low == high {
			trees = append(trees, &TreeNode{
				Val: low,
			})
			return trees
		}

		for index := low; index <= high; index++ {

			left := traverse(low, index-1)
			right := traverse(index+1, high)

			if len(left) == 0 {
				for _, rightTree := range right {
					node := &TreeNode{
						Val:   index,
						Right: rightTree,
					}
					trees = append(trees, node)
				}
			}

			if len(right) == 0 {
				for _, leftTree := range left {
					node := &TreeNode{
						Val:  index,
						Left: leftTree,
					}
					trees = append(trees, node)
				}
			}

			for _, leftTree := range left {
				for _, rightTree := range right {
					node := &TreeNode{
						Val:   index,
						Left:  leftTree,
						Right: rightTree,
					}
					trees = append(trees, node)
				}
			}
		}

		return trees
	}

	return traverse(1, n)
}

func generateTreesPV1(n int) []*TreeNode {

	var helper func(int, int) []*TreeNode

	helper = func(low, high int) []*TreeNode {

		result := []*TreeNode{}

		if low > high {
			result = append(result, nil)
			return result
		}

		for i := low; i <= high; i++ {
			left := helper(low, i-1)
			right := helper(i+1, high)

			for _, treeLeft := range left {
				for _, treeRight := range right {
					node := &TreeNode{
						Val:   i,
						Left:  treeLeft,
						Right: treeRight,
					}
					result = append(result, node)
				}
			}
		}
		return result
	}

	return helper(1, n)
}
