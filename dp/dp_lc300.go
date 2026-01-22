package dp

import (
	"github.com/brofu/leetcode/common"
)

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

func lengthOfLISPV2DPF(nums []int) int {

	memo := make([]int, len(nums)+1)
	for idx := range memo {
		memo[idx] = -1
	}

	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 {
			return 1
		}
		if memo[n] != -1 {
			return memo[n]
		}
		result := 1
		for j := 0; j < n; j++ {
			if nums[j] < nums[n] {
				result = MaxInt(result, 1+dp(j))
			}
		}
		memo[n] = result
		return result
	}

	result := 1
	for n := len(nums) - 1; n >= 0; n-- {
		result = MaxInt(result, dp(n))
	}
	return result
}

/*
KP
1. From Down Top to get the 状态转移方程

TC:
1. O(N^2).

SC:
1. dp, O(N). Can optimized to O(2)
*/

func lengthOfLISPV2DPA(nums []int) int {

	dp := make([]int, len(nums))
	for idx := range dp {
		dp[idx] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = MaxInt(dp[i], dp[j]+1)
			}
		}
	}
	result := 0
	for _, res := range dp {
		result = MaxInt(result, res)
	}
	return result
}

/*
KP
Compare version of lengthOfLISWithBSPV2 and lengthOfLISWithBSPV3
1. V2 using [left, right), half close, half open, while V3 using [left, right]
2. V2 is better for this problem. Because, at the beginning, before there is actually a pile, there is NO data to check, that's [0, 0) can handle
*/

func lengthOfLISWithBSPV2(nums []int) int {

	piles := 0
	top := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		left, right := 0, piles
		current := nums[i]
		for left < right { // KP1. [left, right), since right initialized as piles
			// mid := (left + right) / 2 KP2. Risk of overflow
			mid := left + (right-left)/2
			if current > top[mid] {
				left = mid + 1 // KP3. Close at both side. So, [mid+1, right)
			} else if current < top[mid] {
				right = mid // KP4. Close at both side. So, [left, mid)
			} else {
				right = mid // KP5. We need to find the left-most pile. So, keep going on, [left, mid)
			}
		}
		if left == piles {
			piles++
		}
		top[left] = current
	}
	return piles
}

func lengthOfLISWithBSPV3(nums []int) int {

	piles := 0
	top := make([]int, len(nums))
	top[0] = -10001

	for i := 0; i < len(nums); i++ {
		left, right := 0, piles
		current := nums[i]
		for left <= right { // KP1. [left, right], since right initialized as piles
			// mid := (left + right) / 2 KP2. Risk of overflow
			mid := left + (right-left)/2
			if current > top[mid] {
				left = mid + 1 // KP3. Close at both side. So, [mid+1, right]
			} else if current < top[mid] {
				right = mid - 1 // KP4. Close at both side. So, [left, mid-1]
			} else {
				right = mid - 1 // KP5. We need to find the left-most pile. So, keep going on, [left, mid-1]
			}
		}
		if left == piles+1 {
			piles++
		}
		top[left] = current
	}
	return piles
}

/*
KP:
	1. The dp function version.
	2. dp(i), the LIS which ends with nums[i]

SC:
	1. With memo, the recursive depth should be O(N). N is the number of nums
*/
func lengthOfLISV1(nums []int) int {

	var (
		dp     func(int) int
		memo   = make([]int, len(nums))
		result = 1
	)

	for idx := range memo {
		memo[idx] = 0
	}

	dp = func(idx int) int {
		if idx == 0 {
			return 1
		}

		if memo[idx] != 0 {
			return memo[idx]
		}

		temp := 1
		for i := 0; i < idx; i++ {
			if nums[idx] > nums[i] {
				result := dp(i)
				if temp < result+1 {
					temp = result + 1
				}
			}
		}

		memo[idx] = temp
		return temp
	}

	for i := len(nums) - 1; i >= 0; i-- {
		temp := dp(i)
		if result < temp {
			result = temp
		}
	}

	return result
}

/*

KP:
	1. The dp table version
*/
func lengthOfLISV2(nums []int) int {

	var (
		result int
		dp     = make([]int, len(nums))
	)

	for idx := range dp {
		dp[idx] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	for _, n := range dp {
		if result < n {
			result = n
		}
	}

	return result
}
