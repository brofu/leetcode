package array

/*
KP:
	1. 题目说 hours[i] 以 8 作为分界线，那么我们就要条件反射地想到对数据进行「归一化」处理，
	比如把所有大于 8 的元素视为 +1，把所有小于 8 的元素视为 -1，这样一来，这道题就改造成了：
	计算数组中元素和大于 0 的子数组的最大长度。
	2. 联系到前缀和 preSum[i] - preSum[j] > 0, preSum[i] > preSum[j]
	3. 建立单调栈。注意和453的区别：
		a. 453: Get the next Greater.
		b. 利用 monotonic decreasing stack，筛选掉无需比较的j. 
*/
func longestWPI(hours []int) int {

	result := 0

	preSum := make([]int, len(hours)+1)

	preSum[0] = 0
	for i := 1; i < len(preSum); i++ {
		if hours[i-1] > 8 {
			preSum[i] = preSum[i-1] + 1
		} else {
			preSum[i] = preSum[i-1] - 1
		}
	}

	st := make([]int, 0)
	for idx, sum := range preSum {
		if len(st) == 0 || sum < preSum[st[len(st)-1]] {
			st = append(st, idx)
		}
	}

	for i := len(preSum) - 1; i >= 0; i-- {
		for len(st) > 0 && preSum[i] > preSum[st[len(st)-1]] { // preSum[i] - preSum[j] > 0, good performance
			j := st[len(st)-1]
			if result < i-j {
				result = i - j
			}
			st = st[:len(st)-1]
		}
	}

	return result
}
