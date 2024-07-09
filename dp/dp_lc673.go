package dp

import (
	"github.com/brofu/leetcode/common"
)

/**
KP.
	1. How to setup and update the dp array of `dpCount`?
		*	The number of LIS ended with nums[i]
*/
func findNumberOfLIS(nums []int) int {

	n, maxLen, result := len(nums), 0, 0
	dpLength, dpCount := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		dpLength[i] = 1
		dpCount[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {

				if dpLength[i] < dpLength[j]+1 {
					dpLength[i] = dpLength[j] + 1
					dpCount[i] = 0
				}
				if dpLength[i] == dpLength[j]+1 {
					dpCount[i] += dpCount[j]
				}

			}
		}
	}
	for i := 0; i < n; i++ {
		maxLen = common.MaxInt(maxLen, dpLength[i])
	}

	for i := 0; i < n; i++ {
		if dpLength[i] == maxLen {
			result += dpCount[i]
		}
	}

	return result
}
