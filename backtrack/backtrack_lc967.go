package backtrack

import "github.com/brofu/leetcode/common"

func numsSameConsecDiff(n int, k int) []int {

	var result []int
	temp := [][]int{}

	var bt func([]int)
	bt = func(track []int) {

		//base case
		if len(track) == n {
			m := make([]int, len(track))
			copy(m, track)
			temp = append(temp, m)
			return
		}

		for i := 0; i <= 9; i++ {

			// pruning 1
			if len(track) == 0 && i == 0 {
				continue
			}
			// pruning 1
			if len(track) > 0 && common.AbsIntSub(i, track[len(track)-1]) != k {
				continue
			}

			// choose
			bt(append(track, i))
			// cancel
		}
	}
	bt([]int{})
	result = make([]int, len(temp))
	for idx, t := range temp {
		for _, n := range t {
			result[idx] = result[idx]*10 + n
		}
	}

	return result
}

func numsSameConsecDiffV2(n int, k int) []int {

	var results []int
	var bt func(int, int, int)

	bt = func(digits, pre, result int) {
		// base case
		if digits == n {
			results = append(results, result)
			return
		}

		for i := 0; i <= 9; i++ {
			// pruning 1
			if digits == 0 && i == 0 {
				continue
			}
			if digits > 0 && common.AbsIntSub(i, pre) != k {
				continue
			}

			// choose
			bt(digits+1, i, result*10+i)
			// cancel
		}
	}

	bt(0, 0, 0)

	return results
}
