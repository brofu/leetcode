package backtrack

/*
KP

# Time Complexity
 1. Recursive call time : O(n!)
    count ALL the nodes of the recursive tree:
    1 + n + n*(n-1) + n*(n-1)*(n-2) + ... + n!
    = sum(n!/(n-k)!) where k = 0, 1, ... n,
    = sum(n!/d!) where d = n-k
    = n! * sum(1/d!) where d = 0, 1, ... n, < n! * e
    so, we have O(e*n!)

2. Copy cost: O(n)
3. Overall: O(n*n!)

Space Complexity
1. Recursive stack: O(n)
2. Used array: O(n)
3. Result: O(n*n!)
*/
func permute(nums []int) [][]int {

	result := [][]int{}
	track := []int{}
	used := make(map[int]struct{}, len(nums))

	backtrack(nums, track, used, &result)
	return result
}

func backtrack(nums, track []int, used map[int]struct{}, answer *[][]int) {

	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*answer = append(*answer, temp)
		return
	}

	for _, num := range nums {
		if _, ok := used[num]; ok {
			continue
		}

		used[num] = struct{}{}
		track = append(track, num)

		backtrack(nums, track, used, answer)

		track = track[:len(track)-1]
		delete(used, num)
	}
}

func Permute(num []int) [][]int {
	return permute(num)
}

func permutePV1(nums []int) [][]int {

	result := [][]int{}
	used := make([]bool, len(nums))
	track := []int{}

	var backtrack func([]int, []int, []bool, *[][]int)
	backtrack = func(nums []int, track []int, used []bool, result *[][]int) {

		// the end
		if len(nums) == len(track) {
			temp := make([]int, len(track))
			copy(temp, track)
			*result = append(*result, temp)
			return
		}

		for index, num := range nums {

			if used[index] {
				continue
			}

			// choose
			track = append(track, num)
			// remove from choose list
			used[index] = true

			// go ahead
			backtrack(nums, track, used, result)

			// cancel the choose
			track = track[:len(track)-1]
			// add to the choose list
			used[index] = false
		}
	}

	backtrack(nums, track, used, &result)
	return result
}

// 2024.06.07
func permutePV2(nums []int) [][]int {

	result := [][]int{}
	track := []int{}
	used := make([]bool, len(nums))

	var bt func([]int, []int, []bool, *[][]int)
	bt = func(nums, track []int, used []bool, result *[][]int) {

		// base case
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			*result = append(*result, temp)
			return
		}

		for index, num := range nums {

			// filter
			if used[index] {
				continue
			}

			// choose
			used[index] = true

			// next layer traverse
			bt(nums, append(track, num), used, result)
			// cancel choose
			used[index] = false
		}
	}

	bt(nums, track, used, &result)
	return result
}

/**
KP.
	1.	Use `swap` to reduce the memory complexity
	2.	Select elements for each position.
*/

func permuteV2(nums []int) [][]int {
	result := [][]int{}

	var bt func(int)

	bt = func(start int) {

		// base case
		if start == len(nums) {
			result = append(result, append([]int(nil), nums...))
			return
		}

		for i := start; i < len(nums); i++ {

			// choose, traverse next layer, and cancel choose
			nums[start], nums[i] = nums[i], nums[start]
			bt(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	bt(0)

	return result
}

func permutePV3(nums []int) [][]int {

	result := [][]int{}
	hit := make([]bool, len(nums))

	var backtrack func([]int, []int)

	backtrack = func(path, nums []int) {

		if len(path) == len(nums) { // base case
			result = append(result, path)
			return
		}

		for i, num := range nums {
			if hit[i] {
				continue
			}
			hit[i] = true
			backtrack(append(path, num), nums)
			hit[i] = false
		}
	}

	backtrack([]int{}, nums)
	return result
}

func permutePV4(nums []int) [][]int {

	result := [][]int{}
	used := make([]bool, len(nums))

	var bt func([]int)

	bt = func(track []int) {

		//base case
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			result = append(result, temp)
		}

		for i := 0; i < len(nums); i++ {
			if used[i] { //prune
				continue
			}
			used[i] = true
			bt(append(track, nums[i]))
			used[i] = false
		}
	}

	bt([]int{})
	return result
}
