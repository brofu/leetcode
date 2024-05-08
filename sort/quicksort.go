package sort

import (
	"math/rand"
)

func QuickSort(nums []int, low, high int) {
	if low < high {
		pivot := quickSortPartition(nums, low, high)
		QuickSort(nums, low, pivot-1)
		QuickSort(nums, pivot+1, high)
	}
}

func quickSortPartition(nums []int, low, high int) int {

	pivot := low + rand.Intn(high-low+1)
	val := nums[pivot]
	cursor := low
	nums[pivot], nums[high] = nums[high], nums[pivot]

	for i := low; i < high; i++ {
		if nums[i] < val {
			nums[cursor], nums[i] = nums[i], nums[cursor]
			cursor++
		}
	}
	nums[high], nums[cursor] = nums[cursor], nums[high]
	return cursor
}
