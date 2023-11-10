package dp

import (
	"math"
)

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	min := math.MaxInt

	for _, coin := range coins {

		subproblem := coinChange(coins, amount-coin)

		if subproblem == -1 {
			continue
		}

		min = MinInt(subproblem+1, min)
	}

	if min == math.MaxInt {
		min = -1
	}

	return min
}

func coinChangeRecursiveWithMemo(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for index, _ := range memo {
		memo[index] = 0
	}

	return coinChangeWithMemoWrapper(memo, coins, amount)
}

func coinChangeWithMemoWrapper(memo []int, coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	if memo[amount] != 0 {
		return memo[amount]
	}

	max := amount + 1
	min := max

	for _, coin := range coins {
		subproblem := coinChangeWithMemoWrapper(memo, coins, amount-coin)

		if subproblem == -1 {
			continue
		}

		min = MinInt(subproblem+1, min)
	}

	if min == max {
		min = -1
	}

	memo[amount] = min

	return min
}

func coinChangeRecurrenceWithMemo(coins []int, amount int) int {

	extraLength := amount + 1

	memo := make([]int, extraLength)
	for index, _ := range memo {
		memo[index] = extraLength
	}
	memo[0] = 0

	for i := 1; i < extraLength; i++ {
		for _, coin := range coins {
			if i-coin >= 0 { // no result
				memo[i] = MinInt(memo[i], memo[i-coin]+1)
			}
		}
	}

	if memo[amount] == extraLength {
		memo[amount] = -1
	}

	return memo[amount]
}
