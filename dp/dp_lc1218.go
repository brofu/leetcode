package dp

/*
The Problem of "Longest arithmetic progression with fixed difference k", also known as
"Longest Subsequence with Difference k"（子序列，不要求连续）

KP:
	1.	Sub-sequence, NOT sub-array
	2.	If need to keep the order of array. This problem, need to keep the order.
*/
func longestSubsequence(arr []int, difference int) int {

	result := 0
	best := make(map[int]int)
	for _, x := range arr {
		cur := best[x-difference] + 1
		if best[x] < cur {
			best[x] = cur
		}
		if result < best[x] {
			result = best[x]
		}
	}
	return result
}
