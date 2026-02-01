package array

import "github.com/brofu/leetcode/common"

/*
KP:
	1. The bounder of the case
*/
func matrixBlockSum(mat [][]int, k int) [][]int {

	c := Constructor304(mat)

	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	for idx := range ans {
		ans[idx] = make([]int, n)
	}

	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans[0]); j++ {
			x1 := common.MaxInt(0, i-k)
			y1 := common.MaxInt(0, j-k)
			x2 := common.MinInt(m-1, i+k)
			y2 := common.MinInt(n-1, j+k)
			ans[i][j] = c.SumRegion(x1, y1, x2, y2)
		}
	}
	return ans
}
