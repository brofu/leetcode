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

/*

KP:
	1. 先对宽度 w 进行升序排序，如果遇到 w 相同的情况，则按照高度 h 降序排序；之后把所有的 h 作为一个数组，在这个数组上计算 LIS 的长度就是答案。
	2. The DP version. Can not pass the letcode testing

SC:
	1. sort. O(N*lgN)
	2. LIS on 2nd elements in envelops
*/

func maxEnvelopesV3(envelopes [][]int) int {

	var (
		result int
		dp     = make([]int, len(envelopes))
	)

	for idx := range dp {
		dp[idx] = 1
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	for i := 1; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	for _, temp := range dp {
		if result < temp {
			result = temp
		}
	}

	return result
}

func maxEnvelopesV4(envelopes [][]int) int {

	var (
		piles int
		top   = make([]int, len(envelopes))
	)

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	fmt.Println("flag", envelopes)

	for idx, num := range envelopes {

		left, right := 0, piles
		for left < right {
			mid := (left + right) / 2
			if num[1] < top[mid] {
				right = mid
			} else if num[1] > top[mid] {
				left = mid + 1
			} else {
				right = mid // get the left-most
			}
		}
		if left == piles {
			piles++
		}
		fmt.Println("flag", idx, left, piles)
		top[left] = num[1]
	}

	return piles
}
