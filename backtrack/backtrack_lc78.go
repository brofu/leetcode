package backtrack

import "fmt"

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

// a WRONG version
// understand how to use `visited`
// and how to guarantee there is NO repeat subset
func subsetsV3(nums []int) [][]int {

	result := make([][]int, 0)
	visted := make([]bool, len(nums))

	var bt func(int, []int)

	bt = func(layer int, path []int) {

		fmt.Println("entry", layer, path, visted)
		//base
		if len(path) == layer {
			result = append(result, path)
		}

		for idx, num := range nums {

			// prune
			if visted[idx] {
				continue
			}

			// choose
			visted[idx] = true

			bt(layer+1, append(path, num))

			// cancel
			visted[idx] = false
		}

		fmt.Println("exit", layer, path, visted)
	}

	bt(0, []int{})
	return result
}

func subsetsV4(nums []int) [][]int {

	result := make([][]int, 0)

	var bt func(int, []int)

	bt = func(start int, path []int) {

		fmt.Println("entry:", start, result)
		//base
		result = append(result, path)

		for idx := start; idx < len(nums); idx++ {
			bt(idx+1, append(path, nums[idx]))
		}
		fmt.Println("exist:", start, result)
	}

	bt(0, []int{})
	return result
}
