package heap

import (
	ch "github.com/brofu/leetcode/common/heap"
)

func findKthLargest(nums []int, k int) int {
	h := ch.NewBinaryHeapInt()
	for _, num := range nums {
		h.Push(num)
		if h.Size() > k {
			h.Pop()
		}
	}
	return h.Pop()
}
