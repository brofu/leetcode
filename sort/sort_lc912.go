package sort

import (
	"github.com/brofu/leetcode/common"
)

func sortArray(nums []int) []int {
	//common.MergeSort(nums)
	HeapSortIterativeV3(nums)
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

func sortArrayPV1(nums []int) []int {
	common.MergeSortV2(nums)
	return nums
}
