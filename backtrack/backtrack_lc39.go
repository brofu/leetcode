package backtrack

/*
KP.
1. How to control to select 1 element for multiple times?
2. What's the base case?
3. Compare with problem: 78, 77, and 90, 40

Time Complexity

1. Let's assume
	𝑛 = number of candidates
	𝑇 = target value
	𝑅 = number of valid result combinations
	minVal = min(candidates)

2. Max depth of the recursive tree:
	O(T/minVal) = O(T)

3. Branching Factor
	At each level:
	can pick any of the 𝑛 n candidates starting from start index
	So, branching factor ≤ 𝑛 (due to pruning and reusing only from current index onward, it's not a full n-ary tree)

4. Number of Recursive Calls
	Similar as `enumerating all **integer combinations** (with repetitions allowed) that sum to target using values from candidates`
	So, it's around O(2^T) (the number of path need to explore)

5. Cost per Recursive Call
	Copy cost: O(T)

6. Overall
	O(T*R), R <= 2^T
	Hence we have `O(T*2^T)`

Space Complexity
1. Recursive stack O(T), if 1 contained in the candidates
2. Result storage: O(T * R), R is the validated result
*/

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

func combinationSumPV1(candidates []int, target int) [][]int {

	result := [][]int{}
	track := []int{}

	var bt func([]int, int, int)

	bt = func(track []int, start, sum int) {

		// base case
		if sum == target {
			result = append(result, append([]int(nil), track...))
			return
		}

		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			bt(append(track, candidates[i]), i, sum+candidates[i])
		}
	}

	bt(track, 0, 0)

	return result
}

func combinationSumPV2(candidates []int, target int) [][]int {

	result := [][]int{}

	var bt func(start, sum int, path []int)

	bt = func(start, sum int, path []int) {

		//base case
		if sum == target { // find 1 result
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			bt(i, sum+candidates[i], append(path, candidates[i]))
		}

	}

	bt(0, 0, []int{})
	return result
}
