package sort

import "github.com/brofu/leetcode/common"

func findKthLargest(nums []int, k int) int {
	common.QuickSortInt(nums)
	return nums[len(nums)-k]
}
