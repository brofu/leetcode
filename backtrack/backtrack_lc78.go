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

/**
KP
	1.	Thinking from the perspective of balls
		pay attention to the foor loop. For 1 ball, there is 2 possible situation, selected or NOT selected. So, no need for loop anymore.
*/

func subsetsV2(nums []int) [][]int {

	result := [][]int{}

	var bt func([]int, int)

	bt = func(track []int, index int) {

		// base case
		if index == len(nums) {
			result = append(result, append([]int(nil), track...))
			return
		}

		// choose, next layer traverse and cancel choose
		// if the ball is selected
		bt(append(track, nums[index]), index+1)
		// if the ball is NOT selected
		bt(track, index+1)
	}

	bt([]int{}, 0)
	return result
}
