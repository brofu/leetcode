package backtrack

func subsets(nums []int) [][]int {

	result := [][]int{}
	track := []int{}

	var bt func(int, []int, []int, *[][]int)

	bt = func(start int, nums, track []int, result *[][]int) {
		temp := make([]int, len(track))
		copy(temp, track)
		*result = append(*result, temp)

		for i := start; i < len(nums); i++ {
			// choose
			track = append(track, nums[i])
			bt(i+1, nums, track, result)
			track = track[:len(track)-1]
		}
	}

	bt(0, nums, track, &result)
	return result
}
