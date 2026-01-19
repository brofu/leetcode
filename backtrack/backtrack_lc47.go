package backtrack

import "sort"

/**
KP
	1.	How to do pruning?
		假设输入为 nums = [1,2,2']，标准的全排列算法会得出如下答案：
		[ [1,2,2'],[1,2',2], [2,1,2'],[2,2',1], [2',1,2],[2',2,1] ]
		显然，这个结果存在重复，比如 [1,2,2'] 和 [1,2',2] 应该只被算作同一个排列，但被算作了两个不同的排列。
		所以现在的关键在于，如何设计剪枝逻辑，把这种重复去除掉？

		答案是，保证相同元素在排列中的相对位置保持不变。 比如说 nums = [1,2,2'] 这个例子，我保持排列中 2 一直在 2' 前面。

		这样的话，你从上面 6 个排列中只能挑出 3 个排列符合这个条件：
		[ [1,2,2'],[2,1,2'],[2,2',1] ]
		这也就是正确答案。

		进一步，如果 nums = [1,2,2',2'']，我只要保证重复元素 2 的相对位置固定，比如说 2 -> 2' -> 2''，也可以得到无重复的全排列结果。

		仔细思考，应该很容易明白其中的原理：
		标准全排列算法之所以出现重复，是因为把相同元素形成的排列序列视为不同的序列，但实际上它们应该是相同的；而如果固定相同元素形成的序列顺序，当然就避免了重复。

		那么反映到代码上，你注意看这个剪枝逻辑：

		// 新添加的剪枝逻辑，固定相同的元素在排列中的相对位置
		if (i > 0 && nums[i] == nums[i - 1] && !used[i - 1]) {
			// 如果前面的相邻相等元素没有用过，则跳过
			continue;
		}

		当出现重复元素时，比如输入 nums = [1,2,2',2'']，2' 只有在 2 已经被使用的情况下才会被选择，同理，2'' 只有在 2' 已经被使用的情况下才会被选择，这就保证了相同元素在排列中的相对位置保证固定。

	2. Draw the recursive tree for pruning.
		*	sort, and use `preNum`.
		*	Non-sort, use a map. At each recursive depth, n + (n-1) + (n-2) + ... + 1 = O(n^2)


Time Complexity
1. Recursive call times: ==> Wrong
	* if there are k repeated elements, with pruning:
	* then we have 1 + (n-k) + (n-k)(n-k-1) + ... + (n-k)!
	* totally we have e*(n-k)! times of recursive call
1. Recursive call times: ==> Right
	* let m as the number of distinct elements
	* Ki as the repeat times of the ith number in m
	* then, the `leaf nodes` are n!/(k1!*K2!*...*Km!)
2. Copy time cost O(n)
3. Sort (if we use sort): O(n*lgn)
4. Overall (n*n!/(K1!*K2!*...*Km!))

Space Complexity
1. Recursive stack O(n)
2. `visited` array O(n)
3. `hitMap` (if we use map) n + (n-1) + (n-2) + ... + 1 = O(n^2)
4. Result store: O(n*R), R is the result, R < n!
*/

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	track := []int{}
	used := make([]bool, len(nums))

	sort.Ints(nums)

	var bt func([]int, []int, []bool, *[][]int)

	bt = func(nums []int, track []int, used []bool, result *[][]int) {

		// base case
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			*result = append(*result, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// pruning
			if used[i] {
				continue
			}

			// pruning
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// choose, next layer and cancel choose
			used[i] = true
			bt(nums, append(track, nums[i]), used, result)
			used[i] = false
		}
	}

	bt(nums, track, used, &result)
	return result
}

/*
*
KP
 1. A simpler pruning solution
*/
func permuteUniqueV2(nums []int) [][]int {
	result := [][]int{}
	track := []int{}
	used := make([]bool, len(nums))

	sort.Ints(nums)

	var bt func([]int, []int, []bool, *[][]int)

	bt = func(nums []int, track []int, used []bool, result *[][]int) {

		// base case
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			*result = append(*result, temp)
			return
		}
		preNum := -11

		for i := 0; i < len(nums); i++ {
			// pruning
			if used[i] {
				continue
			}

			// pruning
			if preNum == nums[i] {
				continue
			}

			// choose, next layer and cancel choose
			used[i] = true
			preNum = nums[i]
			bt(nums, append(track, nums[i]), used, result)
			used[i] = false
		}
	}

	bt(nums, track, used, &result)
	return result
}

func permuteUniqueV3(nums []int) [][]int {

	result := make([][]int, 0)
	visted := make([]bool, len(nums))

	var bt func(track []int)

	bt = func(track []int) {

		//base case
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			result = append(result, temp)
			return
		}

		hitMap := make(map[int]struct{})
		for i := 0; i < len(nums); i++ {

			if visted[i] {
				continue
			}
			if _, exists := hitMap[nums[i]]; exists {
				continue
			}

			visted[i] = true
			hitMap[nums[i]] = struct{}{} // no need to cancel

			bt(append(track, nums[i]))

			visted[i] = false

		}
	}

	bt([]int{})
	return result
}

/*
KP
	1. Use layerMap to mark the usage of each layer

TC:
	1. The upper bound of number of search nodes is O(e*n!)
	2. For each search node, need to check if it's used, with TC O(1)
	3. For the final result, if there is U results, then, the tc is O(U*n), with copy n time for each result
	4. U = n!/(d1!*d2!*...*dk!), if there is K duplicated number, and each of them occurs `di` times in numbs
	5. So, the overall TC is around O(U*n)


SC:
	1. Recursive deepth: n
	2. Used: n
	3. Result: S*n
	4. layerMap: n, for each recursive layer. so, n * n
	5. So, overall SC is O(S*n + n^2)
*/

func permuteUniqueV5(nums []int) [][]int {

	var (
		result = make([][]int, 0)
		n      = len(nums)
		used   = make([]bool, n)
		bt     func([]int)
	)

	bt = func(track []int) {

		//base case
		if len(track) == n {
			temp := make([]int, n)
			copy(temp, track)
			result = append(result, temp)
			return
		}

		layerUsed := make(map[int]bool)
		for idx, num := range nums {

			// pruning
			if used[idx] {
				continue
			}
			if layerUsed[num] {
				continue
			}

			layerUsed[num] = true

			// choose
			used[idx] = true

			// explore
			bt(append(track, num))

			// cancel choose
			used[idx] = false
		}
	}

	bt([]int{})

	return result
}

func permuteUniqueV6(nums []int) [][]int {

	var (
		result [][]int
		used   = make([]bool, len(nums))
		bt     func([]int)
	)

	sort.Ints(nums) // O(n*lgn)

	bt = func(track []int) {

		// base case
		if len(track) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, track)
			result = append(result, temp)
			return
		}

		prev := 11
		for idx, num := range nums {

			// pruning
			if used[idx] {
				continue
			}
			if prev == num { // occurs in the same layer
				continue
			}

			// choose
			used[idx] = true
			prev = num

			// explore
			bt(append(track, num))

			// cancel choice
			used[idx] = false

		}
	}

	bt(make([]int, 0, len(nums)))
	return result
}
