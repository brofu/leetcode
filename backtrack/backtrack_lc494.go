package backtrack

import "fmt"

func findTargetSumWays(nums []int, target int) int {

	var bt func(int, int) int

	bt = func(i int, remain int) int {

		// base case
		if i == len(nums) {
			if remain == 0 {
				return 1
			}
			return 0
		}

		result := 0
		// choose, next layer and cancel choose
		for _, op := range []int{-1, 1} {
			result += bt(i+1, remain+op*nums[i])
		}
		return result

	}

	return bt(0, target)
}

/*
*
Version with pruning
*/
func findTargetSumWaysWithMemo(nums []int, target int) int {

	memo := make(map[string]int)

	var bt func(int, int) int

	bt = func(i int, remain int) int {

		// base case
		if i == len(nums) {
			if remain == 0 {
				return 1
			}
			return 0
		}

		//pruning with memo
		str := fmt.Sprintf("%d-%d", i, remain)

		if val, ok := memo[str]; ok {
			return val
		}

		result := 0
		// choose, next layer and cancel choose
		for _, op := range []int{-1, 1} {
			result += bt(i+1, remain+op*nums[i])
		}
		memo[str] = result

		return result

	}

	return bt(0, target)
}
