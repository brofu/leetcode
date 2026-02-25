package dp

import "sort"

/*
The Problem of "Longest arithmetic progression with fixed difference k", also known as
"Longest Subsequence with Difference k"（子序列，不要求连续）

KP:
	1.	Sub-sequence, NOT sub-array
	2.	If need to keep the order of array.
	3.  LeetCode 1218.
*/

/*
KP:
	DP Version 1.

SC:
	1. Sort: O(N*lgN)
	2. Loop: O(N^2)
	3. Result loop: O(N)
	4. Overall: O(N^2)
*/
func findLongestArithmeticProgression(arr []int32, k int32) int32 {
	// Write your code here
	if len(arr) <= 1 {
		return int32(len(arr))
	}

	sort.Slice(arr, func(x, y int) bool {
		return arr[x] < arr[y]
	})
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] == arr[j]+k {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	result := dp[0]
	for i := 1; i < len(dp); i++ {
		if result < dp[i] {
			result = dp[i]
		}
	}

	return int32(result)
}

/*
KP:
	1.	Keep the `order` of the array.

TC: O(N)
SC:	O(n)
*/
func findLongestArithmeticProgressionVersionKeepOrder(arr []int32, k int32) int32 {
	best := make(map[int32]int32)

	result := int32(0)
	for _, a := range arr {
		r := best[a-k] + 1
		if best[a] < r {
			best[a] = r
		}
		if result < best[a] {
			result = best[a]
		}
	}
	return result
}

/*
KP:
	1.	NO need to keep the `order` of the array.

TC: O(N*lgN)
SC: O(N)
*/
func findLongestArithmeticProgressionVersionNoOrder(arr []int32, k int32) int32 {

	sort.Slice(arr, func(x, y int) bool {
		return arr[x] < arr[y]
	})

	result := 0
	best := make(map[int32]int)
	for _, x := range arr {
		cur := best[x-k] + 1
		if best[x] < cur {
			best[x] = cur
		}
		if result < best[x] {
			result = best[x]
		}
	}

	return int32(result)
}
