package tree

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

			subTrees := []*TreeNode{}

			left := traverse(low, index-1)
			right := traverse(index+1, high)

			if len(left) == 0 {
				for _, rightTree := range right {
					node := &TreeNode{
						Val:   index,
						Right: rightTree,
					}
					subTrees = append(subTrees, node)
				}
			}

			if len(right) == 0 {
				for _, leftTree := range left {
					node := &TreeNode{
						Val:  index,
						Left: leftTree,
					}
					subTrees = append(subTrees, node)
				}
			}

			for _, leftTree := range left {
				for _, rightTree := range right {
					node := &TreeNode{
						Val:   index,
						Left:  leftTree,
						Right: rightTree,
					}
					subTrees = append(subTrees, node)
				}
			}
			trees = append(trees, subTrees...)
		}

		return trees
	}

	return traverse(1, n)
}
