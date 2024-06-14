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
