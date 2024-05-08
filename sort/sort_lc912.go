package sort

import (
	"github.com/brofu/leetcode/common"
)

func sortArray(nums []int) []int {
	common.MergeSort(nums)
	return nums
}

func sortArrayHeap(nums []int) []int {

	HeapSort(nums)
	return nums
}

func sortQuickSort(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}
