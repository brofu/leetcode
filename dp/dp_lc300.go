package dp

import "github.com/brofu/leetcode/common"

// nums has at least 1 element
func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = MaxInt(dp[i], dp[j]+1)
			}
		}
	}

	res := 1
	for _, re := range dp {
		if res < re {
			res = re
		}
	}

	return res
}

func lengthOfLISWithBS(nums []int) int {

	top := make([]int, len(nums))

	piles := 0

	for _, number := range nums {

		left, right := 0, piles

		for left < right {
			mid := (left + right) / 2
			if top[mid] < number {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == piles {
			piles += 1
		}

		top[left] = number
	}

	return piles
}

func lengthOfLISWithBSPV1(nums []int) int {

	top := make([]int, len(nums)+1)
	top[0] = -10001
	piles := 0

	for _, num := range nums {

		left, right := 0, piles
		for left <= right {
			mid := left + (right-left)/2
			if top[mid] < num {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left == piles+1 {
			piles++
		}

		top[left] = num
	}
	return piles
}

func lengthOfLISPV1(nums []int) int {

	result := 0
	dp := make([]int, len(nums)) // the LIS ends with nums[i]
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = common.MaxInt(dp[i], dp[j]+1)
			}
		}
	}

	for _, dis := range dp {
		if result < dis {
			result = dis
		}
	}

	return result
}
