package backtrack

/*
Type: 子集/组合（元素可重不可复选）

KP:

1. Why cannot sort the nums?
*/

func findSubsequences(nums []int) [][]int {

	var result [][]int

	var bt func(int, []int)

	bt = func(index int, trace []int) {

		// base case
		if len(trace) >= 2 {
			temp := make([]int, len(trace))
			copy(temp, trace)
			result = append(result, temp)
		}

		used := make(map[int]struct{}) // KP. Record the used info with current layer

		for i := index; i < len(nums); i++ {

			//pruning
			if len(trace) > 0 && trace[len(trace)-1] > nums[i] {
				continue
			}
			if _, exists := used[nums[i]]; exists {
				continue
			}

			// choose
			used[nums[i]] = struct{}{}

			// next layer
			bt(i+1, append(trace, nums[i]))

			// cancel choose
			// delete(used, nums[i]) // KP. Why no need this?
		}
	}

	bt(0, []int{})

	return result
}
