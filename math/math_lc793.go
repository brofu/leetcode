package math

import (
	"math"
)

/*
KP:
	1. Binary Search with function from problem 172.
*/

func preimageSizeFZF(k int) int {

	leftMost, rightMost := 0, 0

	left, right := 0, math.MaxInt64>>31

	for left <= right {
		mid := left + (right-left)/2
		res := trailingZeroes(mid)
		if res >= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	leftMost = left

	left, right = 0, math.MaxInt64>>31
	for left <= right {
		mid := left + (right-left)/2
		res := trailingZeroes(mid)
		if res > k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	rightMost = right

	return rightMost - leftMost + 1
}
