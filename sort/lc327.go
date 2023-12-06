package sort

import (
	"github.com/brofu/leetcode/common"
)

func countRangeSum(nums []int, lower int, upper int) int {

	// pre sum
	preSum := make([]int, len(nums)+1)
	for index, val := range nums {
		preSum[index+1] = val + preSum[index]
	}

	return countRangeSumWrapper(preSum, 0, len(nums), lower, upper)
}

func countRangeSumWrapper(nums []int, left, right, lower, upper int) int {

	if left >= right {
		return 0
	}

	middle := (left + right) / 2

	leftCount := countRangeSumWrapper(nums, left, middle, lower, upper)
	rightCount := countRangeSumWrapper(nums, middle+1, right, lower, upper)

	count, m, n := 0, middle+1, middle+1
	for i := left; i <= middle; i++ {
		for m <= right && nums[m]-nums[i] < lower {
			m++
		}
		for n <= right && nums[n]-nums[i] <= upper {
			n++
		}
		count += n - m
	}

	common.MergeSort(nums[left : right+1])
	return count + leftCount + rightCount
}
