package dp

import (
	"container/heap"
	"math"

	myarray "github.com/brofu/leetcode/common/array"
	myheap "github.com/brofu/leetcode/common/heap"
)

func convertArray(nums []int) int {

	solution := func(nums []int) int {
		max := 5
		min := math.MaxInt

		dp := make([][]int, len(nums))
		for i := 0; i < len(dp); i++ {
			dp[i] = make([]int, max+1)
		}

		for j := 0; j < max+1; j++ {
			min = MinInt(min, int(math.Abs(float64(nums[0]-j))))
			dp[0][j] = min
		}

		for i := 1; i < len(dp); i++ {
			min := math.MaxInt
			for j := 0; j < max+1; j++ {
				min = MinInt(min, dp[i-1][j]+int(math.Abs(float64(nums[i]-j))))
				dp[i][j] = min
			}
		}
		for i := 0; i < len(dp); i++ {
		}
		return dp[len(nums)-1][max]
	}

	res := solution(nums)
	myarray.ReverseInt(nums)

	return MinInt(res, solution(nums))
}

/**
KP
	1. Really hard to understand
*/
func convertArrayHeap(nums []int) int {

	res1 := 0
	res2 := 0

	intHeap := &myheap.IntHeap{}
	for _, num := range nums {
		if intHeap.Len() > 0 && intHeap.Top() > num {
			res1 += heap.Pop(intHeap).(int) - num
			heap.Push(intHeap, num)
		}
		heap.Push(intHeap, num)
	}

	intHeap = &myheap.IntHeap{}
	for _, num := range nums {
		if intHeap.Len() > 0 && intHeap.Top() > -num {
			top := heap.Pop(intHeap).(int)
			res2 += top + num
			heap.Push(intHeap, -num)
		}
		heap.Push(intHeap, -num)
	}

	return MinInt(res1, res2)
}
