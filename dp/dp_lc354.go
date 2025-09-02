package dp

import (
	"fmt"
	"sort"
)

/*
KP:
1. 先对宽度w进行升序排序，如果遇到w相同的情况，则按照高度h降序排序；
2. 之后把所有的h作为一个数组，在这个数组上计算LIS的长度。

TC
1. sort. O(NlgN)
2. construct piles. O(NlgN)
3. Totally, 2O(NlgN), O(N*lgN)

SC:
1. top, O(N)
*/
func maxEnvelopesV2(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	fmt.Println("flag", envelopes)
	piles := 0
	top := make([]int, len(envelopes))

	for i := 0; i < len(envelopes); i++ {
		number := envelopes[i][1]
		left, right := 0, piles // [left, right)
		for left < right {
			mid := left + (right-left)/2
			if number > top[mid] {
				left = mid + 1
			} else if number < top[mid] {
				right = mid
			} else {
				right = mid // find the left most
			}
		}
		if left == piles {
			piles++
		}
		top[left] = number
	}
	return piles
}

/*
KP. Classic DP version. O(N^2)
*/
func maxEnvelopesDP(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, len(envelopes))
	for idx := range dp {
		dp[idx] = 1
	}
	for i := 1; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				dp[i] = MaxInt(dp[i], 1+dp[j])
			}
		}
	}
	result := 1
	for _, res := range dp {
		result = MaxInt(result, res)
	}

	return result
}
