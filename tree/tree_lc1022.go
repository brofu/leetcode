package tree

/*
Traverse
*/

func sumRootToLeaf(root *TreeNode) int {

	if root == nil {
		return 0
	}

	sum := 0
	var traverse func(*TreeNode, []int)

	traverse = func(root *TreeNode, path []int) {

		if root.Left == nil && root.Right == nil {
			sub := calculateSum(append(path, root.Val))
			sum += sub
			return
		}

		if root.Left != nil {
			traverse(root.Left, append(path, root.Val))
		}
		if root.Right != nil {
			traverse(root.Right, append(path, root.Val))
		}
	}

	traverse(root, []int{})

	return sum
}

func calculateSum(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum = sum*2 + n
	}

	return sum
}
