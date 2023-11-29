package tree

import (
	"github.com/brofu/leetcode/common"
)

func countSmaller(nums []int) []int {

	res := make([]int, len(nums))

	if len(nums) <= 1 {
		return res
	}

	// copy S: O(N)
	sorted := make([]int, len(nums))

	// copy T: O(N)
	_ = copy(sorted, nums)

	// sort T: O(N*LogN)
	common.MergeSort(sorted)
	// setup index map. T: O(N) S: O(N)
	indexMap := make(map[int]int)

	index := 0
	for _, val := range sorted {
		if _, exist := indexMap[val]; !exist {
			indexMap[val] = index
			index += 1
		}
	}

	bit := NewBinaryIndexTree(len(indexMap))
	for i := len(nums) - 1; i >= 0; i-- {
		val := nums[i]
		index := indexMap[val]
		count := bit.Sum(index - 1) // query occur times
		res[i] = count
		bit.Add(index, 1)
	}

	return res
}
