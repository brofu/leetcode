package backtrack

import "sort"

func combinationSum2(candidates []int, target int) [][]int {

	result := [][]int{}
	track := []int{}
	sum := 0

	sort.Ints(candidates)

	var bt func([]int, []int, int, int, int, *[][]int)
	bt = func(candidates, track []int, sum, target, start int, result *[][]int) {

		// base case
		if sum == target {
			temp := make([]int, len(track))
			copy(temp, track)
			*result = append(*result, temp)
			return
		}

		if sum >= target {
			return
		}

		for i := start; i < len(candidates); i++ {
			// pruning
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			// choose, depth + 1 traverse, cancel choose
			bt(candidates, append(track, candidates[i]), sum+candidates[i], target, i+1, result)
		}
	}

	bt(candidates, track, sum, target, 0, &result)
	return result
}

func combinationSum2V2(candidates []int, target int) [][]int {

	result := make([][]int, 0)
	sort.Ints(candidates)

	var bt func(int, int, []int)

	bt = func(start, target int, path []int) {

		//base case
		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if target-candidates[i] < 0 { // KP. Another prune
				continue
			}
			bt(i+1, target-candidates[i], append(path, candidates[i]))
		}
	}

	bt(0, target, []int{})
	return result
}

/*

TC:
	1. If there is NO duplicate number, then, the number (upper bound) of results are C(n, target/m), m is the minimum is candidates
	2. The overall time is O(C(n, target/m)*(target/m))

SC:
	1. Recursive depth: target/m
	2. Copy result: target/m
	3. Overall space complexity: O(C(n, target/m)*(target/m))

*/
func combinationSum2V3(candidates []int, target int) [][]int {

	var (
		result [][]int
		bt     func(int, int, []int)
	)

	sort.Ints(candidates)

	bt = func(idx, sum int, track []int) {

		// base case
		if sum == target {
			temp := make([]int, len(track))
			copy(temp, track)
			result = append(result, temp)
			return
		}
		if sum > target {
			return
		}
		if idx == len(candidates) {
			return
		}

		pre := 0
		for i := idx; i < len(candidates); i++ {
			// pruning
			if pre == candidates[i] {
				continue
			}
			pre = candidates[i]
			// choose, explore, cancel-choose
			bt(i+1, sum+candidates[i], append(track, candidates[i]))
		}

	}

	bt(0, 0, []int{})
	return result
}
