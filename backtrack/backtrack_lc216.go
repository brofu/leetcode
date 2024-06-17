package backtrack

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}

	var bt func([]int, int, int)

	bt = func(track []int, start, sum int) {

		// base case
		if sum == n && len(track) == k {
			result = append(result, append([]int(nil), track...))
			return
		}

		if sum > n || len(track) > k {
			return
		}

		// choose, next layer travers and cancel choose
		// may add more pruning logics here
		for i := start; i <= 9; i++ {
			bt(append(track, i), i+1, sum+i)
		}
	}

	bt([]int{}, 1, 0)
	return result
}
