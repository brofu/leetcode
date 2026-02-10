package array

import "sort"

func minimumAbsDifference(arr []int) [][]int {

	sort.Ints(arr)

	m := 1000000 * 2
	result := make([][]int, 0)

	for idx := 1; idx < len(arr); idx++ {
		d := arr[idx] - arr[idx-1]
		if d == m {
			result = append(result, []int{arr[idx-1], arr[idx]})
		}
		if d < m {
			m = d
			result = [][]int{{arr[idx-1], arr[idx]}}
		}
	}

	return result
}
