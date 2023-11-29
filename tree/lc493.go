package tree

import (
	"github.com/brofu/leetcode/common"
)

func reversePairs(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}

	// O(2N)
	sorted := make([]int, len(nums)*2)

	for index, val := range nums {
		sorted[index*2] = val
		sorted[index*2+1] = val * 2
	}

	common.MergeSort(sorted)

	// set up index map
	index := 0
	indexMap := make(map[int]int)
	for _, val := range sorted {
		if _, exists := indexMap[val]; !exists {
			indexMap[val] = index
			index += 1
		}
	}

	bit := NewBinaryIndexTreeReverse(len(indexMap))

	sum := 0
	for index := 0; index < len(nums); index++ {
		val := nums[index]
		sum += bit.Sum(indexMap[val*2] + 1)
		bit.Add(indexMap[val], 1)
	}

	return sum
}
