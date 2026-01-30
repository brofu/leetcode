package array

import "github.com/brofu/leetcode/common"

/*
KP:
	1. Difference from problem 42:
		* When select 2 lines, then, other lines should consider as not exists

*/
func maxArea(height []int) int {

	result := 0

	left, right := 0, len(height)-1

	for left < right {
		r := common.MinInt(height[left], height[right]) * (right - left)
		if result < r {
			result = r
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}
