package dp

import "github.com/brofu/leetcode/common"

/*
TC:
1. Initial dp[0], O(N)
2. Update dp[1:], O((M-1)*N), around O(M*N)
3. For each cell, we may need to 3 previous neighbors, that's O(3*MN)
4. Get the find result from dp[m-1], O(N)
5. Over all, it's O(M*N)

SC:
1. dp: O(M*N)
5. Over all, it's O(M*N)
*/
func minFallingPathSumDPA(matrix [][]int) int {

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	for idx, val := range matrix[0] {
		dp[0][idx] = val
	}

	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		v1, v2, v3 := 10001, 10001, 10001
		for j := 0; j < n; j++ {
			if j-1 >= 0 {
				v1 = dp[i-1][j-1]
			}
			v2 = dp[i-1][j]
			if j+1 < n {
				v3 = dp[i-1][j+1]
			}
			dp[i][j] = MinIntTrix(v1, v2, v3) + matrix[i][j]
		}
	}
	result := dp[m-1][0]
	for i := 1; i < n; i++ {
		result = MinInt(result, dp[m-1][i])
	}

	return result
}

/*
TC:
1. Construct memo: O(M*N)
2. Recursive call: Depth, O(M-1), at each layer, we need to check 3 neighbors, around O(3*(M-1))
3. Fort All the numbers of the last row, need to trigger the recursive call. that's O(N*3*(M-1))
4. Overall, it's O(M*N) + (3*N*(M-1)) = O(M*N)

SC:
1. Memo: O(MN)
*/
func minFallingPathSumDPF(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -10001
		}
	}

	var dp func(int, int) int

	dp = func(i, j int) int {

		// base case
		if i < 0 || i >= m || j < 0 || j >= n {
			return 10001
		}
		if memo[i][j] != -10001 {
			return memo[i][j]
		}
		// base case
		if i == 0 {
			return matrix[0][j]
		}

		memo[i][j] = MinIntTrix(dp(i-1, j-1), dp(i-1, j), dp(i-1, j+1)) + matrix[i][j]
		return memo[i][j]
	}

	result := 10001
	for idx := range matrix[m-1] {
		result = MinInt(result, dp(m-1, idx))
	}
	return result
}

/*
KP
1. dfs version
2. Actually, it's a 3-nary tree, with height as m
*/
func minFallingPathSumDFS(matrix [][]int) int {

	m, n := len(matrix), len(matrix[0])
	result := 10001

	var dfs func(i, j, sum int)
	dfs = func(i, j, sum int) {

		//base case
		if i == m-1 {
			result = MinInt(result, sum)
		}

		for _, neighbor := range [][]int{{i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1}} {
			// pruning
			x, y := neighbor[0], neighbor[1]
			if x >= m || y < 0 || y >= n {
				continue
			}
			dfs(x, y, sum+matrix[x][y])
		}
	}

	for idx, val := range matrix[0] {
		dfs(0, idx, val)
	}

	return result
}

/*

KP.
	1. DP function without Memo
	2. DP(i,j) means the min-fall-path which ends at matrix(i, j)
	3. The state transition equation is
		* matrix(0,j), if i = 0
		* matrix(i,j) + min(dp(i-1, j-1), dp(i-1, j), dp(i-1, j+1)), if i > 0

TC:
	1. The recursive depth is N.
	2. Every 1 upper layer, we need to calculate 3x nodes. That's to say,
		* at N layer, calculate N
		* at N-1 layer, calculate 3*N
		* at N-2 layer, calculate 3^2N
		* at 1 layer, calculate 3^(N-1)*N
		* so, all the recursive nodes are 3^N

SC:
	1. Recursive depth is O(N)

*/
func minFallingPathSumV1(matrix [][]int) int {

	var (
		result     = 10001
		n          = len(matrix)
		directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}}
		dp         func(int, int) int // the min-fall-path which ends at matrix(i, j)
	)

	dp = func(i, j int) int {

		// base case
		if j < 0 || j >= n {
			return -10001
		}
		if i == 0 {
			return matrix[0][j]
		}

		temp := 10001
		for _, direction := range directions {
			r := dp(i+direction[0], j+direction[1])
			if r != -10001 && temp > matrix[i][j]+r {
				temp = matrix[i][j] + r
			}
		}
		return temp
	}

	for j := 0; j < n; j++ {
		t := dp(n-1, j)
		if result > t {
			result = t
		}
	}
	return result
}

/*

TC:
	1. If there is Memo, then, we actually need to calculate ALL the cells in the matrix.
	2. That's O(N^2)
	3. Initial memo, another N^2
	4. Overall, TC is O(N^2)

SC:
	1. Recursive depth: N
	2. Memo size: N^2
	3. Overall O(N^2)
*/
func minFallingPathSumV1WithMemo(matrix [][]int) int {

	var (
		result     = 10001
		n          = len(matrix)
		memo       = make([][]int, len(matrix))
		directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}}
		dp         func(int, int) int // the min-fall-path which ends at matrix(i, j)
	)

	for i := 1; i < n; i++ { // O(N^2)
		memo[i] = make([]int, len(matrix))
		for j := 0; j < n; j++ {
			memo[i][j] = -10001
		}
	}

	dp = func(i, j int) int {

		// base case
		if j < 0 || j >= n {
			return -10001
		}
		if i == 0 {
			return matrix[0][j]
		}

		if memo[i][j] != -10001 {
			return memo[i][j]
		}

		temp := 10001
		for _, direction := range directions {
			r := dp(i+direction[0], j+direction[1])
			if r != -10001 && temp > matrix[i][j]+r {
				temp = matrix[i][j] + r
			}
		}
		memo[i][j] = temp
		return temp
	}

	for j := 0; j < n; j++ {
		t := dp(n-1, j)
		if result > t {
			result = t
		}
	}
	return result
}

/*

KP:
	1. DP table version.


TC:
	1. Overall SC is O(N^2)

SC:
	1. DP table O(N^2)
*/

func minFallingPathSumV1WithDP(matrix [][]int) int {

	if len(matrix) == 1 {
		return matrix[0][0]
	}

	var (
		result int
		n      = len(matrix)
		dp     = make([][]int, n)
	)

	dp[0] = make([]int, n)
	copy(dp[0], matrix[0])
	for i := 1; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = dp[i-1][j]
		}
	}

	for i := 1; i < n; i++ {
		dp[i][0] = common.MinInt(dp[i-1][0], dp[i-1][1]) + matrix[i][0]
		for j := 1; j < n-1; j++ {
			dp[i][j] = common.MinIntThree(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
		}
		dp[i][n-1] = common.MinInt(dp[i-1][n-2], dp[i-1][n-1]) + matrix[i][n-1]
	}

	result = dp[n-1][0]
	for j := 1; j < n; j++ {
		if result > dp[n-1][j] {
			result = dp[n-1][j]
		}
	}
	return result
}

func minFallingPathSumV1WithDPCompressed(matrix [][]int) int {

	if len(matrix) == 1 {
		return matrix[0][0]
	}

	var (
		result int
		n      = len(matrix)
		dp     = make([]int, n)
	)

	copy(dp, matrix[0])

	for i := 1; i < n; i++ {
	}

	result = dp[0]
	for i := 1; i < n; i++ {
		if result > dp[i] {
			result = dp[i]
		}
	}
	return result
}
