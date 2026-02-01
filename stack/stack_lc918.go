package stack

import (
	"math"

	"github.com/brofu/leetcode/list"
)

/*
KP
	1. 2*n to handle the cycle array
	2. Sliding window (window size as n) to make sure the calculation range
	3. Pre-sum + Monotonic Queue
	4. If the array is NOT Cycle, then the problem same as 53, and approaches are:
		a. Presum
		b. DP

TC:
	1. O(N)

SC:
	1. O(N)
*/

func maxSubarraySumCircular(nums []int) int {

	n := len(nums)
	sum := nums[0]
	result := math.MinInt
	q := list.MonotonicQueue{}
	q.Push(nums[0])

	for slow, fast := 0, 1; fast < 2*n; fast++ {
		sum += nums[fast%n]
		if result < sum-q.Min() {
			result = sum - q.Min()
		}
		if fast-slow >= n { //shrink the window
			q.Pop()
			slow++
		}
		q.Push(sum)
	}
	return result
}
