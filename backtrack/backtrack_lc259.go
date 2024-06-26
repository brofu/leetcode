package backtrack

/**
KP.
	1.	backtrack would lead to TLE
*/

func threeSumSmaller(nums []int, target int) int {
	result := 0

	var bt func([]int, int, int)

	bt = func(track []int, start, sum int) {

		// base case
		if len(track) == 3 {
			if sum < target {
				result++
			} else {
				return
			}
		}

		for i := start; i < len(nums); i++ {
			// choose, next layer, cancel choose
			track = append(track, nums[i])
			sum += nums[i]
			bt(track, i+1, sum)
			track = track[:len(track)-1]
			sum -= nums[i]
		}
	}

	bt([]int{}, 0, 0)
	return result
}
