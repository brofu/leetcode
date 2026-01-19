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

/*
KP:
	1. Worse control logic, comparing V6
	2. All the possible results on leaf nodes, and only collect the results on leaf nodes


TC:
	1. All the search nodes are O(2^N+2^N-1), which is O(2^(N+1)-1)
	2. Total results: O(2^n)
	2. For each result, need to copy. Worst case O(n)
	3. Overall, around O(N*2^N)

SC:
	1. Recursive depth: n
	2. For the result, O(n*2^n).
*/

func subsetsV5(nums []int) [][]int {

	var (
		result [][]int
		bt     func(int, []int)
	)

	bt = func(idx int, track []int) {
		if idx == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			result = append(result, temp)
			return
		}

		// choose, explore, cancel-choose
		bt(idx+1, track)
		bt(idx+1, append(track, nums[idx]))

	}

	bt(0, []int{})
	return result
}

/*
KP:
	1. Better control logic, comparing with V5
	2. Results on each node. Collect the results ON all the nodes.
	3. All the search nodes are O(2^N)
	4. Copy with each result, worst are O(N*2^N)


TC:
	1. total results: O(2^n)
	2. for each result, need to copy. Worst case O(n)
	3. Overall, around O(n*2^n)

SC:
	1. Recursive depth: n
	2. For the result, O(n*2^n).
*/

func subsetsV6(nums []int) [][]int {

	var (
		result [][]int
		bt     func(int, []int)
	)

	bt = func(idx int, track []int) {

		// base case
		if idx == len(nums)+1 {
			return
		}

		// collect result
		temp := make([]int, len(track))
		copy(temp, track)
		result = append(result, temp)

		for i := idx; i < len(nums); i++ {

			// choose
			// explore
			bt(i+1, append(track, nums[i]))
			// cancel choose
		}
	}

	bt(0, []int{})
	return result
}
