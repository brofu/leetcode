package array

type NumMatrixV0 struct {
	preSum [][]int
}

// KP. The set calculation
func ConstructorNumMatrix(matrix [][]int) NumMatrixV0 {
	m, n := len(matrix), len(matrix[0])

	preSum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		preSum[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			preSum[i][j] = preSum[i][j-1] + preSum[i-1][j] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}

	return NumMatrixV0{preSum: preSum}
}

// KP. the left number is included in the range
func (this *NumMatrixV0) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}

/*
KP.
	1. (row1, col1, row2, col2) here means, (row1, col1) is included in the range
*/
type NumMatrix struct {
	preSum [][]int
}

func Constructor304(matrix [][]int) NumMatrix {

	sum := make([][]int, len(matrix)+1)
	for idx := range sum {
		sum[idx] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i < len(sum); i++ {
		for j := 1; j < len(sum[0]); j++ {
			sum[i][j] = sum[i][j-1] + sum[i-1][j] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{
		preSum: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1] + this.preSum[row1][col1]
}
