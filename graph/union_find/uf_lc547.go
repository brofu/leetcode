package uf

func findCircleNum(isConnected [][]int) int {

	n := len(isConnected)
	uf := NewPathCompactedUF(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.Connect(i, j)
			}
		}
	}

	return uf.Count()
}
