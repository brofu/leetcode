package stack

import (
	"github.com/brofu/leetcode/common"
)

/*
KP:
	1. Monotonic version.
	2. Use the `increasing stack` to get the leftmost index, of which the there is smaller on on it's right
	3. Use the `decreasing stack` to get the rightmost index, of which there is larger one on it's left

TC:
	O(N)
*/
func findUnsortedSubarray(nums []int) int {

	var (
		left, right = len(nums) - 1, 0
		inSt, deSt  []int
	)

	for idx := 0; idx < len(nums); idx++ {
		for len(inSt) > 0 && nums[inSt[len(inSt)-1]] > nums[idx] {
			left = common.MinInt(left, inSt[len(inSt)-1])
			inSt = inSt[:len(inSt)-1]
		}
		inSt = append(inSt, idx)
	}

	for idx := len(nums) - 1; idx >= 0; idx-- {
		for len(deSt) > 0 && nums[deSt[len(deSt)-1]] < nums[idx] {
			right = common.MaxInt(right, deSt[len(deSt)-1])
			deSt = deSt[:len(deSt)-1]
		}
		deSt = append(deSt, idx)
	}

	if left == len(nums)-1 && right == 0 {
		return 0
	}
	return right - left + 1
}

/*
KP:
	1. A version which is easier to understand
*/

func findUnsortedSubarrayV2(nums []int) int {

	left, right := -1, -1
	maxSofar, minSofar := nums[0], nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		if nums[i] >= maxSofar { // there is a smaller number on it's left
			maxSofar = nums[i]
		} else {
			right = i
		}
	}

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= minSofar {
			minSofar = nums[i]
		} else {
			left = i
		}
	}
	if left == -1 && right == -1 {
		return 0
	}
	return right - left + 1
}
