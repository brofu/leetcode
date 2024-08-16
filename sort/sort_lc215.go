package sort

import "github.com/brofu/leetcode/common"

func findKthLargest(nums []int, k int) int {
	common.QuickSortInt(nums)
	return nums[len(nums)-k]
}

/**
Quick Select Version
KP.
	1.	Have the similar issue of `quick sort`
*/
func findKthLargestQuickSelect(nums []int, k int) int {

	kWant := len(nums) - k

	var quickSlectHelper func(int, int) int
	quickSlectHelper = func(low, high int) int {

		if low <= high {
			pivot := high
			cursor := low

			for i := low; i <= high; i++ {
				if nums[i] < nums[pivot] {
					nums[i], nums[cursor] = nums[cursor], nums[i]
					cursor++
				}
			}
			nums[cursor], nums[pivot] = nums[pivot], nums[cursor]

			if kWant == cursor {
				return nums[cursor]
			}

			if kWant < cursor {
				return quickSlectHelper(low, cursor-1)
			}

			return quickSlectHelper(cursor+1, high)
		}

		return -1
	}

	return quickSlectHelper(0, len(nums)-1)
}

func findKthLargestMergeSort(nums []int, k int) int {
	common.MergeSortPV2(nums)
	return nums[len(nums)-k]
}
