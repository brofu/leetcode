package tree

func constructMaximumBinaryTreeSubTaskP1(nums []int) *TreeNode {
	var traverse func([]int) *TreeNode

	traverse = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		maxValue, maxValueIndex := -1, -1
		for idx, v := range nums {
			if maxValue < v {
				maxValue = v
				maxValueIndex = idx
			}
		}
		node := &TreeNode{
			Val:   maxValue,
			Left:  traverse(nums[:maxValueIndex]),
			Right: traverse(nums[maxValueIndex+1:]),
		}
		return node
	}

	return traverse(nums)
}
