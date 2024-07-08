package tree

import "fmt"

func countPalindromePaths(parent []int, s string) int64 {
	n := len(s)
	con := make([][]int, n)
	for i := 0; i < n; i++ {
		con[i] = []int{}
	}
	for i := 1; i < n; i++ {
		fmt.Println("flag", i, parent[i])
		con[parent[i]] = append(con[parent[i]], i)
	}
	fmt.Println("flag", con)
	have := make(map[int]int)
	have[0] = 1
	//return dfs(0, 0, s, con, have)
	return 0
}
