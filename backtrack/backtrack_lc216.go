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

/*

TC:
	1. Search node number is C(9,1) + C(9,2) + ... + C(9, k),
	2. Upper bound of possible results is C(9, k)
	3. With copy cost, we have C(9, k)*k

SC:
	1. recursive depth: O(k)
*/
func combinationSum3V1(k int, n int) [][]int {

	var (
		result [][]int
		bt     func(int, int, []int)
	)

	bt = func(idx, sum int, track []int) {

		// base case
		if sum == n && len(track) == k { // get 1 result
			temp := make([]int, k)
			copy(temp, track)
			result = append(result, temp)
			return
		}

		if idx == 10 || sum > n || len(track) > k {
			return
		}

		for i := idx; i < 10; i++ {
			// choose, explore, cancel
			bt(i+1, sum+i, append(track, i))
		}
	}

	bt(1, 0, []int{})
	return result
}
