package array

import (
	"sort"
)

func arrayRankTransform(arr []int) []int {

	if len(arr) == 0 {
		return arr
	}

	sour := make([]int, len(arr))
	copy(sour, arr)

	sort.Ints(arr)

	rankMap := make(map[int]int)
	rank := 1
	prev := arr[0]

	for _, x := range arr {
		if x != prev {
			prev = x
			rank++
		}
		rankMap[x] = rank
	}

	result := make([]int, len(arr))

	for idx, v := range sour {
		result[idx] = rankMap[v]
	}

	return result
}
