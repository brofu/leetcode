package array

type NumMatrix struct {
	preSum [][]int
}

// KP. The set calculation
func ConstructorNumMatrix(matrix [][]int) NumMatrix {
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

	return NumMatrix{preSum: preSum}
}

// KP. the left number is included in the range
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}
