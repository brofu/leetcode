package tree

func balanceBST(root *TreeNode) *TreeNode {

	nums := []int{}

	// in-order traverse to get an ordered slice
	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		nums = append(nums, root.Val)
		traverse(root.Right)
	}
	traverse(root)

	// generate a balanced BST from ordered slice
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
