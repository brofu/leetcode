package array

import "sort"

func diagonalSort(mat [][]int) [][]int {

	m := make(map[int][]int)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			id := i - j
			m[id] = append(m[id], mat[i][j])
		}
	}
	for k := range m {
		sort.Ints(m[k])
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			id := i - j
			mat[i][j] = m[id][0]
			m[id] = m[id][1:]
		}
	}

	return mat
}

func diagonalSortV2(mat [][]int) [][]int {

	m := make(map[int][]int)
	p := make(map[int]int)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			id := i - j
			m[id] = append(m[id], mat[i][j])
		}
	}
	for k := range m {
		sort.Ints(m[k])
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			id := i - j
			mat[i][j] = m[id][p[id]]
			p[id]++
		}
	}

	return mat
}
