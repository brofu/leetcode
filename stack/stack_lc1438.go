package stack

import (
	"github.com/brofu/leetcode/list"
)

/*
KP
	1. Utilize the monotonic queue
*/
func longestSubarray(nums []int, limit int) int {

	q := list.MonotonicQueue{}
	result, ws := 0, 0

	for fast := 0; fast < len(nums); fast++ {
		q.Push(nums[fast]) // extend the window
		ws++
		l, s := q.Max(), q.Min()
		if l-s > limit { // shrink the window
			q.Pop()
			ws--
		}
		if result < ws {
			result = ws
		}
	}
	return ws
}
