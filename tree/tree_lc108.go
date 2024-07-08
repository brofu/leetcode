package tree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var helper func([]int, int, int) *TreeNode

	helper = func(nums []int, start, end int) *TreeNode {

		if start > end {
			return nil
		}
		mid := (start + end) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = helper(nums, start, mid-1)
		root.Right = helper(nums, mid+1, end)
		return root
	}

	return helper(nums, 0, len(nums)-1)
}
