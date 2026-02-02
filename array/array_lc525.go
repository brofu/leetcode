package array

import (
	"github.com/brofu/leetcode/common"
)

/*

KP
	1. 涉及到和为 xxx 的子数组，就是要考察前缀和技巧和哈希表的结合使用了
*/

func findMaxLength(nums []int) int {

	sum := make([]int, len(nums)+1)
	for i := 1; i < len(sum); i++ {
		if nums[i-1] == 1 {
			sum[i] = sum[i-1] + 1
		} else {
			sum[i] = sum[i-1] - 1
		}
	}
	m := make(map[int]int)
	res := 0

	for i := 0; i < len(sum); i++ {
		if idx, exists := m[sum[i]]; !exists {
			m[sum[i]] = i
		} else {
			res = common.MaxInt(res, i-idx)
		}
	}
	return res
}
