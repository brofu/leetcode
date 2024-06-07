package backtrack

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	track := []int{}
	sum := 0

	var bt func([]int, []int, int, int, int, *[][]int)
	bt = func(candidates []int, track []int, start, sum, target int, result *[][]int) {

		// base case 1
		if sum == target {
			temp := make([]int, len(track))
			copy(temp, track)
			*result = append(*result, temp)
			return
		}

		// base case 2
		if sum > target {
			return
		}

		// track
		for i := start; i < len(candidates); i++ {

			// choose from list
			sum += candidates[i]
			track = append(track, candidates[i])

			// traverse next layer
			bt(candidates, track, i, sum, target, result)

			// cancel choose
			sum -= candidates[i]
			track = track[:len(track)-1]

		}
	}

	bt(candidates, track, 0, sum, target, &result)
	return result
}
