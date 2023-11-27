package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	index, val := FindMax(nums)

	root := &TreeNode{
		Val: val,
	}

	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])

	return root
}
